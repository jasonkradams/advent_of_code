package day02_test

import (
	"testing"

	"github.com/jasonkradams/advent_of_code/2025/day02"
)

func TestDay02(t *testing.T) {

	t.Run("sums total of all invalid IDs", func(t *testing.T) {
		invalidIDs := []uint64{11, 22, 99, 1010, 1188511885, 222222, 446446, 38593859}
		want := uint64(1227775554)
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
		want := uint64(1227775554)

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("tests Day02Part01", func(t *testing.T) {
		got := day02.PartOne()
		want := uint64(23039913998)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
