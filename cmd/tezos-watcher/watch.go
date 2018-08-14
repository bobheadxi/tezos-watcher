package main

import (
	"encoding/json"
	"fmt"
	"time"

	watcher "github.com/bobheadxi/tezos-watcher"
	"github.com/bobheadxi/tezos-watcher/cmd/tezos-watcher/internal"
	"github.com/spf13/cobra"
)

func init() {
	cmdWatchChain.Flags().String("host", "127.0.0.1", "host of node")
	cmdWatchChain.Flags().String("port", "8732", "port of node")
	cmdWatchChain.Flags().StringP("chain", "c", "main", "chain to watch")
	cmdWatchChain.Flags().StringP("block", "b", "head", "block to watch")
	cmdWatchChain.Flags().IntP("rate", "r", 1, "rate to poll for status in seconds")
	cmdRoot.AddCommand(cmdWatchChain)
}

var cmdWatchChain = &cobra.Command{
	Use:   "watch-chain",
	Short: "Watch node chain",
	Run: func(cmd *cobra.Command, args []string) {
		// Set up connection
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		w, err := watcher.New(watcher.ConnectOpts{Host: host, Port: port})
		if err != nil {
			internal.Fail(err.Error())
		}

		chain, _ := cmd.Flags().GetString("chain")
		block, _ := cmd.Flags().GetString("block")
		rate, _ := cmd.Flags().GetInt("rate")
		quit := make(chan struct{})
		statusCh, errCh := w.WatchBlock(watcher.BlockOptions{
			Chain: chain, Block: block, Rate: time.Duration(rate) * time.Second}, quit)

		// Start watching output
		fmt.Printf("Watching %s/%s on %s:%s...\n", chain, block, host, port)
		for {
			select {
			case err := <-errCh:
				if err != nil {
					close(quit)
					internal.Fail(err.Error())
				}
			case event := <-statusCh:
				b, err := json.Marshal(event)
				if err != nil {
					close(quit)
					internal.Fail(err.Error())
				}
				println(string(b))
			}
		}
	},
}
