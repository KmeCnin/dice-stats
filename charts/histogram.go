package charts

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/kmecnin/dice-stats/stats"
)

// DrawProbabilitiesHistogramWin displays probabilities as a histogram chart.
func DrawProbabilitiesHistogramWin(p stats.VersusProbabilities) {
	fmt.Println()
	win := int(p.Win)
	draw := int(p.Draw)
	fmt.Printf("Win: %.2f%%\t", p.Win)
	fmt.Printf("Draw: %.2f%%\t", p.Draw)
	fmt.Printf("Lose: %.2f%%\n", p.Lose)
	for j := 0; j < 2; j++ {
		for i := 0; i <= win; i++ {
			fmt.Print("█")
		}
		for i := win; i <= win+draw; i++ {
			fmt.Print(" ")
		}
		for i := win + draw; i <= 100; i++ {
			fmt.Print("░")
		}
		fmt.Println()
	}
	fmt.Println()
}

// DrawProbabilitiesHistogramScore displays probabilities as a histogram chart.
func DrawProbabilitiesHistogramScore(p map[int]int) {
	var keys []int
	maxY := 0
	maxX := 0
	for k, freq := range p {
		keys = append(keys, k)
		if maxY < freq {
			maxY = freq
		}
		if maxX < k {
			maxX = k
		}
	}
	sort.Ints(keys)
	fmt.Println()
	digits := int(math.Log10(float64(maxX))) + 1
	for i := maxY; i > 0; i-- {
		// Print Y axes.
		fmt.Printf("%v%%", indentInt(i, 3))
		for _, k := range keys {
			fmt.Print(" ")
			if p[k] == i {
				for j := 0; j < digits; j++ {
					fmt.Print("▄")
				}
			} else if p[k] > i {
				for j := 0; j < digits; j++ {
					fmt.Print("█")
				}
			} else {
				for j := 0; j < digits; j++ {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")
	}
	// Print X axes.
	fmt.Print("    ")
	for _, k := range keys {
		fmt.Printf("%v", indentInt(k, digits+1))
	}
	fmt.Println()
}

func indentInt(n int, length int) string {
	str := ""
	prev := 0
	zero := false
	for pos := length - 1; pos >= 0; pos-- {
		div := 1
		for j := 0; j < pos; j++ {
			div *= 10
		}
		n -= prev * div * 10
		i := n / div
		prev = i
		if i > 0 || zero {
			zero = true
			str += strconv.Itoa(i)
		} else {
			str += " "
		}
	}
	return str
}

func indentFloat(f float64, length int) string {
	str := fmt.Sprintf("%.2f", f)
	for i := len(str); i < length; i++ {
		str = fmt.Sprintf(" %v", str)
	}

	return str
}
