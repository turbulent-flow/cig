package deploy

import (
	"cig/config"
	"cig/internal/pkg/contract"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func Run() {
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
	receipt, err := bind.WaitMined(context.Background(), ec.Client, tx)
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
