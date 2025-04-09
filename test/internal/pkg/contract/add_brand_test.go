package contract

import (
	"cig/internal/core"
	"cig/pkg/helper"

	"testing"

	"github.com/stretchr/testify/assert"
)

// case: It should add a brand.
func TestAddBrand(t *testing.T) {
	ec := DefaultEthConfig
	ec.BeforeTest(t)
	defer ec.Backend.Close() // 测试结束后，清理资源

	// 构造测试的参数
	brand := new(core.Brand).Init()
	brand.Name = "fake-name"
	idBytes32 := helper.UuidToBytes32(brand.Id)

	encodedBrand, err := brand.Encode()
	if err != nil {
		switch err.(type) {
		case *core.ReflectionTypeError:
			t.Fatalf("Failed to create reflectionType: %v", err)
		case *core.ArgumentsPackedError:
			t.Fatalf("Failed to pack cigBrand: %v", err)
		}
	}

	contract := ec.ContractInstance
	tx, err := contract.AddBrand(ec.Auth, encodedBrand)
	if err != nil {
		t.Fatalf("Failed to add a brand: %v", err)
	}

	ec.Backend.Commit() // 把交易提交到区块

	fetchedBrand, err := contract.GetBrand(nil, idBytes32)
	if err != nil {
		t.Fatalf("Failed to fetch the brand: %v", err)
	}

	assert.Equal(t, idBytes32, fetchedBrand.Id)
	assert.Equal(t, "fake-name", fetchedBrand.Name)

	_ = tx
}

// case: It should return an error with invalid name.
func TestAddBrandWithInvalidName(t *testing.T) {
	ec := DefaultEthConfig
	ec.BeforeTest(t)
	defer ec.Backend.Close()

	brand := new(core.Brand).Init()
	encodedBrand, _ := brand.Encode()

	contract := ec.ContractInstance
	_, err := contract.AddBrand(ec.Auth, encodedBrand)
	if assert.Error(t, err) {
		assert.Equal(t, "execution reverted: Brand name is invalid", err.Error())
	}
}

// case: It should return an error with invalid id.
func TestAddBrandWithInvalidId(t *testing.T) {
	ec := DefaultEthConfig
	ec.BeforeTest(t)
	defer ec.Backend.Close()

	brand := &core.Brand{}
	encodedBrand, _ := brand.Encode()

	contract := ec.ContractInstance
	_, err := contract.AddBrand(ec.Auth, encodedBrand)
	if assert.Error(t, err) {
		assert.Equal(t, "execution reverted: Brand ID is invalid", err.Error())
	}
}
