package deploy

import (
	"cig/internal/pkg/deploy"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// case: it should deploy a contract.
func TestRun(t *testing.T) {
	beforeTest(t)
	receipt := deploy.Run()

	assert.Equal(t, 1, int(receipt.Status))
}

func beforeTest(t *testing.T) {
	err := godotenv.Load("../../../../.env.test")
	if err != nil {
		t.Fatalf("Failed to load .env.test: %v", err)
	}
}
