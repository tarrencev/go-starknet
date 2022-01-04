// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package starknet

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TurnLog is an auto generated low-level Go binding around an user-defined struct.
type TurnLog struct {
	user_id                            *big.Int
	location_id                        *big.Int
	buy_or_sell                        *big.Int
	item_id                            *big.Int
	amount_to_give                     *big.Int
	market_pre_trade_item              *big.Int
	market_post_trade_pre_event_item   *big.Int
	market_post_trade_post_event_item  *big.Int
	market_pre_trade_money             *big.Int
	market_post_trade_pre_event_money  *big.Int
	market_post_trade_post_event_money *big.Int
	user_pre_trade_item                *big.Int
	user_post_trade_pre_event_item     *big.Int
	user_post_trade_post_event_item    *big.Int
	user_pre_trade_money               *big.Int
	user_post_trade_pre_event_money    *big.Int
	user_post_trade_post_event_money   *big.Int
	trade_occurs_bool                  *big.Int
	money_reduction_factor             *big.Int
	item_reduction_factor              *big.Int
	regional_item_reduction_factor     *big.Int
	dealer_dash_bool                   *big.Int
	wrangle_dashed_dealer_bool         *big.Int
	mugging_bool                       *big.Int
	run_from_mugging_bool              *big.Int
	gang_war_bool                      *big.Int
	defend_gang_war_bool               *big.Int
	cop_raid_bool                      *big.Int
	bribe_cops_bool                    *big.Int
	find_item_bool                     *big.Int
	local_shipment_bool                *big.Int
	warehouse_seizure_bool             *big.Int
}

// TypeMetaData contains all meta data concerning the Type contract.
var TypeMetaData = &bind.MetaData{
	ABI: "[{\"members\":[{\"name\":\"user_id\",\"offset\":0,\"type\":\"felt\"},{\"name\":\"location_id\",\"offset\":1,\"type\":\"felt\"},{\"name\":\"buy_or_sell\",\"offset\":2,\"type\":\"felt\"},{\"name\":\"item_id\",\"offset\":3,\"type\":\"felt\"},{\"name\":\"amount_to_give\",\"offset\":4,\"type\":\"felt\"},{\"name\":\"market_pre_trade_item\",\"offset\":5,\"type\":\"felt\"},{\"name\":\"market_post_trade_pre_event_item\",\"offset\":6,\"type\":\"felt\"},{\"name\":\"market_post_trade_post_event_item\",\"offset\":7,\"type\":\"felt\"},{\"name\":\"market_pre_trade_money\",\"offset\":8,\"type\":\"felt\"},{\"name\":\"market_post_trade_pre_event_money\",\"offset\":9,\"type\":\"felt\"},{\"name\":\"market_post_trade_post_event_money\",\"offset\":10,\"type\":\"felt\"},{\"name\":\"user_pre_trade_item\",\"offset\":11,\"type\":\"felt\"},{\"name\":\"user_post_trade_pre_event_item\",\"offset\":12,\"type\":\"felt\"},{\"name\":\"user_post_trade_post_event_item\",\"offset\":13,\"type\":\"felt\"},{\"name\":\"user_pre_trade_money\",\"offset\":14,\"type\":\"felt\"},{\"name\":\"user_post_trade_pre_event_money\",\"offset\":15,\"type\":\"felt\"},{\"name\":\"user_post_trade_post_event_money\",\"offset\":16,\"type\":\"felt\"},{\"name\":\"trade_occurs_bool\",\"offset\":17,\"type\":\"felt\"},{\"name\":\"money_reduction_factor\",\"offset\":18,\"type\":\"felt\"},{\"name\":\"item_reduction_factor\",\"offset\":19,\"type\":\"felt\"},{\"name\":\"regional_item_reduction_factor\",\"offset\":20,\"type\":\"felt\"},{\"name\":\"dealer_dash_bool\",\"offset\":21,\"type\":\"felt\"},{\"name\":\"wrangle_dashed_dealer_bool\",\"offset\":22,\"type\":\"felt\"},{\"name\":\"mugging_bool\",\"offset\":23,\"type\":\"felt\"},{\"name\":\"run_from_mugging_bool\",\"offset\":24,\"type\":\"felt\"},{\"name\":\"gang_war_bool\",\"offset\":25,\"type\":\"felt\"},{\"name\":\"defend_gang_war_bool\",\"offset\":26,\"type\":\"felt\"},{\"name\":\"cop_raid_bool\",\"offset\":27,\"type\":\"felt\"},{\"name\":\"bribe_cops_bool\",\"offset\":28,\"type\":\"felt\"},{\"name\":\"find_item_bool\",\"offset\":29,\"type\":\"felt\"},{\"name\":\"local_shipment_bool\",\"offset\":30,\"type\":\"felt\"},{\"name\":\"warehouse_seizure_bool\",\"offset\":31,\"type\":\"felt\"}],\"name\":\"TurnLog\",\"size\":32,\"type\":\"struct\"},{\"inputs\":[{\"name\":\"address_of_controller\",\"type\":\"felt\"}],\"name\":\"constructor\",\"outputs\":[],\"type\":\"constructor\"},{\"inputs\":[{\"name\":\"user_id\",\"type\":\"felt\"},{\"name\":\"location_id\",\"type\":\"felt\"},{\"name\":\"buy_or_sell\",\"type\":\"felt\"},{\"name\":\"item_id\",\"type\":\"felt\"},{\"name\":\"amount_to_give\",\"type\":\"felt\"}],\"name\":\"have_turn\",\"outputs\":[],\"type\":\"function\"},{\"inputs\":[],\"name\":\"read_game_clock\",\"outputs\":[{\"name\":\"clock\",\"type\":\"felt\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"game_clock_at_turn\",\"type\":\"felt\"}],\"name\":\"view_given_turn\",\"outputs\":[{\"name\":\"turn_log\",\"type\":\"TurnLog\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"user_id\",\"type\":\"felt\"}],\"name\":\"check_user_state\",\"outputs\":[{\"name\":\"money\",\"type\":\"felt\"},{\"name\":\"id1\",\"type\":\"felt\"},{\"name\":\"id2\",\"type\":\"felt\"},{\"name\":\"id3\",\"type\":\"felt\"},{\"name\":\"id4\",\"type\":\"felt\"},{\"name\":\"id5\",\"type\":\"felt\"},{\"name\":\"id6\",\"type\":\"felt\"},{\"name\":\"id7\",\"type\":\"felt\"},{\"name\":\"id8\",\"type\":\"felt\"},{\"name\":\"id9\",\"type\":\"felt\"},{\"name\":\"id10\",\"type\":\"felt\"},{\"name\":\"id11\",\"type\":\"felt\"},{\"name\":\"id12\",\"type\":\"felt\"},{\"name\":\"id13\",\"type\":\"felt\"},{\"name\":\"id14\",\"type\":\"felt\"},{\"name\":\"id15\",\"type\":\"felt\"},{\"name\":\"id16\",\"type\":\"felt\"},{\"name\":\"id17\",\"type\":\"felt\"},{\"name\":\"id18\",\"type\":\"felt\"},{\"name\":\"id19\",\"type\":\"felt\"},{\"name\":\"location\",\"type\":\"felt\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"location_id\",\"type\":\"felt\"},{\"name\":\"item_id\",\"type\":\"felt\"}],\"name\":\"check_market_state\",\"outputs\":[{\"name\":\"item_quantity\",\"type\":\"felt\"},{\"name\":\"money_quantity\",\"type\":\"felt\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"user_combat_stats_len\",\"type\":\"felt\"},{\"name\":\"user_combat_stats\",\"type\":\"felt*\"},{\"name\":\"drug_lord_combat_stats_len\",\"type\":\"felt\"},{\"name\":\"drug_lord_combat_stats\",\"type\":\"felt*\"}],\"name\":\"challenge_current_drug_lord\",\"outputs\":[],\"type\":\"function\"}]",
}

// TypeABI is the input ABI used to generate the binding from.
// Deprecated: Use TypeMetaData.ABI instead.
var TypeABI = TypeMetaData.ABI

// Type is an auto generated Go binding around an Ethereum contract.
type Type struct {
	TypeCaller     // Read-only binding to the contract
	TypeTransactor // Write-only binding to the contract
	TypeFilterer   // Log filterer for contract events
}

// TypeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypeSession struct {
	Contract     *Type             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypeCallerSession struct {
	Contract *TypeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TypeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypeTransactorSession struct {
	Contract     *TypeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypeRaw struct {
	Contract *Type // Generic contract binding to access the raw methods on
}

// TypeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypeCallerRaw struct {
	Contract *TypeCaller // Generic read-only contract binding to access the raw methods on
}

// TypeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypeTransactorRaw struct {
	Contract *TypeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewType creates a new instance of Type, bound to a specific deployed contract.
func NewType(address common.Address, backend bind.ContractBackend) (*Type, error) {
	contract, err := bindType(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Type{TypeCaller: TypeCaller{contract: contract}, TypeTransactor: TypeTransactor{contract: contract}, TypeFilterer: TypeFilterer{contract: contract}}, nil
}

// NewTypeCaller creates a new read-only instance of Type, bound to a specific deployed contract.
func NewTypeCaller(address common.Address, caller bind.ContractCaller) (*TypeCaller, error) {
	contract, err := bindType(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypeCaller{contract: contract}, nil
}

// NewTypeTransactor creates a new write-only instance of Type, bound to a specific deployed contract.
func NewTypeTransactor(address common.Address, transactor bind.ContractTransactor) (*TypeTransactor, error) {
	contract, err := bindType(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypeTransactor{contract: contract}, nil
}

// NewTypeFilterer creates a new log filterer instance of Type, bound to a specific deployed contract.
func NewTypeFilterer(address common.Address, filterer bind.ContractFilterer) (*TypeFilterer, error) {
	contract, err := bindType(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypeFilterer{contract: contract}, nil
}

// bindType binds a generic wrapper to an already deployed contract.
func bindType(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Type *TypeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Type.Contract.TypeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Type *TypeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Type.Contract.TypeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Type *TypeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Type.Contract.TypeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Type *TypeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Type.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Type *TypeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Type.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Type *TypeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Type.Contract.contract.Transact(opts, method, params...)
}

// CheckMarketState is a free data retrieval call binding the contract method 0x3990913b.
//
// Solidity: function check_market_state(felt location_id, felt item_id) view returns(felt item_quantity, felt money_quantity)
func (_Type *TypeCaller) CheckMarketState(opts *bind.CallOpts, location_id *big.Int, item_id *big.Int) (struct {
	ItemQuantity  *big.Int
	MoneyQuantity *big.Int
}, error) {
	var out []interface{}
	err := _Type.contract.Call(opts, &out, "check_market_state", location_id, item_id)

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

// CheckMarketState is a free data retrieval call binding the contract method 0x3990913b.
//
// Solidity: function check_market_state(felt location_id, felt item_id) view returns(felt item_quantity, felt money_quantity)
func (_Type *TypeSession) CheckMarketState(location_id *big.Int, item_id *big.Int) (struct {
	ItemQuantity  *big.Int
	MoneyQuantity *big.Int
}, error) {
	return _Type.Contract.CheckMarketState(&_Type.CallOpts, location_id, item_id)
}

// CheckMarketState is a free data retrieval call binding the contract method 0x3990913b.
//
// Solidity: function check_market_state(felt location_id, felt item_id) view returns(felt item_quantity, felt money_quantity)
func (_Type *TypeCallerSession) CheckMarketState(location_id *big.Int, item_id *big.Int) (struct {
	ItemQuantity  *big.Int
	MoneyQuantity *big.Int
}, error) {
	return _Type.Contract.CheckMarketState(&_Type.CallOpts, location_id, item_id)
}

// CheckUserState is a free data retrieval call binding the contract method 0x0ec6e44f.
//
// Solidity: function check_user_state(felt user_id) view returns(felt money, felt id1, felt id2, felt id3, felt id4, felt id5, felt id6, felt id7, felt id8, felt id9, felt id10, felt id11, felt id12, felt id13, felt id14, felt id15, felt id16, felt id17, felt id18, felt id19, felt location)
func (_Type *TypeCaller) CheckUserState(opts *bind.CallOpts, user_id *big.Int) (struct {
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
	err := _Type.contract.Call(opts, &out, "check_user_state", user_id)

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

// CheckUserState is a free data retrieval call binding the contract method 0x0ec6e44f.
//
// Solidity: function check_user_state(felt user_id) view returns(felt money, felt id1, felt id2, felt id3, felt id4, felt id5, felt id6, felt id7, felt id8, felt id9, felt id10, felt id11, felt id12, felt id13, felt id14, felt id15, felt id16, felt id17, felt id18, felt id19, felt location)
func (_Type *TypeSession) CheckUserState(user_id *big.Int) (struct {
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
	return _Type.Contract.CheckUserState(&_Type.CallOpts, user_id)
}

// CheckUserState is a free data retrieval call binding the contract method 0x0ec6e44f.
//
// Solidity: function check_user_state(felt user_id) view returns(felt money, felt id1, felt id2, felt id3, felt id4, felt id5, felt id6, felt id7, felt id8, felt id9, felt id10, felt id11, felt id12, felt id13, felt id14, felt id15, felt id16, felt id17, felt id18, felt id19, felt location)
func (_Type *TypeCallerSession) CheckUserState(user_id *big.Int) (struct {
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
	return _Type.Contract.CheckUserState(&_Type.CallOpts, user_id)
}

// ReadGameClock is a free data retrieval call binding the contract method 0x196d94ee.
//
// Solidity: function read_game_clock() view returns(felt clock)
func (_Type *TypeCaller) ReadGameClock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Type.contract.Call(opts, &out, "read_game_clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadGameClock is a free data retrieval call binding the contract method 0x196d94ee.
//
// Solidity: function read_game_clock() view returns(felt clock)
func (_Type *TypeSession) ReadGameClock() (*big.Int, error) {
	return _Type.Contract.ReadGameClock(&_Type.CallOpts)
}

// ReadGameClock is a free data retrieval call binding the contract method 0x196d94ee.
//
// Solidity: function read_game_clock() view returns(felt clock)
func (_Type *TypeCallerSession) ReadGameClock() (*big.Int, error) {
	return _Type.Contract.ReadGameClock(&_Type.CallOpts)
}

// ViewGivenTurn is a free data retrieval call binding the contract method 0xdc847c54.
//
// Solidity: function view_given_turn(felt game_clock_at_turn) view returns(TurnLog turn_log)
func (_Type *TypeCaller) ViewGivenTurn(opts *bind.CallOpts, game_clock_at_turn *big.Int) (TurnLog, error) {
	var out []interface{}
	err := _Type.contract.Call(opts, &out, "view_given_turn", game_clock_at_turn)

	if err != nil {
		return *new(TurnLog), err
	}

	out0 := *abi.ConvertType(out[0], new(TurnLog)).(*TurnLog)

	return out0, err

}

// ViewGivenTurn is a free data retrieval call binding the contract method 0xdc847c54.
//
// Solidity: function view_given_turn(felt game_clock_at_turn) view returns(TurnLog turn_log)
func (_Type *TypeSession) ViewGivenTurn(game_clock_at_turn *big.Int) (TurnLog, error) {
	return _Type.Contract.ViewGivenTurn(&_Type.CallOpts, game_clock_at_turn)
}

// ViewGivenTurn is a free data retrieval call binding the contract method 0xdc847c54.
//
// Solidity: function view_given_turn(felt game_clock_at_turn) view returns(TurnLog turn_log)
func (_Type *TypeCallerSession) ViewGivenTurn(game_clock_at_turn *big.Int) (TurnLog, error) {
	return _Type.Contract.ViewGivenTurn(&_Type.CallOpts, game_clock_at_turn)
}

// ChallengeCurrentDrugLord is a paid mutator transaction binding the contract method 0xe9dcb95c.
//
// Solidity: function challenge_current_drug_lord(felt user_combat_stats_len, felt* user_combat_stats, felt drug_lord_combat_stats_len, felt* drug_lord_combat_stats) returns()
func (_Type *TypeTransactor) ChallengeCurrentDrugLord(opts *bind.TransactOpts, user_combat_stats_len *big.Int, user_combat_stats *big.Int, drug_lord_combat_stats_len *big.Int, drug_lord_combat_stats *big.Int) (*types.Transaction, error) {
	return _Type.contract.Transact(opts, "challenge_current_drug_lord", user_combat_stats_len, user_combat_stats, drug_lord_combat_stats_len, drug_lord_combat_stats)
}

// ChallengeCurrentDrugLord is a paid mutator transaction binding the contract method 0xe9dcb95c.
//
// Solidity: function challenge_current_drug_lord(felt user_combat_stats_len, felt* user_combat_stats, felt drug_lord_combat_stats_len, felt* drug_lord_combat_stats) returns()
func (_Type *TypeSession) ChallengeCurrentDrugLord(user_combat_stats_len *big.Int, user_combat_stats *big.Int, drug_lord_combat_stats_len *big.Int, drug_lord_combat_stats *big.Int) (*types.Transaction, error) {
	return _Type.Contract.ChallengeCurrentDrugLord(&_Type.TransactOpts, user_combat_stats_len, user_combat_stats, drug_lord_combat_stats_len, drug_lord_combat_stats)
}

// ChallengeCurrentDrugLord is a paid mutator transaction binding the contract method 0xe9dcb95c.
//
// Solidity: function challenge_current_drug_lord(felt user_combat_stats_len, felt* user_combat_stats, felt drug_lord_combat_stats_len, felt* drug_lord_combat_stats) returns()
func (_Type *TypeTransactorSession) ChallengeCurrentDrugLord(user_combat_stats_len *big.Int, user_combat_stats *big.Int, drug_lord_combat_stats_len *big.Int, drug_lord_combat_stats *big.Int) (*types.Transaction, error) {
	return _Type.Contract.ChallengeCurrentDrugLord(&_Type.TransactOpts, user_combat_stats_len, user_combat_stats, drug_lord_combat_stats_len, drug_lord_combat_stats)
}

// HaveTurn is a paid mutator transaction binding the contract method 0x45213ffc.
//
// Solidity: function have_turn(felt user_id, felt location_id, felt buy_or_sell, felt item_id, felt amount_to_give) returns()
func (_Type *TypeTransactor) HaveTurn(opts *bind.TransactOpts, user_id *big.Int, location_id *big.Int, buy_or_sell *big.Int, item_id *big.Int, amount_to_give *big.Int) (*types.Transaction, error) {
	return _Type.contract.Transact(opts, "have_turn", user_id, location_id, buy_or_sell, item_id, amount_to_give)
}

// HaveTurn is a paid mutator transaction binding the contract method 0x45213ffc.
//
// Solidity: function have_turn(felt user_id, felt location_id, felt buy_or_sell, felt item_id, felt amount_to_give) returns()
func (_Type *TypeSession) HaveTurn(user_id *big.Int, location_id *big.Int, buy_or_sell *big.Int, item_id *big.Int, amount_to_give *big.Int) (*types.Transaction, error) {
	return _Type.Contract.HaveTurn(&_Type.TransactOpts, user_id, location_id, buy_or_sell, item_id, amount_to_give)
}

// HaveTurn is a paid mutator transaction binding the contract method 0x45213ffc.
//
// Solidity: function have_turn(felt user_id, felt location_id, felt buy_or_sell, felt item_id, felt amount_to_give) returns()
func (_Type *TypeTransactorSession) HaveTurn(user_id *big.Int, location_id *big.Int, buy_or_sell *big.Int, item_id *big.Int, amount_to_give *big.Int) (*types.Transaction, error) {
	return _Type.Contract.HaveTurn(&_Type.TransactOpts, user_id, location_id, buy_or_sell, item_id, amount_to_give)
}
