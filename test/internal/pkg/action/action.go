package action

import (
	"cig/config"
	"cig/internal/pkg/contract"
	"context"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../../../.env.test")
	if err != nil {
		log.Fatalf("Failed to load .env.test: %v", err)
	}
}

// 在每个测试之前，重新部署合约，并把合约地址写进环境变量。
func setUp(t *testing.T) {
	config.BeforeAction()
	ec := config.DefaultEthConfig
	defer ec.Client.Close()

	contractAddress, tx, _, err := contract.DeployContract(ec.Atuh, ec.Client)
	if err != nil {
		t.Fatalf("Failed to deploy the contract: %v", err)
	}

	receipt, err := bind.WaitMined(context.Background(), ec.Client, tx)
	if err != nil {
		t.Fatalf("Failed to mine the transaction: %v", err)
	}

	if receipt.Status != 1 {
		t.Fatal("Failed to deploy the contract!")
	}

	os.Setenv("CONTRACT_ADDRESS", contractAddress.Hex())
}
