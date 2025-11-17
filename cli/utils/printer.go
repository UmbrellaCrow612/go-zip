package utils

import (
	"fmt"
	"os"
	"time"
)

// PrintStdout prints a message to stdout with the current timestamp.
func PrintStdout(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s\n", timestamp, message)
}

// PrintStderr prints a message to stderr with the current timestamp.
func PrintStderr(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(os.Stderr, "[%s] %s\n", timestamp, message)
}
