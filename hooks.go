package main

import (
	"fmt"
	"go-cli-p/queue"
	"log"

	"github.com/brahma-adshonor/gohook"
)

var (
	q, _ = queue.NewQueue(10)
)

func myPrintf(format string, a ...any) (n int, err error) {
	message := fmt.Sprintf(format, a...)
	q.Enqueue(message)
	fmt.Print(message)
	return
}

func myPrintln(a ...any) (n int, err error) {
	message := fmt.Sprintln(a...)
	q.Enqueue(message)
	fmt.Print(message)
	return
}

func myPrintlnTrampoline(a ...any) (n int, err error) {
	return
}

func myPrintfTrampoline(format string, a ...any) (n int, err error) {
	return
}

// hooks not working gloabally
func addHooks() {
	if err := gohook.Hook(fmt.Println, myPrintln, myPrintlnTrampoline); err != nil {
		log.Fatal(err)
	}
	if err := gohook.Hook(fmt.Printf, myPrintf, myPrintfTrampoline); err != nil {
		log.Fatal(err)
	}

}
