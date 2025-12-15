package main

import (
	"testing"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

var fileName string = "input.txt"

func TestTaskOne(t *testing.T) {
	file := file.Get(filePath)
	expressions := buildExpressions(file)
	t.Run("task 1", func(t *testing.T) {

		got := taskOne(expressions)
		want := 183380722
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestTaskTwo(t *testing.T) {
	file := file.Get(filePath)
	expressions := buildExpressions(file)
	t.Run("task 1", func(t *testing.T) {

		got := taskTwo(expressions)
		want := 82733683
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
