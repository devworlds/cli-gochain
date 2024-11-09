package evm

// WalletsInMemory is a map to store the private keys of the ethereum wallets.
var WalletsInMemory = make(map[string]string)

// NetworkMemory is a map to store the network node of the ethereum client.
var NetworkInMemory = make(map[string]string)
