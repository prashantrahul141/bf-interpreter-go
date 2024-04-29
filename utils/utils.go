package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

// Reverses a given string.
// converts them to rune first for even encoding.
// from : https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Reverses an array of any type.
func ReverseArray[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Our panic implementation
func BfigoPanic(message string) {
	devEnv := os.Getenv("DEV")
	if len(devEnv) > 0 {
		log.Error("PANIC :", "message", message)
	}

	fmt.Print(message + "\n")
	os.Exit(1)
}

// Open and read file from command line.
func GetFileContent(logger *log.Logger) string {
	argPassed := os.Args[1:]
	if len(argPassed) <= 0 {
		BfigoPanic(fmt.Sprint("Not enough arguments were passed.\n", USAGE))
	}

	logger.Debug("", "filename", argPassed[0])

	fileContent, err := os.ReadFile(argPassed[0])
	if err != nil {
		BfigoPanic(fmt.Sprintf("File '%s' does not exists.", argPassed[0]))
	}

	return string(fileContent)
}
