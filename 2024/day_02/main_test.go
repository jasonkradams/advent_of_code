package main

import (
	"aoc/file"
	"testing"
)

var fileName string = "input.txt"

func TestTaskOne(t *testing.T) {
	t.Run("task 1", func(t *testing.T) {
		// Open the file
		file := file.Get(fileName)
		defer file.Close()
		reports := buildReports(file)

		// safeCount := taskOne(reports)
		got := taskOne(reports)
		want := 670
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
