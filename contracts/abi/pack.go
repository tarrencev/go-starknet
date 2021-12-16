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

package abi

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common/math"
)

// packElement packs the given reflect value according to the abi specification in
// t.
func packElement(t Type, reflectValue reflect.Value) ([]byte, error) {
	switch t.T {
	case FeltTy:
		return packNum(reflectValue), nil
	default:
		return []byte{}, fmt.Errorf("could not pack element, unknown type: %v", t.T)
	}
}

// packNum packs the given number (using the reflect value) and will cast it to appropriate number representation.
func packNum(value reflect.Value) []byte {
	switch kind := value.Kind(); kind {
	case reflect.Ptr:
		return math.U256Bytes(new(big.Int).Set(value.Interface().(*big.Int)))
	default:
		panic("abi: fatal error")
	}
}
