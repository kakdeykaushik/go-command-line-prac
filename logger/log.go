package logger

import (
	"fmt"
	"os"
)

const LOG_FILE = "log.log"

func appendToFile(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}

func Println(a ...any) {
	message := fmt.Sprintln(a...)
	appendToFile(LOG_FILE, message)
	fmt.Println(a...)
}

func Printf(format string, a ...any) {
	message := fmt.Sprintf(format, a...)
	appendToFile(LOG_FILE, message)
	fmt.Printf(format, a...)
}

func Fatal(v ...any) {
	message := fmt.Sprintln(v...)
	appendToFile(LOG_FILE, message)
	fmt.Print(message)
	os.Exit(1)
}
