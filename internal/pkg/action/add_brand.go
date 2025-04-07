package action

import (
	"cig/config"
	"cig/internal/core"
	"cig/internal/pkg/contract"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func AddBrand(brand *core.Brand) {
	cf := config.DefaultConfig
	ec := config.DefaultEthConfig
	defer ec.Client.Close()

	contractAddress := common.HexToAddress(cf.ContractAddress)
	instance, err := contract.NewContract(contractAddress, ec.Client)
	if err != nil {
		log.Fatalf("Failed to create an instance of the contract: %v", err)
	}

	encoded, err := brand.Encode()
	if err != nil {
		switch err.(type) {
		case *core.ReflectionTypeError:
			log.Fatalf("Failed to create reflectionType: %v", err)
		case *core.ArgumentsPackedError:
			log.Fatalf("Failed to pack cigBrand: %v", err)
		}
	}

	tx, err := instance.AddBrand(ec.Atuh, encoded)
	if err != nil {
		log.Fatalf("Failed to add a brand: %v", err)
	}

	receipt, err := bind.WaitMined(context.Background(), ec.Client, tx)
	if err != nil {
		log.Fatalf("Failed to execute the transaction: %v", err)
	}

	if receipt.Status == 1 {
		log.Print("The transaction is executed successfully!")
	} else {
		log.Print("Failed to execute the transaction!")
	}
}
