package tezos

import "os"

// GetTestNodeParams retrieves test node settings
func GetTestNodeParams() (string, string) {
	if os.Getenv("TEZOS_RPC_PORT") != "" {
		return "127.0.0.1", os.Getenv("TEZOS_RPC_PORT")
	}
	return "127.0.0.1", "8732"
}
