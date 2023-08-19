package main

import (
	"fmt"
	"go-cli-p/cmd"
	"go-cli-p/cmd/ping"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	addHooks()
	handleSIGINT()

	cmd.Execute()

	showStatistics()
}

func showStatistics() {

	fmt.Print("\n--------- Statistics ---------\n")

	success := ping.Success.Load()
	failure := ping.Failed.Load()

	total := success + failure
	successPercentage := 100 * float64(success) / float64(total)

	fmt.Printf("total hits: %v, success percentage %f\n", total, successPercentage)
}

func handleSIGINT() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		showStatistics()
		fmt.Println()
		os.Exit(1)
	}()
}
