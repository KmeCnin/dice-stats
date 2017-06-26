package input

import (
	"fmt"
	"regexp"
	"strconv"
)

// Throw returns a parsed set of data gave by user.
type Throw struct {
	DicesNumber int
	DicesFaces  int
	KeepNumber  int
}

// NewThrow creates and returns a parsed input from data raw input.
func NewThrow(query string) (Throw, error) {
	// Create regexp.
	re, err := regexp.Compile(`^(\d+)d(\d+)k?(\d+)?`)
	if nil != err {
		return Throw{}, err
	}

	// Apply regexp.
	res := re.FindAllStringSubmatch(query, -1)
	if len(res) == 0 {
		return Throw{}, fmt.Errorf("given query `%v` is badly formatted", query)
	}

	// Cast values to int.
	dicesNumber, err := strconv.Atoi(res[0][1])
	if nil != err {
		return Throw{}, err
	}
	dicesFaces, err := strconv.Atoi(res[0][2])
	if nil != err {
		return Throw{}, err
	}
	keepNumber := 0
	if "" != res[0][3] {
		keepNumber, err = strconv.Atoi(res[0][3])
		if nil != err {
			return Throw{}, err
		}
	}

	return Throw{
		DicesNumber: dicesNumber,
		DicesFaces:  dicesFaces,
		KeepNumber:  keepNumber,
	}, nil
}
