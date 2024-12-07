package main

import (
	"aoc/file"
	"fmt"
)

const (
	filePath string = "input.txt"
)

func taskOne() bool {
	return false
}
func taskTwo() bool {
	return false
}

func main() {
	file := file.Get(filePath)
	defer file.Close()

	taskOneResult := taskOne()
	taskTwoResult := taskTwo()

	fmt.Println("task 1: ", taskOneResult)
	fmt.Println("task 2: ", taskTwoResult)
}
