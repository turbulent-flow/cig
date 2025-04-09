package action

import (
	"cig/internal/core"
	"cig/internal/pkg/action"
	"testing"

	"github.com/stretchr/testify/assert"
)

// case: it should add a brand.
func TestAddBrand(t *testing.T) {
	setUp(t)
	brand := new(core.Brand).Init()
	brand.Name = "fake-name"

	status, _ := action.AddBrand(brand)
	assert.Equal(t, 1, status)
}

func TestAddBrandWithInvalidName(t *testing.T) {
	setUp(t)
	brand := new(core.Brand).Init()
	_, err := action.AddBrand(brand)

	assert.IsType(t, &action.ActionError{}, err)
	assert.Contains(t, err.Error(), "Brand name is invalid")
}
