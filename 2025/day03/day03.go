package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

const (
	partOneBatteriesPerBank = 2
	partTwoBatteriesPerBank = 12
)

func Part01(banks []string) int {
	return maxJoltage(banks, partOneBatteriesPerBank)
}

func Part02(banks []string) int {
	return maxJoltage(banks, partTwoBatteriesPerBank)
}

func SolveDay03() int {
	f := file.Get("input.txt")
	input := file.ReadLines(f)
	part01 := Part01(input)
	part02 := Part02(input)
	fmt.Println("Part 1:", part01)
	fmt.Println("Part 2:", part02)
	return part01
}

// allowed to turn on exactly k batteries
func maxJoltage(banks []string, k int) int {
	// Resulting "turned on" sequence must be in relative order
	// Does not need to be contiguous
	// Fixed number of kept digits
	sum := 0

	for _, bank := range banks {
		sum += largestLexicographicSubsequence(bank, k)
	}

	return sum
}

func largestLexicographicSubsequence(s string, k int) int {
	if len(s) == k {
		s, _ := strconv.Atoi(s)
		return s
	}
	if strings.TrimLeft(s, "0") == "" {
		return 0
	}

	ret := make([]byte, 0, len(s))
	ret = append(ret, s[0])
	l := len(s) - k

	// monotonic decreasing stack
	// remove smaller digits before adding larger ones
	// smaller digits on the left hurts us more
	for right := 1; right < len(s); right++ {
		for len(ret) > 0 && s[right] > ret[len(ret)-1] && l > 0 {
			l--
			ret = ret[:len(ret)-1]
		}
		ret = append(ret, s[right])
	}

	// we may have too many digits, shrink from the right until we have the desired stack size
	if len(ret) > k {
		ret = ret[:k]
	}

	r, _ := strconv.Atoi(string(ret))
	return r
}
