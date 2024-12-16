package services

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yash-dev-ll/eth-wallet/pkg/wallet"
)

type WalletService struct {
	wallet *wallet.Wallet
}

func NewWalletService(client *ethclient.Client, keyStoreManager *wallet.KeyStoreManager) *WalletService {
	w := wallet.NewWallet(client, keyStoreManager)
	return &WalletService{wallet: w}
}

func (ws *WalletService) TransferEth(ctx context.Context, from string, to string, amount *big.Float, password string) (string, error) {
	log.Printf("Transfering %v eth from %v to %v", amount, from, to)
	txHash, err := ws.wallet.TransferEth(ctx, from, to, amount, password)
	if err != nil {
		log.Printf("Failed to transfer eth: %v", err)
		return "", err
	}

	return txHash, nil
}

func (ws *WalletService) CheckBalance(ctx context.Context, address string) (string, error) {
	balance, err := ws.wallet.CheckBalance(ctx, address)
	if err != nil {
		log.Printf("Failed to check balance: %v", err)
		return "", err
	}

	return balance, nil
}

func (ws *WalletService) CreateWallet(password string) (string, error) {
	account, err := ws.wallet.KeyStoreManager.CreateWallet(password)
	if err != nil {
		log.Printf("Failed to create wallet: %v", err)
		return "", err
	}

	return account.Address.Hex(), nil
}

func (ws *WalletService) LoadWallet(address string, password string) (string, error) {
	privateKey, err := ws.wallet.KeyStoreManager.LoadWallet(address, password)
	if err != nil {
		log.Printf("Failed to load wallet: %v", err)
		return "", err
	}

	privateKeyHex := fmt.Sprintf("%x", privateKey.D.Bytes())
	log.Printf("Private key: %v", privateKeyHex)

	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex(), nil
}
