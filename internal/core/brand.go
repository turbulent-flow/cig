package core

import (
	"cig/internal/pkg/contract"
	"cig/pkg/helper"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/google/uuid"
)

type Brand struct {
	Id          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
}

func (brand *Brand) Init() *Brand {
	brand.Id = uuid.New()
	brand.CreatedAt = time.Now().Truncate(time.Second)

	return brand
}

// 使用 abi 工具把 golang 的 struct 转化为合约可以接受的格式，
// 把其作为参数传给合约后，在合约中使用 abi 进行解码，可以得到符合特定的
// 结构体的类型的数据，比如，把 golang 的结构体 Brand 转化为合约中的结构体
// Brand。
func (brand *Brand) Encode() (encoded []byte, err error) {
	brandReflectionType, err := brand.NewTuple()
	if err != nil {
		return
	}

	encoded, err = brand.PackArguments(brandReflectionType)
	if err != nil {
		return
	}

	return
}

type ArgumentsPackedError struct {
	Message string
}

func (e *ArgumentsPackedError) Error() string {
	return e.Message
}

func (brand *Brand) PackArguments(reflectionType *abi.Type) ([]byte, error) {
	brandArgs := abi.Arguments{
		{Type: *reflectionType},
	}

	cigBrand := &contract.CIGBrand{
		Id:          helper.UuidToBytes32(brand.Id),
		Name:        brand.Name,
		Description: brand.Description,
		CreatedAt:   helper.TimeToBigInt(brand.CreatedAt),
	}

	// 再进行打包，构成合约可以接受的格式，在合约中使用 abi 解码后，
	// 符合合约中特定的结构体的类型。
	encoded, err := brandArgs.Pack(cigBrand)

	if err != nil {
		err = &ArgumentsPackedError{
			Message: err.Error(),
		}

		return encoded, err
	}

	return encoded, err
}

type ReflectionTypeError struct {
	Message string
}

func (e *ReflectionTypeError) Error() string {
	return e.Message
}

// 创建一个 tuple 类型，在 abi 中，tuple 表示 struct。
// reflectionType 是参数的类型的映射类型，比如，参数是
// 合约中`addBrand(bytes memory encodedBrand)`中的
// encodedBrand。
func (brand *Brand) NewTuple() (*abi.Type, error) {
	reflectionType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "id", Type: "bytes32"},
		{Name: "name", Type: "string"},
		{Name: "description", Type: "string"},
		{Name: "createdAt", Type: "uint256"},
	})

	if err != nil {
		err = &ReflectionTypeError{
			Message: err.Error(),
		}
	}

	return &reflectionType, err
}
