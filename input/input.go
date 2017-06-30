package input

import (
	"flag"
)

// Input returns parsed input data from user command.
type Input struct {
	Query      string
	Iterations int
	Versus     string
}

// ParseCommand validate and get data from user input.
func ParseCommand() (Input, error) {
	i := flag.Int(
		"i",
		1000000, // 1m
		"number of tries to do to generate stats",
	)
	vs := flag.String(
		"vs",
		"",
		"throw query to try against",
	)

	flag.Parse()

	return Input{
		Query:      flag.Args()[0],
		Iterations: *i,
		Versus:     *vs,
	}, nil
}
