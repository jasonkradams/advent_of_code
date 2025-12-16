package day02_test

import (
	"testing"

	"github.com/jasonkradams/advent_of_code/2025/day02"
)

const (
	partOneExampleAnswer = uint64(1227775554)
	partTwoExampleAnswer = uint64(4174379265)
	partOneInputAnswer   = uint64(23039913998)
	partTwoInputAnswer   = uint64(35950619148)
)

func TestDay02(t *testing.T) {

	t.Run("sums total of all invalid IDs", func(t *testing.T) {
		invalidIDs := []uint64{11, 22, 99, 1010, 1188511885, 222222, 446446, 38593859}
		want := uint64(partOneExampleAnswer)
		got := day02.InvalidIDsTotal(invalidIDs)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("splits input by comma", func(t *testing.T) {
		// 	input := []string{`11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
		// 1698522-1698528,446443-446449,38593856-38593862,565653-565659,
		// 824824821-824824827,2121212118-2121212124`}

		input := []string{
			"11-22",
			"95-115",
			"998-1012",
			"1188511880-1188511890",
			"222220-222224",
			"1698522-1698528",
			"446443-446449",
			"38593856-38593862",
			"565653-565659",
			"824824821-824824827",
			"2121212118-2121212124",
		}

		got := day02.Day02(input)
		want := uint64(partOneExampleAnswer)

		if got.PartOne.Answer != want {
			t.Errorf("got %d, want %d", got, want)
		}

		want = partTwoExampleAnswer
		if got.PartTwo.Answer != want {
			t.Errorf("part2: got %d, want %d", got.PartTwo.Answer, want)
		}
	})

	t.Run("tests Day02Part01", func(t *testing.T) {
		got := day02.SolveDayTwo()
		want := uint64(partOneInputAnswer)
		if got.PartOne.Answer != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	invalidAnswers := []uint64{
		partOneInputAnswer,
		24208750487,      // too low
		3645584041103490, // too high
		8638101580253925,
	}

	t.Run("tests Day02Part02", func(t *testing.T) {
		got := day02.SolveDayTwo()

		for _, a := range invalidAnswers {
			if a == got.PartTwo.Answer {
				t.Fatalf("already tried wrong number %d", got.PartTwo.Answer)
			}
		}
		want := uint64(partTwoInputAnswer)
		if got.PartTwo.Answer != want {
			t.Errorf("got %d want %d", got.PartTwo.Answer, want)
		}
	})

}
