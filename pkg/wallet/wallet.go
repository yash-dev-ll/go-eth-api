package wallet

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckBalance(ctx context.Context, client *ethclient.Client, address string) (string, error) {
	acc := common.HexToAddress(address)

	balance, nil := client.BalanceAt(ctx, acc, nil)

	return balance.String(), nil

}
