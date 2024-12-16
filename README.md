# Ethereum Wallet Service

This project is an Ethereum wallet service built using Go. It provides APIs to create wallets, check balances, load wallets, and transfer Ether. The service uses the Gin web framework for handling HTTP requests and the go-ethereum package for interacting with the Ethereum blockchain.

## Project Structure

```
.env
.gitignore
cmd/
    api/
        main.go
go.mod
go.sum
internal/
    handlers/
        wallet.go
    services/
        wallet_service.go
pkg/
    wallet/
        keystore.go
        wallet.go
```

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yash-dev-ll/eth-wallet.git
   cd eth-wallet
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Create a .envfile in the root directory with the following content:

`ETH_RPC_URL=https://eth-sepolia.g.alchemy.com/v2/x9HZoP9QSiI2amPmlBHsV_lZyhJndl6P`

## Running the Service

1. Start the service:
   ```sh
   go run
   ```

main.go

    ```

2. The service will be running on `http://localhost:9001`.

## API Endpoints

### Check Balance

- **Endpoint:** `GET /wallet/:address/balance`
- **Description:** Checks the balance of the specified Ethereum address.
- **Response:**
  ```json
  {
    "balance": "1000000000000000000"
  }
  ```

### Create Wallet

- **Endpoint:** `POST /wallet/new/keystore`
- **Description:** Creates a new wallet with the specified password.
- **Request Body:**
  ```json
  {
    "password": "your_password"
  }
  ```
- **Response:**
  ```json
  {
    "address": "0xYourNewWalletAddress"
  }
  ```

### Load Wallet

- **Endpoint:** `GET /wallet/keystore`
- **Description:** Loads a wallet from the keystore using the specified address and password.
- **Request Body:**
  ```json
  {
    "address": "0xYourWalletAddress",
    "password": "your_password"
  }
  ```
- **Response:**
  ```json
  {
    "message": "Wallet loaded successfully",
    "private_key": "your_private_key"
  }
  ```

### Transfer Ether

- **Endpoint:** `POST /wallet/transferEth`
- **Description:** Transfers Ether from one address to another.
- **Request Body:**
  ```json
  {
    "from": "0xYourWalletAddress",
    "to": "0xRecipientAddress",
    "amount": 0.1,
    "password": "your_password"
  }
  ```
- **Response:**
  ```json
  {
    "tx_hash": "0xTransactionHash"
  }
  ```

## Project Files

- **[cmd/api/main.go](cmd/api/main.go):** Entry point of the application. Sets up the Gin router and defines the API endpoints.
- **[internal/handlers/wallet.go](internal/handlers/wallet.go):** Contains the HTTP handlers for the wallet-related endpoints.
- **[internal/services/wallet_service.go](internal/services/wallet_service.go):** Implements the business logic for wallet operations.
- **[pkg/wallet/keystore.go](pkg/wallet/keystore.go):** Manages the keystore operations, including creating and loading wallets.
- **[pkg/wallet/wallet.go](pkg/wallet/wallet.go):** Contains the wallet operations, including checking balance and transferring Ether.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.## Project Files

- **[cmd/api/main.go](cmd/api/main.go):** Entry point of the application. Sets up the Gin router and defines the API endpoints.
- **[internal/handlers/wallet.go](internal/handlers/wallet.go):** Contains the HTTP handlers for the wallet-related endpoints.
- **[internal/services/wallet_service.go](internal/services/wallet_service.go):** Implements the business logic for wallet operations.
- **[pkg/wallet/keystore.go](pkg/wallet/keystore.go):** Manages the keystore operations, including creating and loading wallets.
- **[pkg/wallet/wallet.go](pkg/wallet/wallet.go):** Contains the wallet operations, including checking balance and transferring Ether.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
