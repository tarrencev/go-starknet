package starknet

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tarrencev/go-starknet/provider"
)

func Test_Dopewars(t *testing.T) {
	address := common.HexToHash("0x06ab912fbcc27c3f6ed4d998f2cc56fc65ab4fa6796259fc1855a099b0c92804")
	p := provider.NewProvider("https://alpha4.starknet.io")
	dope, err := NewDopeWarsCaller(address, p)
	if err != nil {
		t.Fatalf("Creating dope wars caller: %v+", err)
	}

	tl, err := dope.ViewGivenTurn(nil, big.NewInt(1))
	if err != nil {
		t.Fatalf("Reading turn log: %v+", err)
	}

	_ = tl
}
