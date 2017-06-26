package main

import (
	"fmt"
	"os"

	"github.com/kmecnin/dices-stats/input"
)

func main() {
	// Get user input.
	args := os.Args[1:]

	throw, err := input.NewThrow(args[0])
	if nil != err {
		handleError(err)
	}

	fmt.Println(throw)
}

func handleError(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}
