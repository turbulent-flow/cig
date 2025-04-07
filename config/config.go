package config

import (
	"context"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
)

// 默认的全局配置
type Config struct {
	EthereumEndpoint string
	PrivateKey       string
	WalletAddress    string
	GasLimit         int
	ContractAddress  string
}

// 与以太坊操作相关的配置
type EthConfig struct {
	Client *ethclient.Client
	Atuh   *bind.TransactOpts
}

var (
	DefaultConfig    *Config    = &Config{}
	DefaultEthConfig *EthConfig = &EthConfig{}
)

func init() {
	// 初始化默认配置
	DefaultConfig.Init()

	// 初始化以太坊的操作
	cf := DefaultConfig
	DefaultEthConfig.initEth(cf)
}

func (ec *EthConfig) initEth(cf *Config) {
	// 连接到以太坊
	client, err := ethclient.Dial(DefaultConfig.EthereumEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect the ethereum endpoint: %v", err)
	}

	ec.Client = client

	// 用私钥创建交易签名
	privateKey, err := crypto.HexToECDSA(cf.PrivateKey)
	if err != nil {
		log.Fatalf("The private key is invalid")
	}

	walletAddress := common.HexToAddress(cf.WalletAddress)

	// 创建一个 transaction signer
	chainID, err := ec.Client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch chainID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create a transaction signer: %v", err)
	}

	// 获取 nonce
	nonce, err := ec.Client.PendingNonceAt(context.Background(), walletAddress)
	if err != nil {
		log.Fatalf("Failed to fetch the nonce: %v", err)
	}

	// 获取 SuggestGasPrice
	gasPrice, err := ec.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch SuggestGasPrice: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // 发送的 ETH 的数量
	auth.GasLimit = uint64(cf.GasLimit)
	auth.GasPrice = gasPrice

	ec.Atuh = auth
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
