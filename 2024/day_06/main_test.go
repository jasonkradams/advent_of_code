package main

import (
	"testing"
)

var grid = [][]rune{
	[]rune("....#....."),
	[]rune(".........#"),
	[]rune(".........."),
	[]rune("..#......."),
	[]rune(".......#.."),
	[]rune(".........."),
	[]rune(".#..^....."),
	[]rune("........#."),
	[]rune("#........."),
	[]rune("......#..."),
}

func TestTaskOne(t *testing.T) {
	t.Run("task 1", func(t *testing.T) {

		got := taskOne(grid)
		want := 41
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

// func TestTaskTwo(t *testing.T) {
// 	t.Run("task 2", func(t *testing.T) {
// 		got := taskTwo()
// 		want := 123
// 		if got != want {
// 			t.Errorf("got %d want %d", got, want)
// 		}
// 	})
// }
