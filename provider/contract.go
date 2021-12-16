package provider

import (
	"context"
	"net/http"
)

type Contracts struct {
	Starknet             string `json:"Starknet"`
	GpsStatementVerifier string `json:"GpsStatementVerifier"`
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
