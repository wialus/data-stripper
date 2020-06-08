package main

import "fmt"

var (
	Version = "0.0.0-dev"
)

func handlerVersion() error {
	fmt.Println(Version)

	return nil
}
