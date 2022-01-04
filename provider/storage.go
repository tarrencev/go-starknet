package provider

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-querystring/query"
)

type StorageAtOptions struct {
	BlockNumber int    `url:"blockNumber,omitempty"`
	BlockHash   string `url:"blockHash,omitempty"`
}

// Get a storage slots value.
//
// [Reference](https://github.com/starkware-libs/cairo-lang/blob/fc97bdd8322a7df043c87c371634b26c15ed6cee/src/starkware/starknet/services/api/feeder_gateway/feeder_gateway_client.py#L70)
func (p *Provider) StorageAt(ctx context.Context, address common.Hash, key uint64, opts *StorageAtOptions) (string, error) {
	req, err := p.newRequest(ctx, http.MethodGet, "/get_storage_at", nil)
	if err != nil {
		return "", err
	}

	appendQueryValues(req, url.Values{
		"contractAddress": []string{address.Hex()},
		"key":             []string{strconv.FormatUint(key, 10)},
	})

	if opts != nil {
		vs, err := query.Values(opts)
		if err != nil {
			return "", err
		}
		appendQueryValues(req, vs)
	}

	var resp string
	return resp, p.do(req, &resp)
}
