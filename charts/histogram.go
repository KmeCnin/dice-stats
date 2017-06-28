package charts

import (
	"fmt"
	"sort"
	"strconv"
)

// DrawProbabilitiesHistogram displays probabilities as a histogram chart.
func DrawProbabilitiesHistogram(p map[int]int) {
	var keys []int
	max := 0
	for k, freq := range p {
		keys = append(keys, k)
		if max < freq {
			max = freq
		}
	}
	sort.Ints(keys)
	fmt.Print("\n")
	for i := max; i > 0; i-- {
		// Print Y axes.
		fmt.Printf("%v%%", indentInt(i, 3))
		for _, k := range keys {
			if p[k] == i {
				fmt.Print(" ▄▄")
			} else if p[k] > i {
				fmt.Print(" ██")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Print("\n")
	}
	// Print X axes.
	fmt.Print("    ")
	for _, k := range keys {
		fmt.Printf("%v", indentInt(k, 3))
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
