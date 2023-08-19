package logs

import (
	"fmt"
	"go-cli-p/logger"
	"os"

	"github.com/spf13/cobra"
)

const LOG_FILE = "log.log"

var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		viewLog, _ := cmd.Flags().GetBool("view")
		if viewLog {
			printLogs()
		}

		clearLog, _ := cmd.Flags().GetBool("clear")
		if clearLog {
			clearLogs()
		}
	},
}

func init() {
	LogsCmd.Flags().BoolP("view", "v", false, "view logs")
	LogsCmd.Flags().BoolP("clear", "c", false, "clear logs")
}

func printLogs() {
	dataByte, err := os.ReadFile(LOG_FILE)

	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(string(dataByte))

}

func clearLogs() {
	// os.Create can be used for truncate. docs - https://pkg.go.dev/os#Create:~:text=Create%20creates%20or%20truncates%20the%20named%20file.%20If%20the%20file%20already%20exists%2C%20it%20is%20truncated

	_, err := os.Create(LOG_FILE)

	if err != nil {
		fmt.Println("Unable to clear logs")
	}

	fmt.Println("Logs cleared")

}
