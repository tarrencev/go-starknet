package provider

import (
	"context"
	"math/big"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

type Contracts struct {
	Starknet             string `json:"Starknet"`
	GpsStatementVerifier string `json:"GpsStatementVerifier"`
}

type CallMsg struct {
	Type               string       `json:"type,omitempty"`
	ContractAddress    *common.Hash `json:"contract_address"`
	EntryPointType     string       `json:"entry_point_type,omitempty"`
	EntryPointSelector string       `json:"entry_point_selector"`
	Signature          []*big.Int   `json:"signature"`
	Calldata           []string     `json:"calldata"`
}

type CallContractResponse struct {
	Result []string `json:"result"`
}

// Calls a contract.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/fc97bdd8322a7df043c87c371634b26c15ed6cee/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L25)
func (p *Provider) CallContract(ctx context.Context, call CallMsg, blockNumber *big.Int) ([]string, error) {
	call.Signature = []*big.Int{}

	req, err := p.newRequest(ctx, http.MethodPost, "/call_contract", call)
	if err != nil {
		return nil, err
	}

	if blockNumber != nil {
		appendQueryValues(req, url.Values{"blockNumber": []string{blockNumber.String()}})
	}

	var resp CallContractResponse
	if err := p.do(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
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
