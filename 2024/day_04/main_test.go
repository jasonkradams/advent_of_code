package main

import (
	"aoc/file"
	"testing"
)

var grid = [][]rune{
	[]rune("MMMSXXMASM"),
	[]rune("MSAMXMSMSA"),
	[]rune("AMXSXMAAMM"),
	[]rune("MSAMASMSMX"),
	[]rune("XMASAMXAMM"),
	[]rune("XXAMMXXAMA"),
	[]rune("SMSMSASXSS"),
	[]rune("SAXAMASAAA"),
	[]rune("MAMMMXMMMM"),
	[]rune("MXMXAXMASX"),
}

var fileName string = "input.txt"

func TestTaskOne(t *testing.T) {
	file := file.Get(filePath)
	defer file.Close()
	t.Run("task 1", func(t *testing.T) {

		got := taskOne(grid)
		want := 18
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

// func TestTaskTwo(t *testing.T) {
// 	file := file.Get(filePath)
// 	defer file.Close()
// 	t.Run("task 2", func(t *testing.T) {

// 		got := taskTwo()
// 		want := true
// 		if got != want {
// 			t.Errorf("got %t want %t", got, want)
// 		}
// 	})
// }
