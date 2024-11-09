package evm

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// newClient creates a new Ethereum client and returns it.
func newClient(nodeUrl string) *ethclient.Client {
	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatalf("Error to connect with node: %v", err)
	}
	defer client.Close()

	_, err = client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Error to get Network ID: %v", err)
	}
	return client
}

// GenerateEvmWallet creates a new Ethereum wallet and returns the address and store private key in memory.
func GenerateWallet() (string, error) {
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %v", err)
	}

	// Extract the public key from the private key
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	// Generate the address from the public key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey).Hex()

	WalletsInMemory[address] = privateKeyHex
	return address, nil
}

// CreateEvmTransaction creates a new Ethereum transaction and sends it to the network.
func CreateTransaction(network, privateKeyHex, to string, amount int64) (string, error) {
	client := newClient(NetworkInMemory[network])

	// Privatekey to sign the transaction
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Get the address of the account to send the transaction from
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	// Get the nonce of the account to prevent replay attacks
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}

	// Configure the transaction parameters
	toAddress := common.HexToAddress(to) // address to send the transaction to
	value := big.NewInt(amount)          // value to send in wei (1 ether = 10^18 wei)
	gasLimit := uint64(21000)            // gas limit for the transaction
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("failed to suggest gas price: %v", err)
	}

	// Criar a transação
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// Assinar a transação com a chave privada
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send the transaction to the network
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

func GetBlockNumber(network string) (int64, error) {
	client := newClient(NetworkInMemory[network])
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return -1, err
	}
	return block.Number().Int64(), nil
}

func GetBalance(network, from string) (string, error) {
	client := newClient(NetworkInMemory[network])
	balanceWei, err := client.BalanceAt(context.Background(), common.HexToAddress(from), nil)
	if err != nil {
		return "", err
	}
	balance := new(big.Float).Quo(new(big.Float).SetInt(balanceWei), big.NewFloat(1e18))
	return balance.String(), nil
}

func GetTransactionReceipt(network, txHash string) (*types.Receipt, error) {
	client := newClient(NetworkInMemory[network])

	txHash = strings.TrimPrefix(txHash, "0x")
	txHashBytes, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, err
	}
	txHashHex := common.BytesToHash(txHashBytes)
	receipt, err := client.TransactionReceipt(context.Background(), txHashHex)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func SetNetwork(network, url string) {
	NetworkInMemory[network] = url
}

func GetWalletsInMemory() []string {
	wallets := []string{}
	for key := range WalletsInMemory {
		wallets = append(wallets, key)
	}
	return wallets
}
