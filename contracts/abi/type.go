// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package abi

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
)

// Type enumerator
const (
	FeltTy byte = iota
	TupleTy
	PointerTy
	FunctionTy
	StructTy
)

// Type is the reflection of the supported argument type.
type Type struct {
	Elem *Type
	Size int
	T    byte // Our own type checking

	stringKind string // holds the unparsed string for deriving signatures

	// Tuple relative fields
	TupleRawName  string       // Raw struct name defined in source code, may be empty.
	TupleElems    []*Type      // Type information of all tuple fields
	TupleRawNames []string     // Raw field name of all tuple fields
	TupleType     reflect.Type // Underlying struct of the tuple
}

var (
	// typeRegex parses the abi sub types
	typeRegex = regexp.MustCompile("([a-zA-Z]+)(([0-9]+)(x([0-9]+))?)?")
)

// NewType creates a new reflection type of abi type given in t.
func NewType(t string, internalType string, components []ArgumentMarshaling) (typ Type, err error) {
	typ.stringKind = t

	// parse the type and size of the abi-type.
	matches := typeRegex.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return Type{}, fmt.Errorf("invalid type '%v'", t)
	}
	parsedType := matches[0]

	// varType is the parsed abi type
	switch varType := parsedType[1]; varType {
	case "felt":
		typ.Size = 251
		typ.T = FeltTy
	case "tuple":
		var (
			fields     []reflect.StructField
			elems      []*Type
			names      []string
			expression string // canonical parameter expression
		)
		expression += "("
		overloadedNames := make(map[string]string)
		for idx, c := range components {
			cType, err := NewType(c.Type, c.InternalType, c.Components)
			if err != nil {
				return Type{}, err
			}
			fieldName, err := overloadedArgName(c.Name, overloadedNames)
			if err != nil {
				return Type{}, err
			}
			overloadedNames[fieldName] = fieldName
			fields = append(fields, reflect.StructField{
				Name: fieldName, // reflect.StructOf will panic for any exported field.
				Type: cType.GetType(),
				Tag:  reflect.StructTag("json:\"" + c.Name + "\""),
			})
			elems = append(elems, &cType)
			names = append(names, c.Name)
			expression += cType.stringKind
			if idx != len(components)-1 {
				expression += ","
			}
		}
		expression += ")"

		typ.TupleType = reflect.StructOf(fields)
		typ.TupleElems = elems
		typ.TupleRawNames = names
		typ.T = TupleTy
		typ.stringKind = expression
	case "function":
		typ.T = FunctionTy
	default:
		typ.T = StructTy
	}

	return
}

// GetType returns the reflection type of the ABI type.
func (t Type) GetType() reflect.Type {
	switch t.T {
	case FeltTy:
		return reflect.TypeOf(&big.Int{})
	case PointerTy:
		return reflect.SliceOf(reflect.TypeOf(byte(0)))
	case FunctionTy:
		return reflect.ArrayOf(24, reflect.TypeOf(byte(0)))
	default:
		panic("Invalid type")
	}
}

func overloadedArgName(rawName string, names map[string]string) (string, error) {
	fieldName := ToCamelCase(rawName)
	if fieldName == "" {
		return "", errors.New("abi: purely anonymous or underscored field is not supported")
	}
	// Handle overloaded fieldNames
	_, ok := names[fieldName]
	for idx := 0; ok; idx++ {
		fieldName = fmt.Sprintf("%s%d", ToCamelCase(rawName), idx)
		_, ok = names[fieldName]
	}
	return fieldName, nil
}

// String implements Stringer.
func (t Type) String() (out string) {
	return t.stringKind
}

func (t Type) pack(v reflect.Value) ([]byte, error) {
	// dereference pointer first if it's a pointer
	v = indirect(v)
	if err := typeCheck(t, v); err != nil {
		return nil, err
	}

	switch t.T {
	case TupleTy:
		// (T1,...,Tk) for k >= 0 and any types T1, …, Tk
		// enc(X) = head(X(1)) ... head(X(k)) tail(X(1)) ... tail(X(k))
		// where X = (X(1), ..., X(k)) and head and tail are defined for Ti being a static
		// type as
		//     head(X(i)) = enc(X(i)) and tail(X(i)) = "" (the empty string)
		// and as
		//     head(X(i)) = enc(len(head(X(1)) ... head(X(k)) tail(X(1)) ... tail(X(i-1))))
		//     tail(X(i)) = enc(X(i))
		// otherwise, i.e. if Ti is a dynamic type.
		fieldmap, err := mapArgNamesToStructFields(t.TupleRawNames, v)
		if err != nil {
			return nil, err
		}
		// Calculate prefix occupied size.
		offset := 0
		for _, elem := range t.TupleElems {
			offset += getTypeSize(*elem)
		}
		var ret, tail []byte
		for i, elem := range t.TupleElems {
			field := v.FieldByName(fieldmap[t.TupleRawNames[i]])
			if !field.IsValid() {
				return nil, fmt.Errorf("field %s for tuple not found in the given struct", t.TupleRawNames[i])
			}
			val, err := elem.pack(field)
			if err != nil {
				return nil, err
			}
			if isDynamicType(*elem) {
				ret = append(ret, packNum(reflect.ValueOf(offset))...)
				tail = append(tail, val...)
				offset += len(val)
			} else {
				ret = append(ret, val...)
			}
		}
		return append(ret, tail...), nil

	default:
		return packElement(t, v)
	}
}

// isDynamicType returns true if the type is dynamic.
// The following types are called “dynamic”:
// * bytes
// * string
// * T[] for any T
// * T[k] for any dynamic T and any k >= 0
// * (T1,...,Tk) if Ti is dynamic for some 1 <= i <= k
func isDynamicType(t Type) bool {
	if t.T == TupleTy {
		for _, elem := range t.TupleElems {
			if isDynamicType(*elem) {
				return true
			}
		}
		return false
	}
	return t.T == PointerTy
}

// getTypeSize returns the size that this type needs to occupy.
// We distinguish static and dynamic types. Static types are encoded in-place
// and dynamic types are encoded at a separately allocated location after the
// current block.
// So for a static variable, the size returned represents the size that the
// variable actually occupies.
// For a dynamic variable, the returned size is fixed 32 bytes, which is used
// to store the location reference for actual value storage.
func getTypeSize(t Type) int {
	if t.T == TupleTy && !isDynamicType(t) {
		total := 0
		for _, elem := range t.TupleElems {
			total += getTypeSize(*elem)
		}
		return total
	}
	return 32
}
