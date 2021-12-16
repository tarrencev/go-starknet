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

// Argument holds the name of the argument and the corresponding type.
// Types are used when packing and testing arguments.
type Member struct {
	Name   string
	Type   string
	Offset uint64
}

type Members []Member

// Struct represents
type Struct struct {
	Name    string
	Size    uint64
	Members Members
}

// NewStruct creates a new Struct.
// A struct should always be created using NewStruct.
func NewStruct(name string, size uint64, members Members) Struct {
	return Struct{
		Name:    name,
		Size:    size,
		Members: members,
	}
}
