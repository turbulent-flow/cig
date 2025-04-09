package helper

import (
	"math/big"
	"time"

	"github.com/google/uuid"
)

func UuidToBytes32(uid uuid.UUID) [32]byte {
	idBytes := uid[:] // 格式为 16 字节
	var bytes32Id [32]byte
	copy(bytes32Id[:], idBytes)

	return bytes32Id
}

func Bytes32ToUuid(b32 [32]byte) uuid.UUID {
	var uid [16]byte
	copy(uid[:], b32[:16])
	return uuid.UUID(uid)
}

// 把 time.Time 转换为 *big.Int，对应 solidity 中的类型 uint256。
func TimeToBigInt(t time.Time) *big.Int {
	return big.NewInt(t.Unix())
}

func BigIntToTime(b *big.Int) time.Time {
	return time.Unix(b.Int64(), 0)
}
