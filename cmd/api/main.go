package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yash-dev-ll/eth-wallet/internal/handlers"
	"github.com/yash-dev-ll/eth-wallet/internal/services"
	"github.com/yash-dev-ll/eth-wallet/pkg/wallet"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	keyStoreDir := "./keystore"
	keyStoreManager, err := wallet.NewKeystoreManager(keyStoreDir)

	if err != nil {
		log.Fatal("Failed to create keystore manager: %w", err)
	}

	client, err := ethclient.Dial(os.Getenv("ETH_RPC_URL"))

	if err != nil {
		log.Fatal("Failed to connect to the Ethereum network: %w", err)

	}

	walletService := services.NewWalletService(client, keyStoreManager)

	walletHandler := handlers.WalletHandler{WalletService: walletService}

	r := gin.Default()

	r.GET("/wallet/:address/balance", walletHandler.CheckBalanceHandler)
	r.POST("/wallet/new/keystore", walletHandler.CreateWalletKeyStoreHandler)
	r.GET("/wallet/keystore", walletHandler.LoadWalletKeyStoreHandler)
	r.POST("/wallet/transferEth", walletHandler.TransferEthHandler)

	err = r.Run(":9001")

	if err != nil {
		log.Fatal("Failed to start the server: %w", err)
	}
}
