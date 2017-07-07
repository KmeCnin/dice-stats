package input

import (
	"flag"

	"errors"

	"github.com/segmentio/go-prompt"
)

// CommandHasNoArgs is error code returned when command is executed without any argument. It will trigger the use of the prompt interface.
const CommandHasNoArgs = "error_1"

// DefaultIterations is the number of iterations to do in order to generate probabilities.
const DefaultIterations = 1000000

// Input returns parsed input data from user command.
type Input struct {
	Query      string
	Iterations int
	Versus     string
}

// ParseCommand validate and get data from command args.
func ParseCommand() (Input, error) {
	i := flag.Int(
		"i",
		DefaultIterations, // 1m
		"number of tries to do to generate stats",
	)
	vs := flag.String(
		"vs",
		"",
		"throw query to try against",
	)

	flag.Parse()

	if 0 == len(flag.Args()) {
		return Input{}, errors.New(CommandHasNoArgs)
	}

	return Input{
		Query:      flag.Args()[0],
		Iterations: *i,
		Versus:     *vs,
	}, nil
}

// ParsePrompt validate and get data from use input.
func ParsePrompt() (Input, error) {
	query := prompt.StringRequired("Throw query")
	vs := prompt.String("Versus throw query (leave empty to ignore)")

	return Input{
		Query:      query,
		Iterations: DefaultIterations,
		Versus:     vs,
	}, nil
}
