package config

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	EthereumEndpoint string
	PrivateKey       string
	WalletAddress    string
	GasLimit         int
}

var DefaultConfig *Config = &Config{}

func init() {
	DefaultConfig.Init()
}

func (cf *Config) Init() {
	cf.EthereumEndpoint = os.Getenv("ETHEREUM_ENDPOINT")
	cf.PrivateKey = os.Getenv("PRIVATE_KEY")
	cf.WalletAddress = os.Getenv("WALLET_ADDRESS")

	gasLimit, err := strconv.Atoi(os.Getenv("GAS_LIMIT"))
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}

	cf.GasLimit = gasLimit
}
