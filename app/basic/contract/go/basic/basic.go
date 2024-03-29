// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basic

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
	_ = abi.ConvertType
)

// BasicMetaData contains all meta data concerning the Basic contract.
var BasicMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"Items\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040518060400160405280600381526020017f312e3100000000000000000000000000000000000000000000000000000000008152506000908161005591906102ab565b5061037d565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806100dc57607f821691505b6020821081036100ef576100ee610095565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026101577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261011a565b610161868361011a565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006101a86101a361019e84610179565b610183565b610179565b9050919050565b6000819050919050565b6101c28361018d565b6101d66101ce826101af565b848454610127565b825550505050565b600090565b6101eb6101de565b6101f68184846101b9565b505050565b5b8181101561021a5761020f6000826101e3565b6001810190506101fc565b5050565b601f82111561025f57610230816100f5565b6102398461010a565b81016020851015610248578190505b61025c6102548561010a565b8301826101fb565b50505b505050565b600082821c905092915050565b600061028260001984600802610264565b1980831691505092915050565b600061029b8383610271565b9150826002028217905092915050565b6102b48261005b565b67ffffffffffffffff8111156102cd576102cc610066565b5b6102d782546100c4565b6102e282828561021e565b600060209050601f8311600181146103155760008415610303578287015190505b61030d858261028f565b865550610375565b601f198416610323866100f5565b60005b8281101561034b57848901518255600182019150602085019450602081019050610326565b868310156103685784890151610364601f891682610271565b8355505b6001600288020188555050505b505050505050565b6105e58061038c6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80634547a6b3146100465780638c5cf3ed14610076578063bb62860d14610092575b600080fd5b610060600480360381019061005b9190610326565b6100b0565b60405161006d9190610388565b60405180910390f35b610090600480360381019061008b91906103cf565b6100de565b005b61009a61013e565b6040516100a791906104aa565b60405180910390f35b6001818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b806001836040516100ef9190610508565b9081526020016040518091039020819055507f814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f828260405161013292919061051f565b60405180910390a15050565b6000805461014b9061057e565b80601f01602080910402602001604051908101604052809291908181526020018280546101779061057e565b80156101c45780601f10610199576101008083540402835291602001916101c4565b820191906000526020600020905b8154815290600101906020018083116101a757829003601f168201915b505050505081565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610233826101ea565b810181811067ffffffffffffffff82111715610252576102516101fb565b5b80604052505050565b60006102656101cc565b9050610271828261022a565b919050565b600067ffffffffffffffff821115610291576102906101fb565b5b61029a826101ea565b9050602081019050919050565b82818337600083830152505050565b60006102c96102c484610276565b61025b565b9050828152602081018484840111156102e5576102e46101e5565b5b6102f08482856102a7565b509392505050565b600082601f83011261030d5761030c6101e0565b5b813561031d8482602086016102b6565b91505092915050565b60006020828403121561033c5761033b6101d6565b5b600082013567ffffffffffffffff81111561035a576103596101db565b5b610366848285016102f8565b91505092915050565b6000819050919050565b6103828161036f565b82525050565b600060208201905061039d6000830184610379565b92915050565b6103ac8161036f565b81146103b757600080fd5b50565b6000813590506103c9816103a3565b92915050565b600080604083850312156103e6576103e56101d6565b5b600083013567ffffffffffffffff811115610404576104036101db565b5b610410858286016102f8565b9250506020610421858286016103ba565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561046557808201518184015260208101905061044a565b60008484015250505050565b600061047c8261042b565b6104868185610436565b9350610496818560208601610447565b61049f816101ea565b840191505092915050565b600060208201905081810360008301526104c48184610471565b905092915050565b600081905092915050565b60006104e28261042b565b6104ec81856104cc565b93506104fc818560208601610447565b80840191505092915050565b600061051482846104d7565b915081905092915050565b600060408201905081810360008301526105398185610471565b90506105486020830184610379565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061059657607f821691505b6020821081036105a9576105a861054f565b5b5091905056fea2646970667358221220cf1af4405ebb15e857a546469fc78478069e13944272bf821157b7365245c1ec64736f6c63430008190033",
}

// BasicABI is the input ABI used to generate the binding from.
// Deprecated: Use BasicMetaData.ABI instead.
var BasicABI = BasicMetaData.ABI

// BasicBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BasicMetaData.Bin instead.
var BasicBin = BasicMetaData.Bin

// DeployBasic deploys a new Ethereum contract, binding an instance of Basic to it.
func DeployBasic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Basic, error) {
	parsed, err := BasicMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Basic{BasicCaller: BasicCaller{contract: contract}, BasicTransactor: BasicTransactor{contract: contract}, BasicFilterer: BasicFilterer{contract: contract}}, nil
}

// Basic is an auto generated Go binding around an Ethereum contract.
type Basic struct {
	BasicCaller     // Read-only binding to the contract
	BasicTransactor // Write-only binding to the contract
	BasicFilterer   // Log filterer for contract events
}

// BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicSession struct {
	Contract     *Basic            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicCallerSession struct {
	Contract *BasicCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTransactorSession struct {
	Contract     *BasicTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicRaw struct {
	Contract *Basic // Generic contract binding to access the raw methods on
}

// BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicCallerRaw struct {
	Contract *BasicCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTransactorRaw struct {
	Contract *BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasic creates a new instance of Basic, bound to a specific deployed contract.
func NewBasic(address common.Address, backend bind.ContractBackend) (*Basic, error) {
	contract, err := bindBasic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Basic{BasicCaller: BasicCaller{contract: contract}, BasicTransactor: BasicTransactor{contract: contract}, BasicFilterer: BasicFilterer{contract: contract}}, nil
}

// NewBasicCaller creates a new read-only instance of Basic, bound to a specific deployed contract.
func NewBasicCaller(address common.Address, caller bind.ContractCaller) (*BasicCaller, error) {
	contract, err := bindBasic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicCaller{contract: contract}, nil
}

// NewBasicTransactor creates a new write-only instance of Basic, bound to a specific deployed contract.
func NewBasicTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTransactor, error) {
	contract, err := bindBasic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTransactor{contract: contract}, nil
}

// NewBasicFilterer creates a new log filterer instance of Basic, bound to a specific deployed contract.
func NewBasicFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicFilterer, error) {
	contract, err := bindBasic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicFilterer{contract: contract}, nil
}

// bindBasic binds a generic wrapper to an already deployed contract.
func bindBasic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasicMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Basic *BasicRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Basic.Contract.BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Basic *BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basic.Contract.BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Basic *BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Basic.Contract.BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Basic *BasicCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Basic *BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Basic *BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Basic.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x4547a6b3.
//
// Solidity: function Items(string ) view returns(uint256)
func (_Basic *BasicCaller) Items(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "Items", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x4547a6b3.
//
// Solidity: function Items(string ) view returns(uint256)
func (_Basic *BasicSession) Items(arg0 string) (*big.Int, error) {
	return _Basic.Contract.Items(&_Basic.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x4547a6b3.
//
// Solidity: function Items(string ) view returns(uint256)
func (_Basic *BasicCallerSession) Items(arg0 string) (*big.Int, error) {
	return _Basic.Contract.Items(&_Basic.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Basic *BasicCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Basic.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Basic *BasicSession) Version() (string, error) {
	return _Basic.Contract.Version(&_Basic.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Basic *BasicCallerSession) Version() (string, error) {
	return _Basic.Contract.Version(&_Basic.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0x8c5cf3ed.
//
// Solidity: function SetItem(string key, uint256 value) returns()
func (_Basic *BasicTransactor) SetItem(opts *bind.TransactOpts, key string, value *big.Int) (*types.Transaction, error) {
	return _Basic.contract.Transact(opts, "SetItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0x8c5cf3ed.
//
// Solidity: function SetItem(string key, uint256 value) returns()
func (_Basic *BasicSession) SetItem(key string, value *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.SetItem(&_Basic.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0x8c5cf3ed.
//
// Solidity: function SetItem(string key, uint256 value) returns()
func (_Basic *BasicTransactorSession) SetItem(key string, value *big.Int) (*types.Transaction, error) {
	return _Basic.Contract.SetItem(&_Basic.TransactOpts, key, value)
}

// BasicItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Basic contract.
type BasicItemSetIterator struct {
	Event *BasicItemSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicItemSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicItemSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicItemSet represents a ItemSet event raised by the Basic contract.
type BasicItemSet struct {
	Key   string
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Basic *BasicFilterer) FilterItemSet(opts *bind.FilterOpts) (*BasicItemSetIterator, error) {

	logs, sub, err := _Basic.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &BasicItemSetIterator{contract: _Basic.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Basic *BasicFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *BasicItemSet) (event.Subscription, error) {

	logs, sub, err := _Basic.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicItemSet)
				if err := _Basic.contract.UnpackLog(event, "ItemSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseItemSet is a log parse operation binding the contract event 0x814c96094cf9633fd519eab4539bc5f26a8ab7965b7243057596bafd3318e60f.
//
// Solidity: event ItemSet(string key, uint256 value)
func (_Basic *BasicFilterer) ParseItemSet(log types.Log) (*BasicItemSet, error) {
	event := new(BasicItemSet)
	if err := _Basic.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
