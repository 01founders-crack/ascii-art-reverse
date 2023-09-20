package main

import (
	"os"
	"strings"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 1 {
		return
	}
	if !strings.Contains(arg[0], "--reverse=") {
		return
	}
}