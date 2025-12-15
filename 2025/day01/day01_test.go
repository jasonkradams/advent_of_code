package day01_test

import (
	"reflect"
	"testing"

	"github.com/jasonkradams/advent_of_code/2025/day01"
)

func TestDay01(t *testing.T) {
	t.Run("example input clicks past zero 3 times", func(t *testing.T) {
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

		want := day01.Answer{LandOnZero: 3, ClicksPastZero: 6}
		got := day01.Day01(rotations)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("real input lands on 0 1086 times", func(t *testing.T) {
		const (
			wantLandOnZero     = 1086
			wantClicksPastZero = 6268
		)

		wrongAnswers := []int{
			6294,
		}
		want := wantLandOnZero
		got := day01.PartOne()

		if got.LandOnZero != want {
			t.Errorf("got %d want %d", got, want)
		}

		for _, wrong := range wrongAnswers {
			if got.ClicksPastZero == wrong {
				t.Fatalf("already tried wrong answer: %d", got.ClicksPastZero)
			}
		}

		if got.ClicksPastZero != wantClicksPastZero {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
