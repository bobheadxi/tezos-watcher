package internal

import "os"

// Fail prints message and exits
func Fail(msg string) {
	println(msg)
	os.Exit(1)
}
