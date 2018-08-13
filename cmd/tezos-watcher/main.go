package main

import (
	"github.com/bobheadxi/tezos-watcher/cmd/tezos-watcher/internal"
	"github.com/spf13/cobra"
)

var (
	cmdRoot = &cobra.Command{
		Use:   "tezos-watcher",
		Short: "A tool for watching Tezos nodes",
	}
)

func main() {
	if err := cmdRoot.Execute(); err != nil {
		internal.Fail(err.Error())
	}
}
