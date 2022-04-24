// Copyright 2016 The go-ethereum Authors
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

package bind

import "github.com/tarrencev/go-starknet/contracts/abi"

// tmplData is the data structure required to fill the binding template.
type tmplData struct {
	Package   string                   // Name of the package to place the generated file in
	Contracts map[string]*tmplContract // List of contracts to generate into this file
	Structs   map[string]*tmplStruct   // Contract struct type definitions
}

// tmplContract contains the data needed to generate an individual contract binding.
type tmplContract struct {
	Type        string                 // Type name of the main contract binding
	InputABI    string                 // JSON ABI used as the input to generate the binding from
	InputBin    string                 // Optional EVM bytecode used to generate deploy code from
	FuncSigs    map[string]string      // Optional map: string signature -> 4-byte signature
	Constructor abi.Method             // Contract constructor for deploy parametrization
	Calls       map[string]*tmplMethod // Contract calls that only read state data
	Events      map[string]*tmplEvent  // Contract events accessors
}

// tmplMethod is a wrapper around an abi.Method that contains a few preprocessed
// and cached data fields.
type tmplMethod struct {
	Original   abi.Method // Original method as parsed by the abi package
	Normalized abi.Method // Normalized version of the parsed method (capitalized names, non-anonymous args/returns)
	Structured bool       // Whether the returns should be accumulated into a struct
}

// tmplEvent is a wrapper around an abi.Event that contains a few preprocessed
// and cached data fields.
type tmplEvent struct {
	Original   abi.Event // Original event as parsed by the abi package
	Normalized abi.Event // Normalized version of the parsed fields
}

// tmplField is a wrapper around a struct field with binding language
// struct type definition and relative filed name.
type tmplField struct {
	Type    string   // Field type representation depends on target binding language
	Name    string   // Field name converted from the raw user-defined field name
	SolKind abi.Type // Raw abi type information
}

// tmplStruct is a wrapper around an abi.tuple and contains an auto-generated
// struct name.
type tmplStruct struct {
	Name   string       // Auto-generated struct name(before solidity v0.5.11) or raw name.
	Fields []*tmplField // Struct fields definition depends on the binding language.
}

// tmplSourceGo is the Go source template that the generated Go contract binding
// is based on.
const tmplSourceGo = `
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.Package}}

import (
	"math/big"
	"strings"
	"errors"

	"github.com/tarrencev/go-starknet/contracts/abi"
	"github.com/tarrencev/go-starknet/contracts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = bind.Bind
	_ = common.Big1
)

{{$structs := .Structs}}
{{range $structs}}
	// {{.Name}} is an auto generated low-level Go binding around an user-defined struct.
	type {{.Name}} struct {
	{{range $field := .Fields}}
	{{$field.Name}} {{$field.Type}}{{end}}
	}
{{end}}

{{range $contract := .Contracts}}
	// {{.Type}}MetaData contains all meta data concerning the {{.Type}} contract.
	var {{.Type}}MetaData = &bind.MetaData{
		ABI: "{{.InputABI}}",
		{{if $contract.FuncSigs -}}
		Sigs: map[string]string{
			{{range $strsig, $binsig := .FuncSigs}}"{{$binsig}}": "{{$strsig}}",
			{{end}}
		},
		{{end -}}
		{{if .InputBin -}}
		Bin: "0x{{.InputBin}}",
		{{end}}
	}
	// {{.Type}}ABI is the input ABI used to generate the binding from.
	// Deprecated: Use {{.Type}}MetaData.ABI instead.
	var {{.Type}}ABI = {{.Type}}MetaData.ABI

	{{if $contract.FuncSigs}}
		// Deprecated: Use {{.Type}}MetaData.Sigs instead.
		// {{.Type}}FuncSigs maps the 4-byte function signature to its string representation.
		var {{.Type}}FuncSigs = {{.Type}}MetaData.Sigs
	{{end}}

	// {{.Type}} is an auto generated Go binding around an Ethereum contract.
	type {{.Type}} struct {
	  {{.Type}}Caller     // Read-only binding to the contract
	}

	// {{.Type}}Caller is an auto generated read-only Go binding around an Ethereum contract.
	type {{.Type}}Caller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// {{.Type}}Raw is an auto generated low-level Go binding around an Ethereum contract.
	type {{.Type}}Raw struct {
	  Contract *{{.Type}} // Generic contract binding to access the raw methods on
	}

	// {{.Type}}CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type {{.Type}}CallerRaw struct {
		Contract *{{.Type}}Caller // Generic read-only contract binding to access the raw methods on
	}

	// New{{.Type}} creates a new instance of {{.Type}}, bound to a specific deployed contract.
	func New{{.Type}}(address common.Hash, backend bind.ContractBackend) (*{{.Type}}, error) {
	  contract, err := bind{{.Type}}(address, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &{{.Type}}{ {{.Type}}Caller: {{.Type}}Caller{contract: contract} }, nil
	}

	// New{{.Type}}Caller creates a new read-only instance of {{.Type}}, bound to a specific deployed contract.
	func New{{.Type}}Caller(address common.Hash, caller bind.ContractCaller) (*{{.Type}}Caller, error) {
	  contract, err := bind{{.Type}}(address, caller)
	  if err != nil {
	    return nil, err
	  }
	  return &{{.Type}}Caller{contract: contract}, nil
	}

	// bind{{.Type}} binds a generic wrapper to an already deployed contract.
	func bind{{.Type}}(address common.Hash, caller bind.ContractCaller) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader({{.Type}}ABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_{{$contract.Type}} *{{$contract.Type}}Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _{{$contract.Type}}.Contract.{{$contract.Type}}Caller.contract.Call(opts, result, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_{{$contract.Type}} *{{$contract.Type}}CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
		return _{{$contract.Type}}.Contract.contract.Call(opts, result, method, params...)
	}

	{{range .Calls}}
		// {{.Normalized.Name}} is a free data retrieval call binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}Caller) {{.Normalized.Name}}(opts *bind.CallOpts {{range .Normalized.Inputs}}, {{.Name}} {{bindtype .Type}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type}},{{end}}{{end}} error) {
			var out []interface{}
			err := _{{$contract.Type}}.contract.Call(opts, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
			{{if .Structured}}
			outstruct := new(struct{ {{range .Normalized.Outputs}} {{.Name}} {{bindtype .Type}}; {{end}} })
			if err != nil {
				return *outstruct, err
			}
			{{range $i, $t := .Normalized.Outputs}} 
			outstruct.{{.Name}} = *abi.ConvertType(out[{{$i}}], new({{bindtype .Type}})).(*{{bindtype .Type}}){{end}}

			return *outstruct, err
			{{else}}
			if err != nil {
				return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type}}), {{end}} err
			}
			{{range $i, $t := .Normalized.Outputs}}
			out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type}})).(*{{bindtype .Type}}){{end}}
			
			return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
			{{end}}
		}
	{{end}}
{{end}}
`
