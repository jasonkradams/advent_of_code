package main

import "testing"

var lines string = "2333133121414131402"

func TestTask1(t *testing.T) {
	got := solveTask1(lines)
	want := 1928

	if got != want {
		t.Errorf("expected %d, but got %d", want, got)
	}
}
