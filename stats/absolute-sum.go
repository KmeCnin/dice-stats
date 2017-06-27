package stats

import (
	"fmt"

	"github.com/kmecnin/dice-stats/input"
)

const iterations = 1e6 // 1m

type Distribution struct {
}

func NewDistribution(rawResults map[int]int) (Distribution, error) {

}

// ProbabilityDistributionOfScore returns occurrency probabilities
// for each possible final score.
func ProbabilityDistributionOfScore(throw input.Throw) {
	rawResults := make(map[int]int)
	for i := 0; i < iterations; i++ {
		rawResults[throw.Try()]++
	}
	fmt.Printf("%v", rawResults)
}
