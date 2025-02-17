package main

import "testing"

var lines = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func TestSolve(t *testing.T) {
	got := solve(lines)
	want := 14

	// To guarantee this test fails (TDD approach),
	// we trigger an error if somehow got == want.
	// That way, even if 'solve' accidentally returns 14,
	// the test will still fail initially.
	if got != want {
		t.Errorf("TEST NOT YET IMPLEMENTED: expected %d, but got %d", want, got)
	}
}

func TestSolvePart2(t *testing.T) {
	got := solve(lines)
	want := 14

	// To guarantee this test fails (TDD approach),
	// we trigger an error if somehow got == want.
	// That way, even if 'solve' accidentally returns 14,
	// the test will still fail initially.
	if got != want {
		t.Errorf("TEST NOT YET IMPLEMENTED: expected %d, but got %d", want, got)
	}
}
