package provider

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const transactionDump = `{"status": "ACCEPTED_ON_L2", "transaction": {"entry_point_type": "EXTERNAL", "entry_point_selector": "0x240060cdb34fcc260f41eac7474ee1d7c80b7e3607daff9ac67c7ea2ebb1c44", "contract_address": "0x26704e810b71eb2b99081e5093a01773f69b37163f8256c59dbd323bbab8b3d", "transaction_hash": "0x460c647079d51300fa578ae674d80f74febbb957143929c49599f86707a3850", "calldata": ["2997654770182471460954719312950884680023265536996318019468425755755830882165", "1329909728320632088402217562277154056711815095720684343816173432540100887380", "3", "1086644808292048134107134757058009361393820425772757521779863424964944956221", "1000000000000000000000", "0", "0"], "signature": ["2624462889305163041535477652824376920771700600796484323994636651237400740799", "1159601377385057436007767188985818160384069704534186283200528616638142189836"], "type": "INVOKE_FUNCTION"}, "block_number": 582, "block_hash": "0x182d83f0ed972e97fa529be0088e20b5a7cb32e3bba0d164d668147713e79f9", "transaction_index": 19}`

func TestProvider_Transaction(t *testing.T) {
	var want *Transaction
	if err := json.Unmarshal([]byte(transactionDump), &want); err != nil {
		t.Fatalf("unmarshalling block: %v", err)
	}

	for _, tc := range []struct {
		opts TransactionOptions
	}{{
		opts: TransactionOptions{TransactionHash: "0x460c647079d51300fa578ae674d80f74febbb957143929c49599f86707a3850"},
	}, {
		opts: TransactionOptions{TransactionIndex: 19},
	}} {
		ctx := context.Background()
		p := NewProvider("https://alpha-mainnet.starknet.io")
		got, err := p.GetTransaction(ctx, tc.opts)
		if err != nil {
			t.Fatalf("getting transaction: %v", err)
		}

		if diff := cmp.Diff(want, got, nil); diff != "" {
			t.Errorf("Transaction diff mismatch (-want +got):\n%s", diff)
		}
	}
}
