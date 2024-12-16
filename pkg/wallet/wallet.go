package wallet

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet struct {
	Client          *ethclient.Client
	KeyStoreManager *KeyStoreManager
}

func NewWallet(client *ethclient.Client, keyStoreManager *KeyStoreManager) *Wallet {
	return &Wallet{Client: client, KeyStoreManager: keyStoreManager}
}

func (w *Wallet) CheckBalance(ctx context.Context, address string) (string, error) {
	acc := common.HexToAddress(address)

	balance, nil := w.Client.BalanceAt(ctx, acc, nil)

	return balance.String(), nil
}

func (w *Wallet) TransferEth(ctx context.Context, from string, to string, amount *big.Float, password string) (string, error) {
	toAddress := common.HexToAddress(to)

	privateKey, error := w.KeyStoreManager.LoadWallet(from, password)

	if error != nil {
		return "", error
	}

	nonce, error := w.Client.PendingNonceAt(ctx, common.HexToAddress(from))

	if error != nil {
		return "", error
	}

	gasPrice, err := w.Client.SuggestGasPrice(ctx)

	if err != nil {
		return "", err
	}

	wei := new(big.Int)
	amountWei, _ := amount.Mul(amount, big.NewFloat(1e18)).Int(wei)

	log.Printf("Amount in wei: %v", amountWei)

	tx := types.NewTransaction(nonce, toAddress, amountWei, uint64(21000), gasPrice, nil)

	chainId, err := w.Client.NetworkID(ctx)

	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)

	if err != nil {
		return "", err
	}
	err = w.Client.SendTransaction(ctx, signedTx)

	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil

}
