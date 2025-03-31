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
	Bin: "0x6080604052348015600e575f5ffd5b50610cfc8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80633f9bbdf91461003857806398c5ee5014610054575b5f5ffd5b610052600480360381019061004d9190610527565b610084565b005b61006e600480360381019061006991906105a1565b6101c1565b60405161007b91906106ba565b60405180910390f35b5f81806020019051810190610099919061086b565b90505f815f0151036100e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100d79061090c565b60405180910390fd5b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708160200151805190602001200361014d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014490610974565b60405180910390fd5b6001808160018154018082558091505003905f5260205f20905050805f5f835f015181526020019081526020015f205f820151815f015560208201518160010190816101999190610b8f565b5060408201518160020190816101af9190610b8f565b50606082015181600301559050505050565b6101c96103b4565b5f820361020b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102029061090c565b60405180910390fd5b5f5f5f8481526020019081526020015f205f01540361025f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025690610ca8565b60405180910390fd5b5f5f8381526020019081526020015f206040518060800160405290815f8201548152602001600182018054610293906109bf565b80601f01602080910402602001604051908101604052809291908181526020018280546102bf906109bf565b801561030a5780601f106102e15761010080835404028352916020019161030a565b820191905f5260205f20905b8154815290600101906020018083116102ed57829003601f168201915b50505050508152602001600282018054610323906109bf565b80601f016020809104026020016040519081016040528092919081815260200182805461034f906109bf565b801561039a5780601f106103715761010080835404028352916020019161039a565b820191905f5260205f20905b81548152906001019060200180831161037d57829003601f168201915b505050505081526020016003820154815250509050919050565b60405180608001604052805f815260200160608152602001606081526020015f81525090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610439826103f3565b810181811067ffffffffffffffff8211171561045857610457610403565b5b80604052505050565b5f61046a6103da565b90506104768282610430565b919050565b5f67ffffffffffffffff82111561049557610494610403565b5b61049e826103f3565b9050602081019050919050565b828183375f83830152505050565b5f6104cb6104c68461047b565b610461565b9050828152602081018484840111156104e7576104e66103ef565b5b6104f28482856104ab565b509392505050565b5f82601f83011261050e5761050d6103eb565b5b813561051e8482602086016104b9565b91505092915050565b5f6020828403121561053c5761053b6103e3565b5b5f82013567ffffffffffffffff811115610559576105586103e7565b5b610565848285016104fa565b91505092915050565b5f819050919050565b6105808161056e565b811461058a575f5ffd5b50565b5f8135905061059b81610577565b92915050565b5f602082840312156105b6576105b56103e3565b5b5f6105c38482850161058d565b91505092915050565b6105d58161056e565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61060d826105db565b61061781856105e5565b93506106278185602086016105f5565b610630816103f3565b840191505092915050565b5f819050919050565b61064d8161063b565b82525050565b5f608083015f8301516106685f8601826105cc565b50602083015184820360208601526106808282610603565b9150506040830151848203604086015261069a8282610603565b91505060608301516106af6060860182610644565b508091505092915050565b5f6020820190508181035f8301526106d28184610653565b905092915050565b5f5ffd5b5f5ffd5b5f815190506106f081610577565b92915050565b5f67ffffffffffffffff8211156107105761070f610403565b5b610719826103f3565b9050602081019050919050565b5f610738610733846106f6565b610461565b905082815260208101848484011115610754576107536103ef565b5b61075f8482856105f5565b509392505050565b5f82601f83011261077b5761077a6103eb565b5b815161078b848260208601610726565b91505092915050565b61079d8161063b565b81146107a7575f5ffd5b50565b5f815190506107b881610794565b92915050565b5f608082840312156107d3576107d26106da565b5b6107dd6080610461565b90505f6107ec848285016106e2565b5f83015250602082015167ffffffffffffffff81111561080f5761080e6106de565b5b61081b84828501610767565b602083015250604082015167ffffffffffffffff81111561083f5761083e6106de565b5b61084b84828501610767565b604083015250606061085f848285016107aa565b60608301525092915050565b5f602082840312156108805761087f6103e3565b5b5f82015167ffffffffffffffff81111561089d5761089c6103e7565b5b6108a9848285016107be565b91505092915050565b5f82825260208201905092915050565b7f4272616e6420494420697320696e76616c6964000000000000000000000000005f82015250565b5f6108f66013836108b2565b9150610901826108c2565b602082019050919050565b5f6020820190508181035f830152610923816108ea565b9050919050565b7f4272616e64206e616d6520697320696e76616c696400000000000000000000005f82015250565b5f61095e6015836108b2565b91506109698261092a565b602082019050919050565b5f6020820190508181035f83015261098b81610952565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806109d657607f821691505b6020821081036109e9576109e8610992565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610a4b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a10565b610a558683610a10565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610a90610a8b610a868461063b565b610a6d565b61063b565b9050919050565b5f819050919050565b610aa983610a76565b610abd610ab582610a97565b848454610a1c565b825550505050565b5f5f905090565b610ad4610ac5565b610adf818484610aa0565b505050565b5b81811015610b0257610af75f82610acc565b600181019050610ae5565b5050565b601f821115610b4757610b18816109ef565b610b2184610a01565b81016020851015610b30578190505b610b44610b3c85610a01565b830182610ae4565b50505b505050565b5f82821c905092915050565b5f610b675f1984600802610b4c565b1980831691505092915050565b5f610b7f8383610b58565b9150826002028217905092915050565b610b98826105db565b67ffffffffffffffff811115610bb157610bb0610403565b5b610bbb82546109bf565b610bc6828285610b06565b5f60209050601f831160018114610bf7575f8415610be5578287015190505b610bef8582610b74565b865550610c56565b601f198416610c05866109ef565b5f5b82811015610c2c57848901518255600182019150602085019450602081019050610c07565b86831015610c495784890151610c45601f891682610b58565b8355505b6001600288020188555050505b505050505050565b7f4272616e6420646f6573206e6f742065786973740000000000000000000000005f82015250565b5f610c926014836108b2565b9150610c9d82610c5e565b602082019050919050565b5f6020820190508181035f830152610cbf81610c86565b905091905056fea26469706673582212208d799c2bc198a06865c9b3ee0431a5d283686b823a3d5944dc5f795eab90a44864736f6c634300081d0033",
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
