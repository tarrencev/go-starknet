package bind

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

const abiJSON = `[
    {
        "members": [
            {
                "name": "user_id",
                "offset": 0,
                "type": "felt"
            },
            {
                "name": "location_id",
                "offset": 1,
                "type": "felt"
            },
            {
                "name": "buy_or_sell",
                "offset": 2,
                "type": "felt"
            },
            {
                "name": "item_id",
                "offset": 3,
                "type": "felt"
            },
            {
                "name": "amount_to_give",
                "offset": 4,
                "type": "felt"
            },
            {
                "name": "market_pre_trade_item",
                "offset": 5,
                "type": "felt"
            },
            {
                "name": "market_post_trade_pre_event_item",
                "offset": 6,
                "type": "felt"
            },
            {
                "name": "market_post_trade_post_event_item",
                "offset": 7,
                "type": "felt"
            },
            {
                "name": "market_pre_trade_money",
                "offset": 8,
                "type": "felt"
            },
            {
                "name": "market_post_trade_pre_event_money",
                "offset": 9,
                "type": "felt"
            },
            {
                "name": "market_post_trade_post_event_money",
                "offset": 10,
                "type": "felt"
            },
            {
                "name": "user_pre_trade_item",
                "offset": 11,
                "type": "felt"
            },
            {
                "name": "user_post_trade_pre_event_item",
                "offset": 12,
                "type": "felt"
            },
            {
                "name": "user_post_trade_post_event_item",
                "offset": 13,
                "type": "felt"
            },
            {
                "name": "user_pre_trade_money",
                "offset": 14,
                "type": "felt"
            },
            {
                "name": "user_post_trade_pre_event_money",
                "offset": 15,
                "type": "felt"
            },
            {
                "name": "user_post_trade_post_event_money",
                "offset": 16,
                "type": "felt"
            },
            {
                "name": "trade_occurs_bool",
                "offset": 17,
                "type": "felt"
            },
            {
                "name": "money_reduction_factor",
                "offset": 18,
                "type": "felt"
            },
            {
                "name": "item_reduction_factor",
                "offset": 19,
                "type": "felt"
            },
            {
                "name": "regional_item_reduction_factor",
                "offset": 20,
                "type": "felt"
            },
            {
                "name": "dealer_dash_bool",
                "offset": 21,
                "type": "felt"
            },
            {
                "name": "wrangle_dashed_dealer_bool",
                "offset": 22,
                "type": "felt"
            },
            {
                "name": "mugging_bool",
                "offset": 23,
                "type": "felt"
            },
            {
                "name": "run_from_mugging_bool",
                "offset": 24,
                "type": "felt"
            },
            {
                "name": "gang_war_bool",
                "offset": 25,
                "type": "felt"
            },
            {
                "name": "defend_gang_war_bool",
                "offset": 26,
                "type": "felt"
            },
            {
                "name": "cop_raid_bool",
                "offset": 27,
                "type": "felt"
            },
            {
                "name": "bribe_cops_bool",
                "offset": 28,
                "type": "felt"
            },
            {
                "name": "find_item_bool",
                "offset": 29,
                "type": "felt"
            },
            {
                "name": "local_shipment_bool",
                "offset": 30,
                "type": "felt"
            },
            {
                "name": "warehouse_seizure_bool",
                "offset": 31,
                "type": "felt"
            }
        ],
        "name": "TurnLog",
        "size": 32,
        "type": "struct"
    },
    {
        "inputs": [
            {
                "name": "address_of_controller",
                "type": "felt"
            }
        ],
        "name": "constructor",
        "outputs": [],
        "type": "constructor"
    },
    {
        "inputs": [
            {
                "name": "user_id",
                "type": "felt"
            },
            {
                "name": "location_id",
                "type": "felt"
            },
            {
                "name": "buy_or_sell",
                "type": "felt"
            },
            {
                "name": "item_id",
                "type": "felt"
            },
            {
                "name": "amount_to_give",
                "type": "felt"
            }
        ],
        "name": "have_turn",
        "outputs": [],
        "type": "function"
    },
    {
        "inputs": [],
        "name": "read_game_clock",
        "outputs": [
            {
                "name": "clock",
                "type": "felt"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "name": "game_clock_at_turn",
                "type": "felt"
            }
        ],
        "name": "view_given_turn",
        "outputs": [
            {
                "name": "turn_log",
                "type": "TurnLog"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "name": "user_id",
                "type": "felt"
            }
        ],
        "name": "check_user_state",
        "outputs": [
            {
                "name": "money",
                "type": "felt"
            },
            {
                "name": "id1",
                "type": "felt"
            },
            {
                "name": "id2",
                "type": "felt"
            },
            {
                "name": "id3",
                "type": "felt"
            },
            {
                "name": "id4",
                "type": "felt"
            },
            {
                "name": "id5",
                "type": "felt"
            },
            {
                "name": "id6",
                "type": "felt"
            },
            {
                "name": "id7",
                "type": "felt"
            },
            {
                "name": "id8",
                "type": "felt"
            },
            {
                "name": "id9",
                "type": "felt"
            },
            {
                "name": "id10",
                "type": "felt"
            },
            {
                "name": "id11",
                "type": "felt"
            },
            {
                "name": "id12",
                "type": "felt"
            },
            {
                "name": "id13",
                "type": "felt"
            },
            {
                "name": "id14",
                "type": "felt"
            },
            {
                "name": "id15",
                "type": "felt"
            },
            {
                "name": "id16",
                "type": "felt"
            },
            {
                "name": "id17",
                "type": "felt"
            },
            {
                "name": "id18",
                "type": "felt"
            },
            {
                "name": "id19",
                "type": "felt"
            },
            {
                "name": "location",
                "type": "felt"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "name": "location_id",
                "type": "felt"
            },
            {
                "name": "item_id",
                "type": "felt"
            }
        ],
        "name": "check_market_state",
        "outputs": [
            {
                "name": "item_quantity",
                "type": "felt"
            },
            {
                "name": "money_quantity",
                "type": "felt"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "name": "user_combat_stats_len",
                "type": "felt"
            },
            {
                "name": "user_combat_stats",
                "type": "felt*"
            },
            {
                "name": "drug_lord_combat_stats_len",
                "type": "felt"
            },
            {
                "name": "drug_lord_combat_stats",
                "type": "felt*"
            }
        ],
        "name": "challenge_current_drug_lord",
        "outputs": [],
        "type": "function"
    }
]`

func Test_Bind(t *testing.T) {
	src, err := Bind([]string{"DopeWars"}, []string{abiJSON})
	if err != nil {
		t.Fatalf("creating bindings: %+v", err)
	}

	if err := ioutil.WriteFile(filepath.Join("./test/", "dopewars.go"), []byte(src), 0600); err != nil {
		t.Fatalf("writing ABI binding: %v", err)
	}
}