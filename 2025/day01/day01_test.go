package day01_test

import (
	"testing"

	"github.com/jasonkradams/advent_of_code/2025/day01"
)

func TestDay01Example(t *testing.T) {
	rotations := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	want := 3
	got := day01.Day01(rotations)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDay01PartOne(t *testing.T) {
	want := 1086
	got := day01.PartOne()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
