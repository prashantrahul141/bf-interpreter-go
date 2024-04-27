package utils

import (
	"os"

	"github.com/charmbracelet/log"
)

func GetGlobalLogger() *log.Logger {
	logger := log.New(os.Stderr)
	logger.SetLevel(log.DebugLevel)
	return logger
}

func Error(message string, line uint32) {
	log.Error("'%s' at %d", message, line)

}
