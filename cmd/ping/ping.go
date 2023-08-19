package ping

import (
	"fmt"
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
		fmt.Printf("Pinging each url %v time(s)\n", totalPings)
		for _, url := range args {
			wg.Add(1)
			go ping(url)
		}

		wg.Wait()
	},
}

func init() {
	PingCmd.Flags().IntVarP(&totalPings, "pings", "p", totalPings, "Number of time ping needs to run.")
}

func ping(url string) {
	for i := 0; i < totalPings; i++ {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Printf("%v - %v\n", url, err)
			Failed.Add(1)
		} else {
			defer resp.Body.Close()
			fmt.Printf("%v - %v\n", url, resp.StatusCode)
			Success.Add(1)
		}

		time.Sleep(time.Second * 1)
	}

	wg.Done()
}
