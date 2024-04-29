package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func getLevel() log.Level {
	devEnv := os.Getenv("DEV")
	if len(devEnv) > 0 {
		return log.InfoLevel
	}

	return log.ErrorLevel

}

func GetGlobalLogger() *log.Logger {
	logger := log.New(os.Stderr)
	logger.SetLevel(getLevel())
	return logger
}

func Error(message string, line uint32) {
	BfigoPanic(fmt.Sprintf("%s at line %d", message, line))
}
