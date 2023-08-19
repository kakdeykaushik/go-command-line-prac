package cmd

import (
	"fmt"
	"go-cli-p/cmd/network"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-p",
	Short: "CLI for basic stuff",
	Long:  `CLI for TODO, few utilities and network operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(network.NetworkCmd)
}
