package ping

import (
	"fmt"
	"go-cli-p/logger"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/spf13/cobra"
)

var (
	wg         sync.WaitGroup
	totalPings int = 5
	Success    atomic.Int64
	Failed     atomic.Int64
)

var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			logger.Println("Zero urls passed.")
			return
		}

		logger.Printf("Pinging each url %v time(s)\n", totalPings)
		for _, url := range args {
			wg.Add(1)
			go ping(url)
		}

		wg.Wait()

		showStatistics()
	},
}

func init() {
	PingCmd.Flags().IntVarP(&totalPings, "pings", "p", totalPings, "Number of time ping needs to run.")
}

func ping(url string) {
	for i := 0; i < totalPings; i++ {
		resp, err := http.Get(url)

		if err != nil {
			logger.Printf("%v - %v\n", url, err)
			Failed.Add(1)
		} else {
			defer resp.Body.Close()
			logger.Printf("%v - %v\n", url, resp.StatusCode)
			Success.Add(1)
		}

		time.Sleep(time.Second * 1)
	}

	wg.Done()
}

func showStatistics() {
	fmt.Print("\n--------- Statistics ---------\n")

	success := Success.Load()
	failure := Failed.Load()

	total := success + failure
	successPercentage := 100 * float64(success) / float64(total)

	logger.Printf("total hits: %v, success percentage %f\n", total, successPercentage)
}
