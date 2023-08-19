package utils

import (
	"go-cli-p/cmd/logs"

	"github.com/spf13/cobra"
)

var UtilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	UtilsCmd.AddCommand(logs.LogsCmd)
}
