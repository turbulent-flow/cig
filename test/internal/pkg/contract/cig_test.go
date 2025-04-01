package contract

import (
	"cig/internal/core"
	"cig/internal/pkg/contract"
	"cig/pkg/helper"
	"crypto/ecdsa"

	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/assert"
)

// 在测试可能会用到的结构体，包含一些共享的参数配置
type Config struct {
	PrivateKey       *ecdsa.PrivateKey
	WalletAddress    common.Address
	Backend          *simulated.Backend
	Client           simulated.Client
	Auth             *bind.TransactOpts
	ContractAdress   common.Address
	ContractInstance *contract.Contract
}

// 接收共享配置的变量
var SharedConfig *Config

// case: It should add a brand.
func TestAddBrand(t *testing.T) {
	// 在测试之前，对合约进行部署，并设置新的 nonce
	beforeTest(t)
	defer SharedConfig.Backend.Close() // 测试结束后，清理资源

	// 构造测试的参数
	brand := new(core.Brand).Init()
	brand.Name = "fake-name"
	idBytes32 := helper.UuidToBytes32(brand.Id)

	encodedBrand, err := brand.Encode()
	if err != nil {
		switch err.(type) {
		case *core.ReflectionTypeError:
			t.Fatalf("Failed to create reflectionType: %v", err)
		case *core.ArgumentsPackedError:
			t.Fatalf("Failed to pack cigBrand: %v", err)
		}
	}

	tx, err := SharedConfig.ContractInstance.AddBrand(SharedConfig.Auth, encodedBrand)
	if err != nil {
		t.Fatalf("Failed to add a brand: %v", err)
	}

	SharedConfig.Backend.Commit() // 把交易提交到区块

	fetchedBrand, err := SharedConfig.ContractInstance.GetBrand(nil, idBytes32)
	if err != nil {
		t.Fatalf("Failed to fetch the brand: %v", err)
	}

	assert.Equal(t, idBytes32, fetchedBrand.Id)
	assert.Equal(t, "fake-name", fetchedBrand.Name)

	_ = tx
}

// case: It should return an error with invalid name.
func TestAddBrandWithInvalidName(t *testing.T) {
	beforeTest(t)
	defer SharedConfig.Backend.Close()

	brand := new(core.Brand).Init()
	encodedBrand, _ := brand.Encode()
	_, err := SharedConfig.ContractInstance.AddBrand(SharedConfig.Auth, encodedBrand)
	if assert.Error(t, err) {
		assert.Equal(t, "execution reverted: Brand name is invalid", err.Error())
	}
}

// case: It should return an error with invalid id.
func TestAddBrandWithInvalidId(t *testing.T) {
	beforeTest(t)
	defer SharedConfig.Backend.Close()

	brand := &core.Brand{}
	encodedBrand, _ := brand.Encode()
	_, err := SharedConfig.ContractInstance.AddBrand(SharedConfig.Auth, encodedBrand)
	if assert.Error(t, err) {
		assert.Equal(t, "execution reverted: Brand ID is invalid", err.Error())
	}
}

func beforeTest(t *testing.T) {
	deploy(t)
	setNewNonce(t)
}

func setNewNonce(t *testing.T) {
	// 重新设置 nonce
	nonce, err := SharedConfig.Client.PendingNonceAt(context.Background(), SharedConfig.WalletAddress)
	if err != nil {
		t.Fatalf("Failed to fetch the nonce: %v", err)
	}

	SharedConfig.Auth.Nonce = big.NewInt(int64(nonce))
}

func deploy(t *testing.T) {
	SharedConfig = &Config{}

	// 生成测试账户
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate a private key: %v", err)
	}

	SharedConfig.PrivateKey = privateKey
	SharedConfig.WalletAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

	// 分配初始资金 1 ETH
	alloc := types.GenesisAlloc{
		SharedConfig.WalletAddress: {Balance: big.NewInt(1000000000000000000)},
	}

	// 创建一个模拟区块链的后端
	SharedConfig.Backend = simulated.NewBackend(alloc)
	SharedConfig.Client = SharedConfig.Backend.Client()

	// 创建一个 transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(SharedConfig.PrivateKey, big.NewInt(int64(1337)))
	if err != nil {
		t.Fatalf("Failed to create a transaction signer: %v", err)
	}

	SharedConfig.Auth = auth

	// 获取 nonce
	nonce, err := SharedConfig.Client.PendingNonceAt(context.Background(), SharedConfig.WalletAddress)
	if err != nil {
		t.Fatalf("Failed to fetch the nonce: %v", err)
	}

	SharedConfig.Auth.Value = big.NewInt(0) // 发送的 ETH 的数量
	SharedConfig.Auth.Nonce = big.NewInt(int64(nonce))

	// 这里未设置 gasLimit 是因为如果设置了它的值，比如：`config.Auth.GasLimit = uint64(3000000)`，等于禁止了 Gas 的自动估算；
	// 这样的话，当调用由 abigen 生成的“写方法”（带状态修改的合约函数）时，不会做本地的模拟执行，例如：当测试走到下面的代码行时，err 为 nil，
	// 表示交易发送成功，而不是交易执行成功；
	//
	// ```
	// tx, err := instance.AddBrand(config.Auth, encodedBrand)
	// if err != nil {
	// 	t.Fatalf("Failed to add a brand: %v", err)
	// }
	//
	// backend.Commit() // 把交易提交到区块
	// ```
	// 在不启用 Gas 自动估算的情况下，它会把交易广播出去，绕过本地的模拟执行，如果想要查看执行的结果，需要添加如下代码：
	//
	// ```
	// receipt, err := bind.WaitMined(context.Background(), client, tx)
	// if err != nil {
	// 	t.Fatalf("Failed to mine transaction: %v", err)
	// }

	// if receipt.Status == 1 {
	// 	t.Log("The transaction is executed successfully!")
	// } else {
	// 	t.Fatal("Failed to execute the transaction!")
	// }
	// ```
	//
	// 如果想要拿到具体的失败原因，得另外想办法，比如，`rever reason: Brand is invalid`；
	// 基于上述的观点，我们在测试中可以启用 Gas 的自动估算，也就是，把它设置为 0，或者干脆删掉，
	// 这样的话，就可以强制做本地的模拟执行，当运行到 `instanse.AddBrand(....)` 时，会在交易发送之前抛出`execution reverted`的错误。

	// 部署合约
	contractAddress, _, instance, err := contract.DeployContract(auth, SharedConfig.Client)
	if err != nil {
		t.Fatalf("Failed to deploy the contract: %v", err)
	}

	SharedConfig.ContractAdress = contractAddress
	SharedConfig.ContractInstance = instance

	// 强制“挖矿”，确认交易
	SharedConfig.Backend.Commit()

	// 验证部署结果
	decodedCode, err := SharedConfig.Client.CodeAt(context.Background(), SharedConfig.ContractAdress, nil)
	if err != nil {
		t.Fatalf("Something went wrong: %v", err)
	}

	if len(decodedCode) == 0 {
		t.Fatalf("The code of the contract is not stored.")
	}
}
