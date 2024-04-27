package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func GetGlobalLogger() *log.Logger {
	logger := log.New(os.Stderr)
	logger.SetLevel(log.DebugLevel)
	return logger
}

func Error(message string, line uint32) {
	BfigoPanic(fmt.Sprintf("%s at line %d", message, line))
}
