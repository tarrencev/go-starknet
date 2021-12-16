package abi

import (
	"encoding/json"
	"fmt"
	"io"
)

// The ABI holds information about a contract's context and available
// invokable methods. It will allow you to type check function calls and
// packs data accordingly.
type ABI struct {
	Constructor Method
	Methods     map[string]Method
	Structs     map[string]Struct
}

// JSON returns a parsed ABI interface and error if it failed.
func JSON(reader io.Reader) (ABI, error) {
	dec := json.NewDecoder(reader)

	var abi ABI
	if err := dec.Decode(&abi); err != nil {
		return ABI{}, err
	}
	return abi, nil
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (abi *ABI) UnmarshalJSON(data []byte) error {
	var fields []struct {
		Type    string
		Name    string
		Size    uint64
		Inputs  []Argument
		Outputs []Argument
		Members Members

		// Status indicator which can be: "view",
		StateMutability string

		// Event relevant indicator represents the event is
		// declared as anonymous.
		Anonymous bool
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	abi.Methods = make(map[string]Method)
	abi.Structs = make(map[string]Struct)
	for _, field := range fields {
		switch field.Type {
		case "constructor":
			abi.Constructor = NewMethod("", "", Constructor, field.StateMutability, field.Inputs, nil)
		case "function":
			name := overloadedName(field.Name, func(s string) bool { _, ok := abi.Methods[s]; return ok })
			abi.Methods[name] = NewMethod(name, field.Name, Function, field.StateMutability, field.Inputs, field.Outputs)
		case "struct":
			abi.Structs[field.Name] = NewStruct(field.Name, field.Size, field.Members)
		default:
			return fmt.Errorf("abi: could not recognize type %v of field %v", field.Type, field.Name)
		}
	}
	return nil
}

// overloadedName returns the next available name for a given thing.
// Needed since solidity allows for overloading.
//
// e.g. if the abi contains Methods send, send1
// overloadedName would return send2 for input send.
//
// overloadedName works for methods, events and errors.
func overloadedName(rawName string, isAvail func(string) bool) string {
	name := rawName
	ok := isAvail(name)
	for idx := 0; ok; idx++ {
		name = fmt.Sprintf("%s%d", rawName, idx)
		ok = isAvail(name)
	}
	return name
}
