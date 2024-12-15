package wallet

import (
	"crypto/ecdsa"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type KeyStoreManager struct {
	Keystore   *keystore.KeyStore
	StorageDir string
}

func NewKeystoreManager(keyStoreDir string) (*KeyStoreManager, error) {
	if _, err := os.Stat(keyStoreDir); os.IsNotExist(err) {
		err := os.MkdirAll(keyStoreDir, 0700)
		if err != nil {
			return nil, err
		}
	}

	ks := keystore.NewKeyStore(keyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

	return &KeyStoreManager{Keystore: ks, StorageDir: keyStoreDir}, nil
}

func (km *KeyStoreManager) CreateWallet(password string) (accounts.Account, error) {
	acc, err := km.Keystore.NewAccount(password)

	if err != nil {
		return accounts.Account{}, err
	}

	return acc, nil
}

func (km *KeyStoreManager) LoadWallet(add string, password string) (*ecdsa.PrivateKey, error) {

	address := convertHexAddress(add)
	var walletFile string

	err := filepath.WalkDir(km.StorageDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Failed to read directory: %v", err)
			return err
		}
		if strings.Contains(path, address) {
			log.Printf("Address: %v", address)
			walletFile = path

			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if walletFile == "" {
		log.Printf("Wallet file not found")
		return nil, os.ErrNotExist
	}

	keyJson, err := os.ReadFile(walletFile)

	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil, err
	}

	key, err := keystore.DecryptKey(keyJson, password)

	if err != nil {
		log.Printf("Failed to decrypt key: %v", err)
		return nil, err
	}

	return key.PrivateKey, nil
}

func convertHexAddress(address string) string {
	address = strings.TrimPrefix(address, "0x")
	return strings.ToLower(address)
}
