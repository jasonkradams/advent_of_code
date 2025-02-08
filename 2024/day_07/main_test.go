package main

import (
	"testing"
)

var calibrations = []Calibration{
	{result: 190, expression: []int{10, 19}},
	{result: 3267, expression: []int{81, 40, 27}},
	{result: 83, expression: []int{17, 5}},
	{result: 156, expression: []int{15, 6}},
	{result: 7290, expression: []int{6, 8, 6, 15}},
	{result: 161011, expression: []int{16, 10, 13}},
	{result: 192, expression: []int{17, 8, 14}},
	{result: 21037, expression: []int{9, 7, 18, 13}},
	{result: 292, expression: []int{11, 6, 16, 20}},
}

func TestTaskOne(t *testing.T) {
	t.Run("task 1", func(t *testing.T) {
		got := taskOne(calibrations)
		want := 3749
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

// Test Task Two
func TestTaskTwo(t *testing.T) {
	var testCases = []struct {
		name   string
		expr   []int
		target int
		want   bool
	}{
		{
			name:   "190: 10 19",
			expr:   []int{10, 19},
			target: 190,
			want:   true,
		},
		{
			name:   "7290:6 8 6 15",
			expr:   []int{6, 8, 6, 15},
			target: 7290,
			want:   true,
		},
		{
			name:   "192:17 8 14",
			expr:   []int{17, 8, 14},
			target: 192,
			want:   true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := canBeMade(tt.expr, tt.target)
			if got != tt.want {
				t.Errorf("got %v want %v for equation %s", got, tt.want, tt.name)
			}
		})
	}
}
