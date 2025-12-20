package day03_test

import (
	"testing"

	"github.com/jasonkradams/advent_of_code/2025/day03"
)

const (
	partOneExampleAnswer = 357
	partTwoExampleAnswer = 3121910778619
	partOneAnswer        = 17412
)

func TestDay03(t *testing.T) {
	t.Run("test example input", func(t *testing.T) {
		input := []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		}

		part01 := day03.Part01(input)

		if part01 != partOneExampleAnswer {
			t.Errorf("got %d, want %d", part01, partOneExampleAnswer)
		}

		part02 := day03.Part02(input)

		if part02 != partTwoExampleAnswer {
			t.Errorf("got %d, want %d", part02, partTwoExampleAnswer)
		}
	})

	t.Run("test input.txt", func(t *testing.T) {
		got := day03.SolveDay03()
		if got != partOneAnswer {
			t.Errorf("got %d, want %d", got, partOneAnswer)
		}
	})
}
