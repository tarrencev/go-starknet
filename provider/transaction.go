package provider

import (
	"context"
	"math/big"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-querystring/query"
)

type TransactionStatus struct {
	TxStatus  string `json:"tx_status"`
	BlockHash string `json:"block_hash"`
}

type Transaction struct {
	Status      string `json:"status"`
	Transaction struct {
		EntryPointType     string   `json:"entry_point_type"`
		EntryPointSelector string   `json:"entry_point_selector"`
		ContractAddress    string   `json:"contract_address"`
		TransactionHash    string   `json:"transaction_hash"`
		Calldata           []string `json:"calldata"`
		Signature          []string `json:"signature"`
		Type               string   `json:"type"`
	} `json:"transaction"`
	BlockNumber      int    `json:"block_number"`
	BlockHash        string `json:"block_hash"`
	TransactionIndex int    `json:"transaction_index"`
}

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

type TransactionOptions struct {
	TransactionId   uint64 `url:"transactionId,omitempty"`
	TransactionHash string `url:"transactionHash,omitempty"`
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

type TransactionStatusOptions struct {
	TransactionId   uint64 `url:"transactionId,omitempty"`
	TransactionHash string `url:"transactionHash,omitempty"`
}

// Gets the transaction status from a txn.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/fc97bdd8322a7df043c87c371634b26c15ed6cee/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L87)
func (p *Provider) TransactionStatus(ctx context.Context, opts TransactionStatusOptions) (*TransactionStatus, error) {
	req, err := p.newRequest(ctx, http.MethodGet, "/get_transaction_status", nil)
	if err != nil {
		return nil, err
	}
	vs, err := query.Values(opts)
	if err != nil {
		return nil, err
	}
	appendQueryValues(req, vs)

	var resp TransactionStatus
	return &resp, p.do(req, &resp)
}

// Gets the transaction id from its hash.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/fc97bdd8322a7df043c87c371634b26c15ed6cee/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L137)
func (p *Provider) TransactionID(ctx context.Context, hash common.Hash) (*big.Int, error) {
	req, err := p.newRequest(ctx, http.MethodGet, "/get_transaction_id_by_hash", nil)
	if err != nil {
		return nil, err
	}

	appendQueryValues(req, url.Values{
		"transactionHash": []string{hash.Hex()},
	})

	var resp big.Int
	return &resp, p.do(req, &resp)
}

// Gets the transaction hash from its id.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/fc97bdd8322a7df043c87c371634b26c15ed6cee/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L130)
func (p *Provider) TransactionHash(ctx context.Context, id *big.Int) (*common.Hash, error) {
	req, err := p.newRequest(ctx, http.MethodGet, "/get_transaction_hash_by_id", nil)
	if err != nil {
		return nil, err
	}

	appendQueryValues(req, url.Values{
		"transactionId": []string{id.String()},
	})

	var resp string
	if err := p.do(req, &resp); err != nil {
		return nil, err
	}

	hash := common.HexToHash(resp)
	return &hash, nil
}
