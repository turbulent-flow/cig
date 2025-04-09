package action

import (
	"cig/config"
	"cig/internal/core"
	"cig/internal/pkg/contract"
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ActionError struct {
	Message string
}

func (e *ActionError) Error() string {
	return e.Message
}

func AddBrand(brand *core.Brand) (status int, err error) {
	config.BeforeAction()
	ec := config.DefaultEthConfig
	defer ec.Client.Close()

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
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
		err = &ActionError{
			Message: err.Error(),
		}

		return 0, err
	}

	receipt, err := bind.WaitMined(context.Background(), ec.Client, tx)
	if err != nil {
		log.Fatalf("Failed to execute the transaction: %v", err)
	}

	status = int(receipt.Status)

	return status, nil
}
