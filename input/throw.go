package input

import (
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
)

// Throw returns a parsed set of data gave by user.
type Throw struct {
	DiceNumber int
	DiceFaces  int
	KeepNumber int
}

// GetThrow creates and returns a parsed input from data raw input.
func GetThrow(query string) (Throw, error) {
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
func (t *Throw) Try(r *rand.Rand) int {
	launches := make([]int, t.DiceNumber)
	for i := 0; i < t.DiceNumber; i++ {
		launches[i] = launchDie(t.DiceFaces, r)
	}

	if t.KeepNumber > 0 {
		sort.Sort(sort.Reverse(sort.IntSlice(launches)))
		launches = launches[:t.KeepNumber]
	}

	sum := 0
	for _, i := range launches {
		sum += i
	}

	return sum
}

func launchDie(faces int, r *rand.Rand) int {
	return r.Intn(faces) + 1
}
