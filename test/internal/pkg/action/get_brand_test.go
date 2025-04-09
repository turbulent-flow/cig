package action

import (
	"cig/internal/core"
	"cig/internal/pkg/action"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var brand *core.Brand

// case: it should fetch the brand.
func TestGetBrand(t *testing.T) {
	beforeTest(t)

	fetchedBrand, _ := action.GetBrand(brand.Id)
	assert.Equal(t, brand, fetchedBrand)
}

// case: it should return an error with invalid id.
func TestGetBrandWithInvalidId(t *testing.T) {
	beforeTest(t)

	InvalidId := uuid.New()
	_, err := action.GetBrand(InvalidId)
	assert.IsType(t, &action.ActionError{}, err)
	assert.Contains(t, err.Error(), "Brand does not exist")
}

func beforeTest(t *testing.T) {
	setUp(t)

	brand = new(core.Brand).Init()
	brand.Name = "fake-name"
	status, _ := action.AddBrand(brand)
	assert.Equal(t, 1, status)
}
