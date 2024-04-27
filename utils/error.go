package utils

import "fmt"

func Error(message string, line uint32) {
	panic(fmt.Sprintf("ERROR: %s at Line:%v", message, line))
}
