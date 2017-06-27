package main

import (
	"fmt"
	"os"

	"github.com/kmecnin/dice-stats/input"
	"github.com/kmecnin/dice-stats/stats"
)

func main() {
	// Get user input.
	args := os.Args[1:]

	throw, err := input.NewThrow(args[0])
	if nil != err {
		displayError(err)
	}

	keepMessage := ""
	if throw.KeepNumber > 0 {
		keepMessage = fmt.Sprintf(
			" and keeping %v best dice",
			throw.KeepNumber,
		)
	}
	fmt.Printf(
		"Generating probability distributions using %v dice with %v faces%v...\n",
		throw.DiceNumber,
		throw.DiceFaces,
		keepMessage,
	)

	stats.ProbabilityDistributionOfScore(throw)
}

func displayError(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}
