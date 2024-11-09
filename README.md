# GoChain CLI

GoChain CLI is a basic blockchain implementation in Go. This application allows interactions with the Ethereum network, including wallet generation, transaction creation, and checking balances and network data.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/cli-gochain.git
   cd cli-gochain
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build the application:
   ```bash
   go build -o gochain
   ```
## Usage
### Available commands:
- `generate-wallet`: Generates a new Ethereum wallet and saves it in memory.
   ```bash
   ./gochain generate-wallet
   ```
   Output:
    ```
    PublicKey: 0x123...abc
    PrivateKey: 0xdef...456
    ```
- `get-wallets-inmemory`: Displays all Ethereum wallets currently stored in memory.
   ```bash
   ./gochain get-wallets-inmemory
    ```
    Output:
    ```
    Ethereum wallets in memory: [0x123...abc, 0xdef...456]
    ```
- `set-network-node <network> <nodeUrl>`: Sets the network node for the Ethereum client.
   ```bash
   ./gochain set-network-node mainnet https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID
   ```
   Output:
    ```
    Network mainnet successfully added
    ```
- `get-blocknumber <network>`: Gets the current block number of a specified network.
   ```bash
   ./gochain get-blocknumber mainnet
   ```
   Output:
    ```
    Current blockNumber: 13245678
    ```
- `create-transaction <network> <from> <to> <amount>`: Creates an Ethereum transaction.
    ```bash
   ./gochain create-transaction mainnet 0x123...abc 0x456...def 1000000000000000000
   ```
   Output:
    ```
    Transaction created successfully with hash: 0xabc123...789
    ```
    Note: The `from` address must be one of the addresses generated by the `generate-wallet` command.
- `get-balance <network> <address>`: Retrieves the balance of an Ethereum address.
    ```bash
   ./gochain get-balance mainnet 0x123...abc
    ```
    Output:
    ```
    Balance: 1.234567890123456789 ETH
    ```
- `get-transaction-receipt <network> <txHash>`: Gets the receipt of a specific transaction.
    ```bash
   ./gochain get-transaction-receipt mainnet 0xabc123...789
    ```
    Output:
    ```
    Transaction receipt: {Status: 1, BlockNumber: 13245678, GasUsed: 21000, ...}
    ```

## Code structure
- `main.go`: The main file that defines the CLI using `cobra.Command`.
- `database/memory.go`: Responsible for in-memory storage of data such as wallets and network settings.
- `evm`:  Module that interacts directly with the Ethereum Virtual Machine (EVM), including wallet creation, transactions, and network data retrieval.

## Dependencies
- `Cobra`: A library for creating beautiful command-line interfaces in Go.
- `Go-Ethereum`: A Go implementation of the Ethereum protocol.

## License
This project is licensed under the MIT License.