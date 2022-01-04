package provider

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

type TransactionReceipt struct {
	TransactionIndex   int `json:"transaction_index"`
	ExecutionResources struct {
		NSteps                 int `json:"n_steps"`
		NMemoryHoles           int `json:"n_memory_holes"`
		BuiltinInstanceCounter struct {
			PedersenBuiltin   int `json:"pedersen_builtin"`
			RangeCheckBuiltin int `json:"range_check_builtin"`
			BitwiseBuiltin    int `json:"bitwise_builtin"`
			OutputBuiltin     int `json:"output_builtin"`
			EcdsaBuiltin      int `json:"ecdsa_builtin"`
			EcOpBuiltin       int `json:"ec_op_builtin"`
		} `json:"builtin_instance_counter"`
	} `json:"execution_resources"`
	L2ToL1Messages  []interface{} `json:"l2_to_l1_messages"`
	BlockHash       string        `json:"block_hash"`
	TransactionHash string        `json:"transaction_hash"`
	Status          string        `json:"status"`
	BlockNumber     int           `json:"block_number"`
}

type Transaction struct {
	Signature          []interface{} `json:"signature"`
	EntryPointSelector string        `json:"entry_point_selector"`
	EntryPointType     string        `json:"entry_point_type"`
	ContractAddress    string        `json:"contract_address"`
	Calldata           []string      `json:"calldata"`
	TransactionHash    string        `json:"transaction_hash"`
	Type               string        `json:"type"`
}

type TransactionOptions struct {
	TransactionIndex uint64 `url:"transactionId,omitempty"`
	TransactionHash  string `url:"transactionHash,omitempty"`
}

// Gets the transaction information from a tx id.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/f464ec4797361b6be8989e36e02ec690e74ef285/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L54-L58)
func (p *Provider) Transaction(ctx context.Context, opts TransactionOptions) (*Transaction, error) {
	req, err := p.newRequest(ctx, http.MethodGet, "/get_transaction", nil)
	if err != nil {
		return nil, err
	}
	vs, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	appendQueryValues(req, vs)

	var resp Transaction
	return &resp, p.do(req, &resp)
}
