package network

import (
	"go-cli-p/cmd/ping"

	"github.com/spf13/cobra"
)

var NetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Network command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	NetworkCmd.AddCommand(ping.PingCmd)
}
