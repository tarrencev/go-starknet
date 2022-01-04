package provider

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Block struct {
	TransactionReceipts []TransactionReceipt `json:"transaction_receipts"`
	Transactions        []Transaction        `json:"transactions"`
	Timestamp           int                  `json:"timestamp"`
	ParentBlockHash     string               `json:"parent_block_hash"`
	BlockHash           string               `json:"block_hash"`
	Status              string               `json:"status"`
	BlockNumber         int                  `json:"block_number"`
	StateRoot           string               `json:"state_root"`
}

type BlockOptions struct {
	BlockNumber int    `url:"blockNumber,omitempty"`
	BlockHash   string `url:"blockHash,omitempty"`
}

// Gets the block information from a block ID.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/f464ec4797361b6be8989e36e02ec690e74ef285/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L27-L31)
func (p *Provider) GetBlock(ctx context.Context, opts *BlockOptions) (*Block, error) {
	req, err := p.newRequest(ctx, http.MethodGet, "/get_block", nil)
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

	var resp Block
	return &resp, p.do(req, &resp)
}
