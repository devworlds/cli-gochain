package main

import (
	"fmt"
	"math/big"

	inmemory "github.com/devworlds/cli-gochain/internal/database/memory"
	"github.com/devworlds/cli-gochain/internal/evm"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gochain",
		Short: "A blockchain implementation in Go",
		Long:  `A blockchain implementation in Go`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	generateWalletCmd := &cobra.Command{
		Use:   "generate-wallet",
		Short: "Generate an ethereum wallet",
		Long:  `Generate an ethereum wallet`,
		Run: func(cmd *cobra.Command, args []string) {
			inmemory.LoadWalletsData()
			wallet, err := evm.GenerateWallet()
			if err != nil {
				fmt.Println(err)
				return
			}
			inmemory.SaveWalletsData()
			fmt.Printf("PublicKey: %s\nPrivateKey: %s\n", wallet, evm.WalletsInMemory[wallet])
		},
	}

	getWalletsCmd := &cobra.Command{
		Use:   "get-wallets-inmemory",
		Short: "Get all ethereum wallets in memory",
		Long:  `Get all ethereum wallets in memory`,
		Run: func(cmd *cobra.Command, args []string) {
			inmemory.LoadWalletsData()
			wallets := evm.GetWalletsInMemory()
			fmt.Printf("Ethereum wallets in memory: %v\n", wallets)
		},
	}

	setNetworkCmd := &cobra.Command{
		Use:   "set-network-node",
		Short: "Set the node network to use on EthereumClient",
		Long:  `Set the node network to use on EthereumClient`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Please provide the network and nodeUrl")
				return
			}
			inmemory.LoadNetworkData()
			network := args[0]
			nodeUrl := args[1]
			evm.SetNetwork(network, nodeUrl)
			inmemory.SaveNetworkData()
			fmt.Printf("Network %v successfully added\n", network)
		},
	}

	getBlockNumberCmd := &cobra.Command{
		Use:   "get-blocknumber",
		Short: "Get the current block number",
		Long:  `Get the current block number`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide the network")
				return
			}
			inmemory.LoadNetworkData()
			blockNumber, err := evm.GetBlockNumber(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Current blockNumber: %d\n", blockNumber)
		},
	}

	createTxCmd := &cobra.Command{
		Use:   "create-transaction",
		Short: "Create a transaction",
		Long:  `Create a transaction`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 4 {
				fmt.Println("Please provide the network, from, to, amount")
				return
			}
			inmemory.LoadNetworkData()
			inmemory.LoadWalletsData()
			network := args[0]
			from := args[1]
			to := args[2]
			amount, _ := new(big.Int).SetString(args[3], 10)
			txHash, err := evm.CreateTransaction(network, evm.WalletsInMemory[from], to, amount.Int64())
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Transaction created successfully with hash: %s\n", txHash)
		},
	}

	getBalanceCmd := &cobra.Command{
		Use:   "get-balance",
		Short: "Get the balance of an ethereum address",
		Long:  `Get the balance of an ethereum address`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Please provide the network and address")
				return
			}
			inmemory.LoadNetworkData()
			balance, err := evm.GetBalance(args[0], args[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Balance: %s\n", balance)
		},
	}

	getTransactionReceiptCmd := &cobra.Command{
		Use:   "get-transaction-receipt",
		Short: "Get the transaction receipt of a transaction",
		Long:  `Get the transaction receipt of a transaction`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("Please provide the network and transaction hash")
				return
			}
			inmemory.LoadNetworkData()
			receipt, err := evm.GetTransactionReceipt(args[0], args[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Transaction receipt: %+v\n", receipt)
		},
	}

	rootCmd.AddCommand(
		generateWalletCmd,
		getWalletsCmd,
		setNetworkCmd,
		getBlockNumberCmd,
		createTxCmd,
		getBalanceCmd,
		getTransactionReceiptCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
