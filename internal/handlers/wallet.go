package handlers

import (
	"log"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yash-dev-ll/eth-wallet/internal/services"
)

type WalletHandler struct {
	WalletService *services.WalletService
}

func (h *WalletHandler) CheckBalanceHandler(c *gin.Context) {
	address := c.Param("address")

	balance, err := h.WalletService.CheckBalance(c, address)

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

	account, err := h.WalletService.CreateWallet(req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"address": account})

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

	privateKey, err := h.WalletService.LoadWallet(req.Address, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Wallet loaded successfully",
		"private_key": privateKey,
	})
}

func (h *WalletHandler) TransferEthHandler(c *gin.Context) {
	var req struct {
		From     string  `json:"from"`
		To       string  `json:"to"`
		Amount   float64 `json:"amount"`
		Password string  `json:"password"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	txHash, err := h.WalletService.TransferEth(c, req.From, req.To, big.NewFloat(req.Amount), req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to transfer eth"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tx_hash": txHash})
}
