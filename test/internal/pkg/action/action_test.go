package action

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../../../.env.test")
	if err != nil {
		log.Fatalf("Failed to load .env.test: %v", err)
	}
}
