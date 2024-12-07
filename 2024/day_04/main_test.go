package main

import (
	"aoc/file"
	"testing"
)

var fileName string = "input.txt"

func TestTaskOne(t *testing.T) {
	file := file.Get(filePath)
	defer file.Close()
	t.Run("task 1", func(t *testing.T) {

		got := taskOne()
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestTaskTwo(t *testing.T) {
	file := file.Get(filePath)
	defer file.Close()
	t.Run("task 1", func(t *testing.T) {

		got := taskTwo()
		want := true
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
