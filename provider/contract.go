package provider

import (
	"context"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-querystring/query"
)

type Contracts struct {
	Starknet             string `json:"Starknet"`
	GpsStatementVerifier string `json:"GpsStatementVerifier"`
}

type CallContractOptions struct {
	BlockNumber int    `url:"blockNumber,omitempty"`
	BlockHash   string `url:"blockHash,omitempty"`
}

type CallContractRequest struct {
	Type               string      `json:"type"`
	ContractAddress    common.Hash `json:"contractAddress,omitempty"`
	EntryPointType     string      `json:"entry_point_type,omitempty"`
	EntryPointSelector common.Hash `json:"entry_point_selector,omitempty"`
	Signature          [2]*big.Int `json:"signature"`
	Calldata           []string    `json:"calldata"`
}

type CallContractResponse struct {
	Result []string `json:"result"`
}

// Calls a contract.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/fc97bdd8322a7df043c87c371634b26c15ed6cee/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L25)
func (p *Provider) CallContract(ctx context.Context, call CallContractRequest, opts *CallContractOptions) (*CallContractResponse, error) {
	call.Type = "INVOKE_FUNCTION"
	call.EntryPointType = "EXTERNAL"

	req, err := p.newRequest(ctx, http.MethodPost, "/call_contract", call)
	if err != nil {
		return nil, err
	}

	if opts != nil {
		vs, err := query.Values(opts)
		if err != nil {
			return nil, err
		}
		appendQueryValues(req, vs)
	}

	var resp CallContractResponse
	return &resp, p.do(req, &resp)
}

// Gets the smart contract address on the goerli testnet.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/f464ec4797361b6be8989e36e02ec690e74ef285/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L13-L15)
func (p *Provider) GetContractAddresses(ctx context.Context) (*Contracts, error) {
	req, err := p.newRequest(ctx, http.MethodGet, "/get_contract_addresses", nil)
	if err != nil {
		return nil, err
	}

	var resp Contracts
	return &resp, p.do(req, &resp)
}
