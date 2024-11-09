package database

import (
	"encoding/json"
	"os"

	"github.com/devworlds/cli-gochain/internal/evm"
)

const walletsFileName = "wallets.json"

// Function to load data from JSON file
func LoadWalletsData() error {
	file, err := os.ReadFile(walletsFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &evm.WalletsInMemory)
}

// Function to save data to JSON file
func SaveWalletsData() error {
	data, err := json.MarshalIndent(evm.WalletsInMemory, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(walletsFileName, data, 0644)
}
