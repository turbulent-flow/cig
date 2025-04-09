package action

import (
	"cig/config"
	"cig/internal/core"
	"cig/pkg/helper"

	"github.com/google/uuid"
)

func GetBrand(id uuid.UUID) (*core.Brand, error) {
	config.BeforeAction()
	ec := config.DefaultEthConfig
	defer ec.Client.Close()

	bytes32Id := helper.UuidToBytes32(id)
	cigBrand, err := newInstance(ec.Client).GetBrand(nil, bytes32Id)
	if err != nil {
		err = &ActionError{
			Message: err.Error(),
		}

		return nil, err
	}

	return CIGBrandToBrand(&cigBrand), err
}
