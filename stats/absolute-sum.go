package stats

import (
	"fmt"

	"github.com/kmecnin/dice-stats/charts"
	"github.com/kmecnin/dice-stats/input"
)

const iterations = 100000 // 1m

// ProbabilityDistributionOfScore returns occurrency probabilities
// for each possible final score.
func ProbabilityDistributionOfScore(throw input.Throw) {
	fmt.Printf("%v\n", iterations)
	statistics := make(map[int]int)
	for i := 0; i < iterations; i++ {
		statistics[throw.Try()]++
	}
	fmt.Printf("%v\n", statistics)

	probabilities := make(map[int]int)
	for score, occurrencies := range statistics {
		probabilities[score] = int((float32(occurrencies) / float32(iterations)) * 100)
	}
	fmt.Printf("%v\n", probabilities)
	charts.DrawProbabilitiesHistogram(probabilities)
}
