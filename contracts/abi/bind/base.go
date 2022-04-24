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

package bind

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tarrencev/go-starknet/contracts/abi"
	"github.com/tarrencev/go-starknet/provider"
)

// SignerFn is a signer function callback when a contract requires a method to
// sign the transaction before submission.
type SignerFn func(common.Hash, *types.Transaction) (*types.Transaction, error)

// CallOpts is the collection of options to fine tune a contract call request.
type CallOpts struct {
	Pending     bool            // Whether to operate on the pending state or the last known one
	From        common.Hash     // Optional the sender address, otherwise the first account is used
	BlockNumber *big.Int        // Optional the block number on which the call should be performed
	Context     context.Context // Network context to support cancellation and timeouts (nil = no timeout)
}

// MetaData collects all metadata for a bound contract.
type MetaData struct {
	mu   sync.Mutex
	Sigs map[string]string
	Bin  string
	ABI  string
	ab   *abi.ABI
}

func (m *MetaData) GetAbi() (*abi.ABI, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.ab != nil {
		return m.ab, nil
	}
	if parsed, err := abi.JSON(strings.NewReader(m.ABI)); err != nil {
		return nil, err
	} else {
		m.ab = &parsed
	}
	return m.ab, nil
}

// BoundContract is the base wrapper object that reflects a contract on the
// Ethereum network. It contains a collection of methods that are used by the
// higher level contract bindings to operate.
type BoundContract struct {
	address common.Hash    // Deployment address of the contract on the Ethereum blockchain
	abi     abi.ABI        // Reflect based ABI to access the correct Ethereum methods
	caller  ContractCaller // Read interface to interact with the blockchain
}

// NewBoundContract creates a low level contract interface through which calls
// and transactions may be made through.
func NewBoundContract(address common.Hash, abi abi.ABI, caller ContractCaller) *BoundContract {
	return &BoundContract{
		address: address,
		abi:     abi,
		caller:  caller,
	}
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (c *BoundContract) Call(opts *CallOpts, results *[]interface{}, method string, params ...interface{}) error {
	// Don't crash on a lazy user
	if opts == nil {
		opts = new(CallOpts)
	}
	if results == nil {
		results = new([]interface{})
	}

	var input []string
	for _, p := range params {
		switch t := p.(type) {
		case *big.Int:
			input = append(input, t.String())
		}
	}

	selector, exist := c.abi.Methods[method]
	if !exist {
		return fmt.Errorf("method '%s' not found", method)
	}

	var (
		msg = provider.CallMsg{ContractAddress: &c.address, EntryPointSelector: hexutil.Encode(selector.ID), Calldata: input}
		ctx = ensureContext(opts.Context)
	)

	output, err := c.caller.CallContract(ctx, msg, opts.BlockNumber)
	if err != nil {
		return err
	}
	if len(output) == 0 {
		// Make sure we have a contract to operate on, and bail out otherwise.
		if cc, err := c.caller.CodeAt(ctx, c.address, opts.BlockNumber); err != nil {
			return err
		} else if len(cc.Bytecode) == 0 {
			return ErrNoCode
		}
	}

	if len(*results) == 0 {
		res, err := c.abi.Unpack(method, output)
		*results = res
		return err
	}
	// res := *results
	return nil //c.abi.UnpackIntoInterface(res[0], method, output)
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
