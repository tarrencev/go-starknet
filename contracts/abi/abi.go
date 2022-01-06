package abi

import (
	"encoding/json"
	"errors"
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

func (abi ABI) getArguments(name string, data []byte) (Arguments, error) {
	// since there can't be naming collisions with contracts and events,
	// we need to decide whether we're calling a method or an event
	var args Arguments
	if method, ok := abi.Methods[name]; ok {
		if len(data)%32 != 0 {
			return nil, fmt.Errorf("abi: improperly formatted output: %s - Bytes: [%+v]", string(data), data)
		}
		args = method.Outputs
	}
	if args == nil {
		return nil, errors.New("abi: could not locate named method or event")
	}
	return args, nil
}

// Pack the given method name to conform the ABI. Method call's data
// will consist of method_id, args0, arg1, ... argN. Method id consists
// of 4 bytes and arguments are all 32 bytes.
// Method ids are created from the first 4 bytes of the hash of the
// methods string signature. (signature = baz(uint32,string32))
func (abi ABI) Pack(name string, args ...interface{}) ([]byte, error) {
	// Fetch the ABI of the requested method
	if name == "" {
		// constructor
		arguments, err := abi.Constructor.Inputs.Pack(args...)
		if err != nil {
			return nil, err
		}
		return arguments, nil
	}
	method, exist := abi.Methods[name]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", name)
	}
	arguments, err := method.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}
	// Pack up the method ID too if not a constructor and return
	return append(method.ID, arguments...), nil
}

// Unpack unpacks the output according to the abi specification.
func (abi ABI) Unpack(name string, data []byte) ([]interface{}, error) {
	args, err := abi.getArguments(name, data)
	if err != nil {
		return nil, err
	}
	return args.Unpack(data)
}

// UnpackIntoInterface unpacks the output in v according to the abi specification.
// It performs an additional copy. Please only use, if you want to unpack into a
// structure that does not strictly conform to the abi structure (e.g. has additional arguments)
func (abi ABI) UnpackIntoInterface(v interface{}, name string, data []byte) error {
	args, err := abi.getArguments(name, data)
	if err != nil {
		return err
	}
	unpacked, err := args.Unpack(data)
	if err != nil {
		return err
	}
	return args.Copy(v, unpacked)
}
