package main

import (
	"aoc/file"
	"testing"
)

var fileName string = "input.txt"

func TestTaskOne(t *testing.T) {
	t.Run("task 1", func(t *testing.T) {
		reports := buildTestReports()
		got := taskOne(reports)
		want := 670
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestTaskTwo(t *testing.T) {
	t.Run("task 2", func(t *testing.T) {
		reports := buildTestReports()
		taskOne(reports)
		got := taskTwo(reports)
		want := 700
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func buildTestReports() []Report {
	file := file.Get(fileName)
	defer file.Close()
	reports := buildReports(file)

	return reports
}
