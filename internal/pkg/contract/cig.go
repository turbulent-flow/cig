// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// CIGBrand is an auto generated low-level Go binding around an user-defined struct.
type CIGBrand struct {
	Id          [32]byte
	Name        string
	Description string
	CreatedAt   *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedBrand\",\"type\":\"bytes\"}],\"name\":\"addBrand\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"brandID\",\"type\":\"bytes32\"}],\"name\":\"getBrand\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structCIG.Brand\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50610d888061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633f9bbdf91461003b57806398c5ee5014610057575b600080fd5b6100556004803603810190610050919061055b565b610087565b005b610071600480360381019061006c91906105da565b6101d1565b60405161007e9190610718565b60405180910390f35b60008180602001905181019061009d91906108d6565b905060008019168160000151036100e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100e09061097c565b60405180910390fd5b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081602001518051906020012003610156576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014d906109e8565b60405180910390fd5b600180816001815401808255809150500390600052602060002090505080600080836000015181526020019081526020016000206000820151816000015560208201518160010190816101a99190610c14565b5060408201518160020190816101bf9190610c14565b50606082015181600301559050505050565b6101d96103d6565b6000801916820361021f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102169061097c565b60405180910390fd5b6000801916600080848152602001908152602001600020600001540361027a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027190610d32565b60405180910390fd5b600080838152602001908152602001600020604051806080016040529081600082015481526020016001820180546102b190610a37565b80601f01602080910402602001604051908101604052809291908181526020018280546102dd90610a37565b801561032a5780601f106102ff5761010080835404028352916020019161032a565b820191906000526020600020905b81548152906001019060200180831161030d57829003601f168201915b5050505050815260200160028201805461034390610a37565b80601f016020809104026020016040519081016040528092919081815260200182805461036f90610a37565b80156103bc5780601f10610391576101008083540402835291602001916103bc565b820191906000526020600020905b81548152906001019060200180831161039f57829003601f168201915b505050505081526020016003820154815250509050919050565b6040518060800160405280600080191681526020016060815260200160608152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104688261041f565b810181811067ffffffffffffffff8211171561048757610486610430565b5b80604052505050565b600061049a610401565b90506104a6828261045f565b919050565b600067ffffffffffffffff8211156104c6576104c5610430565b5b6104cf8261041f565b9050602081019050919050565b82818337600083830152505050565b60006104fe6104f9846104ab565b610490565b90508281526020810184848401111561051a5761051961041a565b5b6105258482856104dc565b509392505050565b600082601f83011261054257610541610415565b5b81356105528482602086016104eb565b91505092915050565b6000602082840312156105715761057061040b565b5b600082013567ffffffffffffffff81111561058f5761058e610410565b5b61059b8482850161052d565b91505092915050565b6000819050919050565b6105b7816105a4565b81146105c257600080fd5b50565b6000813590506105d4816105ae565b92915050565b6000602082840312156105f0576105ef61040b565b5b60006105fe848285016105c5565b91505092915050565b610610816105a4565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610650578082015181840152602081019050610635565b60008484015250505050565b600061066782610616565b6106718185610621565b9350610681818560208601610632565b61068a8161041f565b840191505092915050565b6000819050919050565b6106a881610695565b82525050565b60006080830160008301516106c66000860182610607565b50602083015184820360208601526106de828261065c565b915050604083015184820360408601526106f8828261065c565b915050606083015161070d606086018261069f565b508091505092915050565b6000602082019050818103600083015261073281846106ae565b905092915050565b600080fd5b600080fd5b600081519050610753816105ae565b92915050565b600067ffffffffffffffff82111561077457610773610430565b5b61077d8261041f565b9050602081019050919050565b600061079d61079884610759565b610490565b9050828152602081018484840111156107b9576107b861041a565b5b6107c4848285610632565b509392505050565b600082601f8301126107e1576107e0610415565b5b81516107f184826020860161078a565b91505092915050565b61080381610695565b811461080e57600080fd5b50565b600081519050610820816107fa565b92915050565b60006080828403121561083c5761083b61073a565b5b6108466080610490565b9050600061085684828501610744565b600083015250602082015167ffffffffffffffff81111561087a5761087961073f565b5b610886848285016107cc565b602083015250604082015167ffffffffffffffff8111156108aa576108a961073f565b5b6108b6848285016107cc565b60408301525060606108ca84828501610811565b60608301525092915050565b6000602082840312156108ec576108eb61040b565b5b600082015167ffffffffffffffff81111561090a57610909610410565b5b61091684828501610826565b91505092915050565b600082825260208201905092915050565b7f4272616e6420494420697320696e76616c696400000000000000000000000000600082015250565b600061096660138361091f565b915061097182610930565b602082019050919050565b6000602082019050818103600083015261099581610959565b9050919050565b7f4272616e64206e616d6520697320696e76616c69640000000000000000000000600082015250565b60006109d260158361091f565b91506109dd8261099c565b602082019050919050565b60006020820190508181036000830152610a01816109c5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610a4f57607f821691505b602082108103610a6257610a61610a08565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610aca7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a8d565b610ad48683610a8d565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610b11610b0c610b0784610695565b610aec565b610695565b9050919050565b6000819050919050565b610b2b83610af6565b610b3f610b3782610b18565b848454610a9a565b825550505050565b600090565b610b54610b47565b610b5f818484610b22565b505050565b5b81811015610b8357610b78600082610b4c565b600181019050610b65565b5050565b601f821115610bc857610b9981610a68565b610ba284610a7d565b81016020851015610bb1578190505b610bc5610bbd85610a7d565b830182610b64565b50505b505050565b600082821c905092915050565b6000610beb60001984600802610bcd565b1980831691505092915050565b6000610c048383610bda565b9150826002028217905092915050565b610c1d82610616565b67ffffffffffffffff811115610c3657610c35610430565b5b610c408254610a37565b610c4b828285610b87565b600060209050601f831160018114610c7e5760008415610c6c578287015190505b610c768582610bf8565b865550610cde565b601f198416610c8c86610a68565b60005b82811015610cb457848901518255600182019150602085019450602081019050610c8f565b86831015610cd15784890151610ccd601f891682610bda565b8355505b6001600288020188555050505b505050505050565b7f4272616e6420646f6573206e6f74206578697374000000000000000000000000600082015250565b6000610d1c60148361091f565b9150610d2782610ce6565b602082019050919050565b60006020820190508181036000830152610d4b81610d0f565b905091905056fea264697066735822122042004c843abf59b107489e6ba0636458e5f43f79f9266b92fcbda90bbb1ccd6d64736f6c634300081d0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetBrand is a free data retrieval call binding the contract method 0x98c5ee50.
//
// Solidity: function getBrand(bytes32 brandID) view returns((bytes32,string,string,uint256))
func (_Contract *ContractCaller) GetBrand(opts *bind.CallOpts, brandID [32]byte) (CIGBrand, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBrand", brandID)

	if err != nil {
		return *new(CIGBrand), err
	}

	out0 := *abi.ConvertType(out[0], new(CIGBrand)).(*CIGBrand)

	return out0, err

}

// GetBrand is a free data retrieval call binding the contract method 0x98c5ee50.
//
// Solidity: function getBrand(bytes32 brandID) view returns((bytes32,string,string,uint256))
func (_Contract *ContractSession) GetBrand(brandID [32]byte) (CIGBrand, error) {
	return _Contract.Contract.GetBrand(&_Contract.CallOpts, brandID)
}

// GetBrand is a free data retrieval call binding the contract method 0x98c5ee50.
//
// Solidity: function getBrand(bytes32 brandID) view returns((bytes32,string,string,uint256))
func (_Contract *ContractCallerSession) GetBrand(brandID [32]byte) (CIGBrand, error) {
	return _Contract.Contract.GetBrand(&_Contract.CallOpts, brandID)
}

// AddBrand is a paid mutator transaction binding the contract method 0x3f9bbdf9.
//
// Solidity: function addBrand(bytes encodedBrand) returns()
func (_Contract *ContractTransactor) AddBrand(opts *bind.TransactOpts, encodedBrand []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addBrand", encodedBrand)
}

// AddBrand is a paid mutator transaction binding the contract method 0x3f9bbdf9.
//
// Solidity: function addBrand(bytes encodedBrand) returns()
func (_Contract *ContractSession) AddBrand(encodedBrand []byte) (*types.Transaction, error) {
	return _Contract.Contract.AddBrand(&_Contract.TransactOpts, encodedBrand)
}

// AddBrand is a paid mutator transaction binding the contract method 0x3f9bbdf9.
//
// Solidity: function addBrand(bytes encodedBrand) returns()
func (_Contract *ContractTransactorSession) AddBrand(encodedBrand []byte) (*types.Transaction, error) {
	return _Contract.Contract.AddBrand(&_Contract.TransactOpts, encodedBrand)
}
