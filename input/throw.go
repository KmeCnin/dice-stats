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
	launches   []int
}

// GetThrow creates and returns a parsed input from data raw input.
func GetThrow(query string) (Throw, error) {
	// Create regexp.
	re, err := regexp.Compile(`^(\d+)d(\d+)(k(\d+))?$`)
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
		keepNumber, err = strconv.Atoi(res[0][4])
		if nil != err {
			return Throw{}, err
		}
	}

	if keepNumber > diceNumber {
		keepNumber = diceNumber
	}

	return Throw{
		DiceNumber: diceNumber,
		DiceFaces:  diceFaces,
		KeepNumber: keepNumber,
		launches:   make([]int, diceNumber),
	}, nil
}

// Try launches the dice defined in given throw and sum all the results up.
func (t *Throw) Try(r *rand.Rand) int {
	if t.KeepNumber == 0 {
		return t.simpleTry(r)
	}
	return t.bestTry(r)
}

func (t *Throw) simpleTry(r *rand.Rand) int {
	sum := 0
	for i := 0; i < t.DiceNumber; i++ {
		sum += launchDie(t.DiceFaces, r)
	}
	return sum
}

func (t *Throw) bestTry(r *rand.Rand) int {
	for i := 0; i < t.DiceNumber; i++ {
		t.launches[i] = launchDie(t.DiceFaces, r)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(t.launches)))
	sum := 0
	for _, i := range t.launches[:t.KeepNumber] {
		sum += i
	}
	return sum
}

func launchDie(faces int, r *rand.Rand) int {
	return r.Intn(faces) + 1
}
