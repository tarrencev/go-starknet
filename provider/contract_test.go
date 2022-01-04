package provider

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-cmp/cmp"
)

func TestProvider_CallContract(t *testing.T) {
	// TODO: Find a better test case
	want := "0x0"

	ctx := context.Background()
	p := NewProvider("https://alpha4.starknet.io")
	got, err := p.CallContract(ctx, CallContractRequest{
		ContractAddress: common.HexToHash("0x06ab912fbcc27c3f6ed4d998f2cc56fc65ab4fa6796259fc1855a099b0c92804"),
	}, nil)
	if err != nil {
		t.Fatalf("getting storage at: %v", err)
	}

	if diff := cmp.Diff(want, got, nil); diff != "" {
		t.Errorf("Storage value diff mismatch (-want +got):\n%s", diff)
	}
}
