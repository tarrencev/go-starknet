// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package starknet

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tarrencev/go-starknet/contracts/abi"
	"github.com/tarrencev/go-starknet/contracts/abi/bind"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = bind.Bind
	_ = common.Big1
)

// TurnLog is an auto generated low-level Go binding around an user-defined struct.
type TurnLog struct {
	UserId                        *big.Int
	LocationId                    *big.Int
	BuyOrSell                     *big.Int
	ItemId                        *big.Int
	AmountToGive                  *big.Int
	MarketPreTradeItem            *big.Int
	MarketPostTradePreEventItem   *big.Int
	MarketPostTradePostEventItem  *big.Int
	MarketPreTradeMoney           *big.Int
	MarketPostTradePreEventMoney  *big.Int
	MarketPostTradePostEventMoney *big.Int
	UserPreTradeItem              *big.Int
	UserPostTradePreEventItem     *big.Int
	UserPostTradePostEventItem    *big.Int
	UserPreTradeMoney             *big.Int
	UserPostTradePreEventMoney    *big.Int
	UserPostTradePostEventMoney   *big.Int
	TradeOccursBool               *big.Int
	MoneyReductionFactor          *big.Int
	ItemReductionFactor           *big.Int
	RegionalItemReductionFactor   *big.Int
	DealerDashBool                *big.Int
	WrangleDashedDealerBool       *big.Int
	MuggingBool                   *big.Int
	RunFromMuggingBool            *big.Int
	GangWarBool                   *big.Int
	DefendGangWarBool             *big.Int
	CopRaidBool                   *big.Int
	BribeCopsBool                 *big.Int
	FindItemBool                  *big.Int
	LocalShipmentBool             *big.Int
	WarehouseSeizureBool          *big.Int
}

// DopeWarsMetaData contains all meta data concerning the DopeWars contract.
var DopeWarsMetaData = &bind.MetaData{
	ABI: "[{\"members\":[{\"name\":\"user_id\",\"offset\":0,\"type\":\"felt\"},{\"name\":\"location_id\",\"offset\":1,\"type\":\"felt\"},{\"name\":\"buy_or_sell\",\"offset\":2,\"type\":\"felt\"},{\"name\":\"item_id\",\"offset\":3,\"type\":\"felt\"},{\"name\":\"amount_to_give\",\"offset\":4,\"type\":\"felt\"},{\"name\":\"market_pre_trade_item\",\"offset\":5,\"type\":\"felt\"},{\"name\":\"market_post_trade_pre_event_item\",\"offset\":6,\"type\":\"felt\"},{\"name\":\"market_post_trade_post_event_item\",\"offset\":7,\"type\":\"felt\"},{\"name\":\"market_pre_trade_money\",\"offset\":8,\"type\":\"felt\"},{\"name\":\"market_post_trade_pre_event_money\",\"offset\":9,\"type\":\"felt\"},{\"name\":\"market_post_trade_post_event_money\",\"offset\":10,\"type\":\"felt\"},{\"name\":\"user_pre_trade_item\",\"offset\":11,\"type\":\"felt\"},{\"name\":\"user_post_trade_pre_event_item\",\"offset\":12,\"type\":\"felt\"},{\"name\":\"user_post_trade_post_event_item\",\"offset\":13,\"type\":\"felt\"},{\"name\":\"user_pre_trade_money\",\"offset\":14,\"type\":\"felt\"},{\"name\":\"user_post_trade_pre_event_money\",\"offset\":15,\"type\":\"felt\"},{\"name\":\"user_post_trade_post_event_money\",\"offset\":16,\"type\":\"felt\"},{\"name\":\"trade_occurs_bool\",\"offset\":17,\"type\":\"felt\"},{\"name\":\"money_reduction_factor\",\"offset\":18,\"type\":\"felt\"},{\"name\":\"item_reduction_factor\",\"offset\":19,\"type\":\"felt\"},{\"name\":\"regional_item_reduction_factor\",\"offset\":20,\"type\":\"felt\"},{\"name\":\"dealer_dash_bool\",\"offset\":21,\"type\":\"felt\"},{\"name\":\"wrangle_dashed_dealer_bool\",\"offset\":22,\"type\":\"felt\"},{\"name\":\"mugging_bool\",\"offset\":23,\"type\":\"felt\"},{\"name\":\"run_from_mugging_bool\",\"offset\":24,\"type\":\"felt\"},{\"name\":\"gang_war_bool\",\"offset\":25,\"type\":\"felt\"},{\"name\":\"defend_gang_war_bool\",\"offset\":26,\"type\":\"felt\"},{\"name\":\"cop_raid_bool\",\"offset\":27,\"type\":\"felt\"},{\"name\":\"bribe_cops_bool\",\"offset\":28,\"type\":\"felt\"},{\"name\":\"find_item_bool\",\"offset\":29,\"type\":\"felt\"},{\"name\":\"local_shipment_bool\",\"offset\":30,\"type\":\"felt\"},{\"name\":\"warehouse_seizure_bool\",\"offset\":31,\"type\":\"felt\"}],\"name\":\"TurnLog\",\"size\":32,\"type\":\"struct\"},{\"inputs\":[{\"name\":\"address_of_controller\",\"type\":\"felt\"}],\"name\":\"constructor\",\"outputs\":[],\"type\":\"constructor\"},{\"inputs\":[{\"name\":\"user_id\",\"type\":\"felt\"},{\"name\":\"location_id\",\"type\":\"felt\"},{\"name\":\"buy_or_sell\",\"type\":\"felt\"},{\"name\":\"item_id\",\"type\":\"felt\"},{\"name\":\"amount_to_give\",\"type\":\"felt\"}],\"name\":\"have_turn\",\"outputs\":[],\"type\":\"function\"},{\"inputs\":[],\"name\":\"read_game_clock\",\"outputs\":[{\"name\":\"clock\",\"type\":\"felt\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"game_clock_at_turn\",\"type\":\"felt\"}],\"name\":\"view_given_turn\",\"outputs\":[{\"name\":\"turn_log\",\"type\":\"TurnLog\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"user_id\",\"type\":\"felt\"}],\"name\":\"check_user_state\",\"outputs\":[{\"name\":\"money\",\"type\":\"felt\"},{\"name\":\"id1\",\"type\":\"felt\"},{\"name\":\"id2\",\"type\":\"felt\"},{\"name\":\"id3\",\"type\":\"felt\"},{\"name\":\"id4\",\"type\":\"felt\"},{\"name\":\"id5\",\"type\":\"felt\"},{\"name\":\"id6\",\"type\":\"felt\"},{\"name\":\"id7\",\"type\":\"felt\"},{\"name\":\"id8\",\"type\":\"felt\"},{\"name\":\"id9\",\"type\":\"felt\"},{\"name\":\"id10\",\"type\":\"felt\"},{\"name\":\"id11\",\"type\":\"felt\"},{\"name\":\"id12\",\"type\":\"felt\"},{\"name\":\"id13\",\"type\":\"felt\"},{\"name\":\"id14\",\"type\":\"felt\"},{\"name\":\"id15\",\"type\":\"felt\"},{\"name\":\"id16\",\"type\":\"felt\"},{\"name\":\"id17\",\"type\":\"felt\"},{\"name\":\"id18\",\"type\":\"felt\"},{\"name\":\"id19\",\"type\":\"felt\"},{\"name\":\"location\",\"type\":\"felt\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"location_id\",\"type\":\"felt\"},{\"name\":\"item_id\",\"type\":\"felt\"}],\"name\":\"check_market_state\",\"outputs\":[{\"name\":\"item_quantity\",\"type\":\"felt\"},{\"name\":\"money_quantity\",\"type\":\"felt\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"user_combat_stats_len\",\"type\":\"felt\"},{\"name\":\"user_combat_stats\",\"type\":\"felt*\"},{\"name\":\"drug_lord_combat_stats_len\",\"type\":\"felt\"},{\"name\":\"drug_lord_combat_stats\",\"type\":\"felt*\"}],\"name\":\"challenge_current_drug_lord\",\"outputs\":[],\"type\":\"function\"}]",
}

// DopeWarsABI is the input ABI used to generate the binding from.
// Deprecated: Use DopeWarsMetaData.ABI instead.
var DopeWarsABI = DopeWarsMetaData.ABI

// DopeWars is an auto generated Go binding around an Ethereum contract.
type DopeWars struct {
	DopeWarsCaller // Read-only binding to the contract
}

// DopeWarsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DopeWarsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DopeWarsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DopeWarsRaw struct {
	Contract *DopeWars // Generic contract binding to access the raw methods on
}

// DopeWarsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DopeWarsCallerRaw struct {
	Contract *DopeWarsCaller // Generic read-only contract binding to access the raw methods on
}

// NewDopeWars creates a new instance of DopeWars, bound to a specific deployed contract.
func NewDopeWars(address common.Hash, backend bind.ContractBackend) (*DopeWars, error) {
	contract, err := bindDopeWars(address, backend)
	if err != nil {
		return nil, err
	}
	return &DopeWars{DopeWarsCaller: DopeWarsCaller{contract: contract}}, nil
}

// NewDopeWarsCaller creates a new read-only instance of DopeWars, bound to a specific deployed contract.
func NewDopeWarsCaller(address common.Hash, caller bind.ContractCaller) (*DopeWarsCaller, error) {
	contract, err := bindDopeWars(address, caller)
	if err != nil {
		return nil, err
	}
	return &DopeWarsCaller{contract: contract}, nil
}

// bindDopeWars binds a generic wrapper to an already deployed contract.
func bindDopeWars(address common.Hash, caller bind.ContractCaller) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DopeWarsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DopeWars *DopeWarsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DopeWars.Contract.DopeWarsCaller.contract.Call(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DopeWars *DopeWarsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DopeWars.Contract.contract.Call(opts, result, method, params...)
}

// CheckMarketState is a free data retrieval call binding the contract method 0x3990913b.
//
// Solidity: function check_market_state(felt location_id, felt item_id) view returns(felt item_quantity, felt money_quantity)
func (_DopeWars *DopeWarsCaller) CheckMarketState(opts *bind.CallOpts, location_id *big.Int, item_id *big.Int) (struct {
	ItemQuantity  *big.Int
	MoneyQuantity *big.Int
}, error) {
	var out []interface{}
	err := _DopeWars.contract.Call(opts, &out, "check_market_state", location_id, item_id)

	outstruct := new(struct {
		ItemQuantity  *big.Int
		MoneyQuantity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ItemQuantity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MoneyQuantity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CheckUserState is a free data retrieval call binding the contract method 0x0ec6e44f.
//
// Solidity: function check_user_state(felt user_id) view returns(felt money, felt id1, felt id2, felt id3, felt id4, felt id5, felt id6, felt id7, felt id8, felt id9, felt id10, felt id11, felt id12, felt id13, felt id14, felt id15, felt id16, felt id17, felt id18, felt id19, felt location)
func (_DopeWars *DopeWarsCaller) CheckUserState(opts *bind.CallOpts, user_id *big.Int) (struct {
	Money    *big.Int
	Id1      *big.Int
	Id2      *big.Int
	Id3      *big.Int
	Id4      *big.Int
	Id5      *big.Int
	Id6      *big.Int
	Id7      *big.Int
	Id8      *big.Int
	Id9      *big.Int
	Id10     *big.Int
	Id11     *big.Int
	Id12     *big.Int
	Id13     *big.Int
	Id14     *big.Int
	Id15     *big.Int
	Id16     *big.Int
	Id17     *big.Int
	Id18     *big.Int
	Id19     *big.Int
	Location *big.Int
}, error) {
	var out []interface{}
	err := _DopeWars.contract.Call(opts, &out, "check_user_state", user_id)

	outstruct := new(struct {
		Money    *big.Int
		Id1      *big.Int
		Id2      *big.Int
		Id3      *big.Int
		Id4      *big.Int
		Id5      *big.Int
		Id6      *big.Int
		Id7      *big.Int
		Id8      *big.Int
		Id9      *big.Int
		Id10     *big.Int
		Id11     *big.Int
		Id12     *big.Int
		Id13     *big.Int
		Id14     *big.Int
		Id15     *big.Int
		Id16     *big.Int
		Id17     *big.Int
		Id18     *big.Int
		Id19     *big.Int
		Location *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Money = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Id1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Id2 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Id3 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Id4 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Id5 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Id6 = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Id7 = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Id8 = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Id9 = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Id10 = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.Id11 = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.Id12 = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	outstruct.Id13 = *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)
	outstruct.Id14 = *abi.ConvertType(out[14], new(*big.Int)).(**big.Int)
	outstruct.Id15 = *abi.ConvertType(out[15], new(*big.Int)).(**big.Int)
	outstruct.Id16 = *abi.ConvertType(out[16], new(*big.Int)).(**big.Int)
	outstruct.Id17 = *abi.ConvertType(out[17], new(*big.Int)).(**big.Int)
	outstruct.Id18 = *abi.ConvertType(out[18], new(*big.Int)).(**big.Int)
	outstruct.Id19 = *abi.ConvertType(out[19], new(*big.Int)).(**big.Int)
	outstruct.Location = *abi.ConvertType(out[20], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReadGameClock is a free data retrieval call binding the contract method 0x196d94ee.
//
// Solidity: function read_game_clock() view returns(felt clock)
func (_DopeWars *DopeWarsCaller) ReadGameClock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DopeWars.contract.Call(opts, &out, "read_game_clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ViewGivenTurn is a free data retrieval call binding the contract method 0xdc847c54.
//
// Solidity: function view_given_turn(felt game_clock_at_turn) view returns(TurnLog turn_log)
func (_DopeWars *DopeWarsCaller) ViewGivenTurn(opts *bind.CallOpts, game_clock_at_turn *big.Int) (TurnLog, error) {
	var out []interface{}
	err := _DopeWars.contract.Call(opts, &out, "view_given_turn", game_clock_at_turn)

	if err != nil {
		return *new(TurnLog), err
	}

	out0 := *abi.ConvertType(out[0], new(TurnLog)).(*TurnLog)

	return out0, err

}
