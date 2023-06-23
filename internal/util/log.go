package util

import "fmt"

func Log(logPrefix string, message string) {
	fmt.Printf("[%s] %s\n", logPrefix, message)
}
