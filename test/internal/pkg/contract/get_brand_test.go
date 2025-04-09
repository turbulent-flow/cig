package contract

import (
	"cig/pkg/helper"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// case: It should return an error with empty id.
func TestGetBrandWithEmptyId(t *testing.T) {
	ec := DefaultEthConfig
	ec.BeforeTest(t)
	defer ec.Backend.Close()

	var emptyBytes32Id [32]byte
	_, err := ec.ContractInstance.GetBrand(nil, emptyBytes32Id)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Brand ID is invalid")
	}
}

// case: It should return an error with invalid id.
func TestGetBrandWithInvalidId(t *testing.T) {
	ec := DefaultEthConfig
	ec.BeforeTest(t)
	defer ec.Backend.Close()

	id := uuid.New()
	invalidBytes32Id := helper.UuidToBytes32(id)
	_, err := ec.ContractInstance.GetBrand(nil, invalidBytes32Id)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Brand does not exist")
	}
}
