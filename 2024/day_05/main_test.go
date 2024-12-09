package main

import (
	"testing"
)

var pageRules = pageUpdates{
	rules: []orderingRule{
		{first: 47, last: 53},
		{first: 97, last: 13},
		{first: 97, last: 61},
		{first: 97, last: 47},
		{first: 75, last: 29},
		{first: 61, last: 13},
		{first: 75, last: 53},
		{first: 29, last: 13},
		{first: 97, last: 29},
		{first: 53, last: 29},
		{first: 61, last: 53},
		{first: 97, last: 53},
		{first: 61, last: 29},
		{first: 47, last: 13},
		{first: 75, last: 47},
		{first: 97, last: 75},
		{first: 47, last: 61},
		{first: 75, last: 61},
		{first: 47, last: 29},
		{first: 75, last: 13},
		{first: 53, last: 13},
	},
	pages: []pagesToProduce{
		{[]int{75, 47, 61, 53, 29}},
		{[]int{97, 61, 53, 29, 13}},
		{[]int{75, 29, 13}},
		{[]int{75, 97, 47, 61, 53}},
		{[]int{61, 13, 29}},
		{[]int{97, 13, 75, 29, 47}},
	},
}

var fileName string = "input.txt"

func TestTaskOne(t *testing.T) {
	t.Run("task 1", func(t *testing.T) {

		got := taskOne(pageRules)
		want := 143
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

// func TestTaskTwo(t *testing.T) {
// 	t.Run("task 2", func(t *testing.T) {
// 		got := taskTwo(pageRules)
// 		want := 9
// 		if got != want {
// 			t.Errorf("got %d want %d", got, want)
// 		}
// 	})
// }
