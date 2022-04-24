package bind

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tarrencev/go-starknet/provider"
)

var (
	// ErrNoCode is returned by call and transact operations for which the requested
	// recipient contract to operate on does not exist in the state db or does not
	// have any code associated with it (i.e. suicided).
	ErrNoCode = errors.New("no contract code at given address")

	// ErrNoPendingState is raised when attempting to perform a pending state action
	// on a backend that doesn't implement PendingContractCaller.
	ErrNoPendingState = errors.New("backend does not support pending state")

	// ErrNoCodeAfterDeploy is returned by WaitDeployed if contract creation leaves
	// an empty contract behind.
	ErrNoCodeAfterDeploy = errors.New("no contract code after deployment")
)

// ContractCaller defines the methods needed to allow operating with a contract on a read
// only basis.
type ContractCaller interface {
	// CodeAt returns the code of the given account. This is needed to differentiate
	// between contract internal errors and the local chain being out of sync.
	CodeAt(ctx context.Context, contract common.Hash, blockNumber *big.Int) (*provider.Code, error)

	// CallContract executes an Ethereum contract call with the specified data as the
	// input.
	CallContract(ctx context.Context, call provider.CallMsg, blockNumber *big.Int) ([]string, error)
}

// PendingContractCaller defines methods to perform contract calls on the pending state.
// Call will try to discover this interface when access to the pending state is requested.
// If the backend does not support the pending state, Call returns ErrNoPendingState.
type PendingContractCaller interface {
	// PendingCodeAt returns the code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, contract common.Hash) ([]byte, error)

	// PendingCallContract executes an Ethereum contract call against the pending state.
	PendingCallContract(ctx context.Context, call provider.CallMsg) ([]byte, error)
}

// ContractBackend defines the methods needed to work with contracts on a read-write basis.
type ContractBackend interface {
	ContractCaller
}
