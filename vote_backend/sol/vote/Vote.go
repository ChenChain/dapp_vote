// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vote

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

// VoteCandidate is an auto generated low-level Go binding around an user-defined struct.
type VoteCandidate struct {
	Name      string
	VoteCount *big.Int
}

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"candidateIndex\",\"type\":\"uint256\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"addCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"internalType\":\"structVote.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_candidateIndex\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610fa38061005b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80630121b93f1461006457806306a49fce1461008057806309eef43e1461009e5780633477ee2e146100ce578063462e91ec146100ff5780638da5cb5b1461011b575b5f5ffd5b61007e60048036038101906100799190610606565b610139565b005b6100886102e9565b60405161009591906107a5565b60405180910390f35b6100b860048036038101906100b3919061081f565b6103de565b6040516100c59190610864565b60405180910390f35b6100e860048036038101906100e39190610606565b6103fb565b6040516100f69291906108d4565b60405180910390f35b61011960048036038101906101149190610a2e565b6104b0565b005b61012361059e565b6040516101309190610a84565b60405180910390f35b60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156101c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ba90610ae7565b60405180910390fd5b600180549050811061020a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020190610b4f565b60405180910390fd5b600160025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600180828154811061027457610273610b6d565b5b905f5260205f2090600202016001015f8282546102919190610bc7565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b826040516102de9190610bfa565b60405180910390a250565b60606001805480602002602001604051908101604052809291908181526020015f905b828210156103d5578382905f5260205f2090600202016040518060400160405290815f8201805461033c90610c40565b80601f016020809104026020016040519081016040528092919081815260200182805461036890610c40565b80156103b35780601f1061038a576101008083540402835291602001916103b3565b820191905f5260205f20905b81548152906001019060200180831161039657829003601f168201915b505050505081526020016001820154815250508152602001906001019061030c565b50505050905090565b6002602052805f5260405f205f915054906101000a900460ff1681565b6001818154811061040a575f80fd5b905f5260205f2090600202015f91509050805f01805461042990610c40565b80601f016020809104026020016040519081016040528092919081815260200182805461045590610c40565b80156104a05780601f10610477576101008083540402835291602001916104a0565b820191905f5260205f20905b81548152906001019060200180831161048357829003601f168201915b5050505050908060010154905082565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461053e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053590610ce0565b60405180910390fd5b600160405180604001604052808381526020015f815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01908161058e9190610e9e565b5060208201518160010155505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b6105e5816105d3565b81146105ef575f5ffd5b50565b5f81359050610600816105dc565b92915050565b5f6020828403121561061b5761061a6105cb565b5b5f610628848285016105f2565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61069c8261065a565b6106a68185610664565b93506106b6818560208601610674565b6106bf81610682565b840191505092915050565b6106d3816105d3565b82525050565b5f604083015f8301518482035f8601526106f38282610692565b915050602083015161070860208601826106ca565b508091505092915050565b5f61071e83836106d9565b905092915050565b5f602082019050919050565b5f61073c82610631565b610746818561063b565b9350836020820285016107588561064b565b805f5b8581101561079357848403895281516107748582610713565b945061077f83610726565b925060208a0199505060018101905061075b565b50829750879550505050505092915050565b5f6020820190508181035f8301526107bd8184610732565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107ee826107c5565b9050919050565b6107fe816107e4565b8114610808575f5ffd5b50565b5f81359050610819816107f5565b92915050565b5f60208284031215610834576108336105cb565b5b5f6108418482850161080b565b91505092915050565b5f8115159050919050565b61085e8161084a565b82525050565b5f6020820190506108775f830184610855565b92915050565b5f82825260208201905092915050565b5f6108978261065a565b6108a1818561087d565b93506108b1818560208601610674565b6108ba81610682565b840191505092915050565b6108ce816105d3565b82525050565b5f6040820190508181035f8301526108ec818561088d565b90506108fb60208301846108c5565b9392505050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61094082610682565b810181811067ffffffffffffffff8211171561095f5761095e61090a565b5b80604052505050565b5f6109716105c2565b905061097d8282610937565b919050565b5f67ffffffffffffffff82111561099c5761099b61090a565b5b6109a582610682565b9050602081019050919050565b828183375f83830152505050565b5f6109d26109cd84610982565b610968565b9050828152602081018484840111156109ee576109ed610906565b5b6109f98482856109b2565b509392505050565b5f82601f830112610a1557610a14610902565b5b8135610a258482602086016109c0565b91505092915050565b5f60208284031215610a4357610a426105cb565b5b5f82013567ffffffffffffffff811115610a6057610a5f6105cf565b5b610a6c84828501610a01565b91505092915050565b610a7e816107e4565b82525050565b5f602082019050610a975f830184610a75565b92915050565b7f596f75206861766520616c726561647920766f746564000000000000000000005f82015250565b5f610ad160168361087d565b9150610adc82610a9d565b602082019050919050565b5f6020820190508181035f830152610afe81610ac5565b9050919050565b7f496e76616c69642063616e64696461746520696e6465780000000000000000005f82015250565b5f610b3960178361087d565b9150610b4482610b05565b602082019050919050565b5f6020820190508181035f830152610b6681610b2d565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610bd1826105d3565b9150610bdc836105d3565b9250828201905080821115610bf457610bf3610b9a565b5b92915050565b5f602082019050610c0d5f8301846108c5565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610c5757607f821691505b602082108103610c6a57610c69610c13565b5b50919050565b7f4f6e6c7920746865206f776e65722063616e20706572666f726d2074686973205f8201527f616374696f6e0000000000000000000000000000000000000000000000000000602082015250565b5f610cca60268361087d565b9150610cd582610c70565b604082019050919050565b5f6020820190508181035f830152610cf781610cbe565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610d5a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d1f565b610d648683610d1f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610d9f610d9a610d95846105d3565b610d7c565b6105d3565b9050919050565b5f819050919050565b610db883610d85565b610dcc610dc482610da6565b848454610d2b565b825550505050565b5f5f905090565b610de3610dd4565b610dee818484610daf565b505050565b5b81811015610e1157610e065f82610ddb565b600181019050610df4565b5050565b601f821115610e5657610e2781610cfe565b610e3084610d10565b81016020851015610e3f578190505b610e53610e4b85610d10565b830182610df3565b50505b505050565b5f82821c905092915050565b5f610e765f1984600802610e5b565b1980831691505092915050565b5f610e8e8383610e67565b9150826002028217905092915050565b610ea78261065a565b67ffffffffffffffff811115610ec057610ebf61090a565b5b610eca8254610c40565b610ed5828285610e15565b5f60209050601f831160018114610f06575f8415610ef4578287015190505b610efe8582610e83565b865550610f65565b601f198416610f1486610cfe565b5f5b82811015610f3b57848901518255600182019150602085019450602081019050610f16565b86831015610f585784890151610f54601f891682610e67565b8355505b6001600288020188555050505b50505050505056fea264697066735822122062816f02b9094b0c375df979f66c2c72a5bac76315cd052337d4ad7f48f6677d64736f6c634300081c0033",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// VoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoteMetaData.Bin instead.
var VoteBin = VoteMetaData.Bin

// DeployVote deploys a new Ethereum contract, binding an instance of Vote to it.
func DeployVote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vote, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount)
func (_Vote *VoteCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount)
func (_Vote *VoteSession) Candidates(arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	return _Vote.Contract.Candidates(&_Vote.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount)
func (_Vote *VoteCallerSession) Candidates(arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	return _Vote.Contract.Candidates(&_Vote.CallOpts, arg0)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() view returns((string,uint256)[])
func (_Vote *VoteCaller) GetCandidates(opts *bind.CallOpts) ([]VoteCandidate, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getCandidates")

	if err != nil {
		return *new([]VoteCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]VoteCandidate)).(*[]VoteCandidate)

	return out0, err

}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() view returns((string,uint256)[])
func (_Vote *VoteSession) GetCandidates() ([]VoteCandidate, error) {
	return _Vote.Contract.GetCandidates(&_Vote.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() view returns((string,uint256)[])
func (_Vote *VoteCallerSession) GetCandidates() ([]VoteCandidate, error) {
	return _Vote.Contract.GetCandidates(&_Vote.CallOpts)
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Vote *VoteCaller) HasVoted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "hasVoted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Vote *VoteSession) HasVoted(arg0 common.Address) (bool, error) {
	return _Vote.Contract.HasVoted(&_Vote.CallOpts, arg0)
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Vote *VoteCallerSession) HasVoted(arg0 common.Address) (bool, error) {
	return _Vote.Contract.HasVoted(&_Vote.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCallerSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Vote *VoteTransactor) AddCandidate(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "addCandidate", _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Vote *VoteSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Vote.Contract.AddCandidate(&_Vote.TransactOpts, _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Vote *VoteTransactorSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Vote.Contract.AddCandidate(&_Vote.TransactOpts, _name)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateIndex) returns()
func (_Vote *VoteTransactor) Vote(opts *bind.TransactOpts, _candidateIndex *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "vote", _candidateIndex)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateIndex) returns()
func (_Vote *VoteSession) Vote(_candidateIndex *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, _candidateIndex)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateIndex) returns()
func (_Vote *VoteTransactorSession) Vote(_candidateIndex *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, _candidateIndex)
}

// VoteVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Vote contract.
type VoteVoteCastIterator struct {
	Event *VoteVoteCast // Event containing the contract specifics and raw log

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
func (it *VoteVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteVoteCast)
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
		it.Event = new(VoteVoteCast)
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
func (it *VoteVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteVoteCast represents a VoteCast event raised by the Vote contract.
type VoteVoteCast struct {
	Voter          common.Address
	CandidateIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address indexed voter, uint256 candidateIndex)
func (_Vote *VoteFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*VoteVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &VoteVoteCastIterator{contract: _Vote.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address indexed voter, uint256 candidateIndex)
func (_Vote *VoteFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *VoteVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteVoteCast)
				if err := _Vote.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address indexed voter, uint256 candidateIndex)
func (_Vote *VoteFilterer) ParseVoteCast(log types.Log) (*VoteVoteCast, error) {
	event := new(VoteVoteCast)
	if err := _Vote.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
