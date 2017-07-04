package stats

import (
	"github.com/kmecnin/dice-stats/src/input"
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

// VersusProbabilities hold all probabilities from versus.
type VersusProbabilities struct {
	Win  float64
	Lose float64
	Draw float64
}

// DistributionOfWin returns percent of chances to win/lose/draw for
// the player1 using throw1 against player2 using throw2.
func DistributionOfWin(throw1, throw2 input.Throw, iterations int) VersusProbabilities {
	progress := pb.StartNew(iterations)

	winStats := 0
	loseStats := 0
	drawStats := 0
	for i := 0; i < iterations; i++ {
		diff := throw1.Try() - throw2.Try()
		if diff > 0 {
			winStats++
		} else if diff < 0 {
			loseStats++
		} else {
			drawStats++
		}
		progress.Increment()
	}

	progress.Finish()
	return VersusProbabilities{
		(float64(winStats) / float64(iterations)) * 100,
		(float64(loseStats) / float64(iterations)) * 100,
		(float64(drawStats) / float64(iterations)) * 100,
	}
}
