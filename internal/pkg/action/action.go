package action

import (
	"cig/internal/core"
	"cig/internal/pkg/contract"
	"cig/pkg/helper"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ActionError struct {
	Message string
}

func (e *ActionError) Error() string {
	return e.Message
}

func newInstance(client *ethclient.Client) (instance *contract.Contract) {
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create an instance of the contract: %v", err)
	}

	return
}

func CIGBrandToBrand(cbrand *contract.CIGBrand) *core.Brand {
	return &core.Brand{
		Id:          helper.Bytes32ToUuid(cbrand.Id),
		Name:        cbrand.Name,
		Description: cbrand.Description,
		CreatedAt:   helper.BigIntToTime(cbrand.CreatedAt),
	}
}
