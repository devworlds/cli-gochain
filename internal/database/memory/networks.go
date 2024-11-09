package database

import (
	"encoding/json"
	"os"

	"github.com/devworlds/cli-gochain/internal/evm"
)

const networkFileName = "networks.json"

// Function to load data from JSON file
func LoadNetworkData() error {
	file, err := os.ReadFile(networkFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &evm.NetworkInMemory)
}

// Function to save data to JSON file
func SaveNetworkData() error {
	data, err := json.MarshalIndent(evm.NetworkInMemory, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(networkFileName, data, 0644)
}
