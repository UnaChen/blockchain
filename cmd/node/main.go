package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use:   "node",
		Short: "The Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(runCmd())
	cmd.AddCommand(addTxCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(nodeStatusCmd())

	err := cmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
