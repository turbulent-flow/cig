package config

import (
	"cig/internal/pkg/contract"
	"context"
	"math/big"

	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
)

type EthConfig struct {
	Client           simulated.Client
	Auth             *bind.TransactOpts
	WalletAddress    common.Address
	Backend          *simulated.Backend
	ContractAddress  common.Address
	ContractInstance *contract.Contract
}

var DefaultEthConfig *EthConfig = &EthConfig{}

// 在测试之前，对合约进行部署，并设置新的 nonce
func (ec *EthConfig) BeforeTest(t *testing.T) {
	ec.deploy(t)
	ec.setNewNonce(t)
}

func (ec *EthConfig) setNewNonce(t *testing.T) {
	nonce, err := ec.Client.PendingNonceAt(context.Background(), ec.WalletAddress)
	if err != nil {
		t.Fatalf("Failed to fetch the nonce: %v", err)
	}

	ec.Auth.Nonce = big.NewInt(int64(nonce))
}

func (ec *EthConfig) deploy(t *testing.T) {
	// 生成测试账户
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate a private key: %v", err)
	}

	ec.WalletAddress = crypto.PubkeyToAddress(privateKey.PublicKey)

	// 分配初始资金 1 ETH
	alloc := types.GenesisAlloc{
		ec.WalletAddress: {Balance: big.NewInt(1000000000000000000)},
	}

	// 创建一个模拟区块链的后端
	ec.Backend = simulated.NewBackend(alloc)
	ec.Client = ec.Backend.Client()

	// 创建一个 transaction signer
	ec.Auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(1337)))
	if err != nil {
		t.Fatalf("Failed to create a transaction signer: %v", err)
	}

	// 获取 nonce
	nonce, err := ec.Client.PendingNonceAt(context.Background(), ec.WalletAddress)
	if err != nil {
		t.Fatalf("Failed to fetch the nonce: %v", err)
	}

	ec.Auth.Value = big.NewInt(0) // 发送的 ETH 的数量
	ec.Auth.Nonce = big.NewInt(int64(nonce))

	// 这里未设置 gasLimit 是因为如果设置了它的值，比如：`ec.Auth.GasLimit = uint64(3000000)`，等于禁止了 Gas 的自动估算；
	// 这样的话，当调用由 abigen 生成的“写方法”（带状态修改的合约函数）时，不会做本地的模拟执行，例如：当测试走到下面的代码行时，err 为 nil，
	// 表示交易发送成功，而不是交易执行成功；
	//
	// ```
	// tx, err := instance.AddBrand(ec.Auth, encodedBrand)
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
	contractAddress, tx, instance, err := contract.DeployContract(ec.Auth, ec.Client)
	if err != nil {
		t.Fatalf("Failed to deploy the contract: %v", err)
	}

	ec.ContractAddress = contractAddress
	ec.ContractInstance = instance

	// 强制“挖矿”，确认交易
	ec.Backend.Commit()

	// 验证部署结果
	decodedCode, err := ec.Client.CodeAt(context.Background(), ec.ContractAddress, nil)
	if err != nil {
		t.Fatalf("Something went wrong: %v", err)
	}

	if len(decodedCode) == 0 {
		t.Fatalf("The code of the contract is not stored.")
	}

	_ = tx
}
