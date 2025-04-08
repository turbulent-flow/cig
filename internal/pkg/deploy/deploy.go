package deploy

import (
	"cig/config"
	"cig/internal/pkg/contract"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func Run() (receipt *types.Receipt) {
	config.BeforeAction()
	ec := config.DefaultEthConfig
	defer ec.Client.Close()

	// 部署合约
	contractAddress, tx, instance, err := contract.DeployContract(ec.Atuh, ec.Client)
	if err != nil {
		log.Fatalf("Failed to deploy the contract: %v", err)
	}

	log.Printf("合约地址是：%v\n", contractAddress)
	log.Printf("交易哈希是：%v\n", tx.Hash().Hex())

	// 验证部署的结果
	receipt, err = bind.WaitMined(context.Background(), ec.Client, tx)
	if err != nil {
		log.Fatalf("Failed to mine the transaction: %v", err)
	}

	_ = instance

	return
}
