package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/segmentio/go-prompt"

	"github.com/kmecnin/dice-stats/charts"
	"github.com/kmecnin/dice-stats/input"
	"github.com/kmecnin/dice-stats/stats"
)

func main() {
	f, err := os.Create("/tmp/prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	isPrompt := false
	userInput, err := input.ParseCommand()
	if nil != err {
		if err.Error() == input.CommandHasNoArgs {
			userInput, err = input.ParsePrompt()
			isPrompt = true
		}
		if nil != err {
			displayError(err)
		}
	}

	if userInput.Versus == "" {
		scoreDistribution(userInput)
	} else {
		winDistribution(userInput)
	}

	if isPrompt {
		prompt.String("Press ENTER to exit")
	}
}

func winDistribution(userInput input.Input) {
	throw1, err := input.GetThrow(userInput.Query)
	if nil != err {
		displayError(err)
	}
	throw2, err := input.GetThrow(userInput.Versus)
	if nil != err {
		displayError(err)
	}

	displayMsgWin(throw1, throw2)
	proba := stats.DistributionOfWin(throw1, throw2, userInput.Iterations)
	charts.DrawProbabilitiesHistogramWin(proba)
}

func scoreDistribution(userInput input.Input) {
	throw, err := input.GetThrow(userInput.Query)
	if nil != err {
		displayError(err)
	}

	displayMsgScore(throw)
	proba := stats.DistributionOfScore(throw, userInput.Iterations)
	charts.DrawProbabilitiesHistogramScore(proba)
}

func displayError(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}

func displayMsgScore(throw input.Throw) {
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

func displayMsgWin(throw1, throw2 input.Throw) {
	fmt.Printf("Generating probabilities distributions:\n")
	keepMessage1 := ""
	if throw1.KeepNumber > 0 {
		keepMessage1 = fmt.Sprintf(
			" and keeping %v best dice",
			throw1.KeepNumber,
		)
	}
	fmt.Printf(
		"\tPlayer1: %v dice with %v faces%v\n",
		throw1.DiceNumber,
		throw1.DiceFaces,
		keepMessage1,
	)
	keepMessage2 := ""
	if throw2.KeepNumber > 0 {
		keepMessage2 = fmt.Sprintf(
			" and keeping %v best dice",
			throw1.KeepNumber,
		)
	}
	fmt.Printf(
		"vs\tPlayer2: %v dice with %v faces%v\n",
		throw2.DiceNumber,
		throw2.DiceFaces,
		keepMessage2,
	)
}
