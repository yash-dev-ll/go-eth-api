package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/yash-dev-ll/eth-wallet/internal/handlers"
	"github.com/yash-dev-ll/eth-wallet/pkg/wallet"
)

func main() {
	keyStoreDir := "./keystore"
	keyStoreManager, err := wallet.NewKeystoreManager(keyStoreDir)

	if err != nil {
		log.Fatal("Failed to create keystore manager: %w", err)
	}

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/x9HZoP9QSiI2amPmlBHsV_lZyhJndl6P")

	if err != nil {
		log.Fatal("Failed to connect to the Ethereum network: %w", err)

	}

	walletHandler := handlers.WalletHandler{Client: client, KeyStore: keyStoreManager}

	r := gin.Default()

	r.GET("/wallet/:address/balance", walletHandler.CheckBalanceHandler)
	r.POST("/wallet/new/keystore", walletHandler.CreateWalletKeyStoreHandler)
	r.GET("/wallet/keystore", walletHandler.LoadWalletKeyStoreHandler)

	err = r.Run(":9001")

	if err != nil {
		log.Fatal("Failed to start the server: %w", err)
	}
}
