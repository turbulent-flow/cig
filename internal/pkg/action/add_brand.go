package action

import (
	"cig/config"
	"cig/internal/core"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func AddBrand(brand *core.Brand) (status int, err error) {
	config.BeforeAction()
	ec := config.DefaultEthConfig
	defer ec.Client.Close()

	encoded, err := brand.Encode()
	if err != nil {
		switch err.(type) {
		case *core.ReflectionTypeError:
			log.Fatalf("Failed to create reflectionType: %v", err)
		case *core.ArgumentsPackedError:
			log.Fatalf("Failed to pack cigBrand: %v", err)
		}
	}

	tx, err := newInstance(ec.Client).AddBrand(ec.Atuh, encoded)
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
