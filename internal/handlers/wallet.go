package handlers

import (
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/yash-dev-ll/eth-wallet/pkg/wallet"
)

type WalletHandler struct {
	KeyStore *wallet.KeyStoreManager
	Client   *ethclient.Client
}

func (h *WalletHandler) CheckBalanceHandler(c *gin.Context) {
	address := c.Param("address")

	balance, err := wallet.CheckBalance(c, h.Client, address)

	if err != nil {
		log.Printf("Failed to check balance: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})

}

func (h *WalletHandler) CreateWalletKeyStoreHandler(c *gin.Context) {
	var req struct {
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	account, err := h.KeyStore.CreateWallet(req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": account.Address.Hex()})

}

func (h *WalletHandler) LoadWalletKeyStoreHandler(c *gin.Context) {
	var req struct {
		Address  string `json:"address"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	privateKey, err := h.KeyStore.LoadWallet(req.Address, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Wallet loaded successfully",
		"private_key": crypto.PubkeyToAddress(privateKey.PublicKey).Hex(),
	})
}
