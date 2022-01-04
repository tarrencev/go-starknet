package starknet

import "github.com/tarrencev/go-starknet/provider"

type Contract struct {
}

func NewContract(provider *provider.Provider) (*Contract, error) {
	return &Contract{}, nil
}
