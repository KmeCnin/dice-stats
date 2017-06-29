package main

import (
	"fmt"
	"os"

	"github.com/kmecnin/dice-stats/charts"
	"github.com/kmecnin/dice-stats/input"
	"github.com/kmecnin/dice-stats/stats"
)

func main() {
	// Get user input.
	userInput, err := input.ParseCommand()
	if nil != err {
		displayError(err)
	}

	throw, err := input.GetThrow(userInput.Query)
	if nil != err {
		displayError(err)
	}

	displayMsgGeneration(throw)
	proba := stats.DistributionOfScore(throw, userInput.Iterations)
	charts.DrawProbabilitiesHistogram(proba)
}

func displayError(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

func displayMsgGeneration(throw input.Throw) {
	keepMessage := ""
	if throw.KeepNumber > 0 {
		keepMessage = fmt.Sprintf(
			" and keeping %v best dice",
			throw.KeepNumber,
		)
	}
	fmt.Printf(
		"Generating probabilities distributions using %v dice with %v faces%v\n",
		throw.DiceNumber,
		throw.DiceFaces,
		keepMessage,
	)
}
