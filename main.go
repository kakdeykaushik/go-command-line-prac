package main

import (
	"fmt"
	"go-cli-p/cmd"
	"go-cli-p/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// addHooks()
	handleSIGINT()

	cmd.Execute()
}

func cleanup() {
	fmt.Println("Program killed. Cleaning.")
}

func handleSIGINT() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cleanup()
		logger.Println()
		os.Exit(1)
	}()
}
