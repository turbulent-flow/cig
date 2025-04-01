package deploy

import (
	"cig/config"
	"cig/internal/pkg/contract"
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run() {
	// 获取默认的全局配置
	cf := config.DefaultConfig

	// 连接到以太坊
	client, err := ethclient.Dial(cf.EthereumEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect the ethereum endpoint: %v", err)
	}

	defer client.Close()

	// 用私钥创建交易签名
	privateKey, err := crypto.HexToECDSA(cf.PrivateKey)
	if err != nil {
		log.Fatalf("The private key is invalid")
	}

	walletAddress := common.HexToAddress(cf.WalletAddress)

	// 创建一个 transaction signer
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch chainID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create a transaction signer: %v", err)
	}

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), walletAddress)
	if err != nil {
		log.Fatalf("Failed to fetch the nonce: %v", err)
	}

	// 获取 SuggestGasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch SuggestGasPrice: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // 发送的 ETH 的数量
	auth.GasLimit = uint64(cf.GasLimit)
	auth.GasPrice = gasPrice

	// 部署合约
	contractAddress, tx, instance, err := contract.DeployContract(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy the contract: %v", err)
	}

	log.Printf("合约地址是：%v\n", contractAddress)
	log.Printf("交易哈希是：%v\n", tx.Hash().Hex())

	// 验证部署的结果
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("Failed to mine the transaction: %v", err)
	}

	if receipt.Status == 1 {
		log.Println("部署成功！")
	} else {
		log.Println("部署失败！")
	}

	_ = instance
}
