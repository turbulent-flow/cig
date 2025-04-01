package helper

import "github.com/google/uuid"

func UuidToBytes32(uid uuid.UUID) [32]byte {
	idBytes := uid[:] // 格式为 16 字节
	var bytes32Id [32]byte
	copy(bytes32Id[:], idBytes)

	return bytes32Id
}
