package main

import (
	"os"
)

func main() {
	cmd := getRootCmd(handler)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
