package input

import (
	"flag"
)

// Input returns parsed input data from user command.
type Input struct {
	Query      string
	Iterations int
}

// ParseCommand validate and get data from user input.
func ParseCommand() (Input, error) {
	i := flag.Int(
		"i",
		1000000,
		"number of tries to do to generate stats",
	)

	flag.Parse()

	return Input{
		Query:      flag.Args()[0],
		Iterations: *i,
	}, nil
}
