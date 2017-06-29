package stats

import (
	"github.com/kmecnin/dice-stats/input"
	pb "gopkg.in/cheggaaa/pb.v1"
)

// DistributionOfScore returns occurrency probabilities
// for each possible final score.
func DistributionOfScore(throw input.Throw, iterations int) map[int]int {
	progress := pb.StartNew(iterations)

	statistics := make(map[int]int)
	for i := 0; i < iterations; i++ {
		statistics[throw.Try()]++
		progress.Increment()
	}

	probabilities := make(map[int]int)
	for score, occurrencies := range statistics {
		pct := int((float32(occurrencies) / float32(iterations)) * 100)
		if pct > 0 {
			probabilities[score] = pct
		}
	}

	progress.Finish()
	return probabilities
}
