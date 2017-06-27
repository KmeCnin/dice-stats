package input

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// Throw returns a parsed set of data gave by user.
type Throw struct {
	DiceNumber int
	DiceFaces  int
	KeepNumber int
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
	diceNumber, err := strconv.Atoi(res[0][1])
	if nil != err {
		return Throw{}, err
	}
	diceFaces, err := strconv.Atoi(res[0][2])
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
		DiceNumber: diceNumber,
		DiceFaces:  diceFaces,
		KeepNumber: keepNumber,
	}, nil
}

// Try launches the dice defined in given throw and sum all the results up.
func (t *Throw) Try() int {
	rand.Seed(time.Now().UTC().UnixNano())
	sum := 0
	for i := 0; i < t.DiceNumber; i++ {
		sum += launchDie(t.DiceFaces)
	}
	return sum
}

func launchDie(faces int) int {
	return rand.Intn(faces) + 1
}
