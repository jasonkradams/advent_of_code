package main

import (
	"aoc/file"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	filePath        string = "input.txt"
	mulRegexPattern string = `(mul|do|don't)\((\d{1,3})?,?(\d{1,3})?\)`
)

var (
	mulRegex = regexp.MustCompile(mulRegexPattern)
)

type expression struct {
	operand string
	x       int
	y       int
}

func newMultiplier(x, y int) expression {
	return expression{x: x, y: y}
}

func buildExpressions(file *os.File) []expression {
	var matches []expression

	// Read from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Actions per line below
		lineMatch := mulRegex.FindAllStringSubmatch(line, -1)

		for _, match := range lineMatch {
			operand := match[1]
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])
			matches = append(matches, expression{operand: operand, x: x, y: y})
		}
	}

	return matches
}

func multiply(x, y int) int {
	return x * y
}

func taskOne(expressions []expression) int {
	var sum int

	for _, expression := range expressions {
		sum += multiply(expression.x, expression.y)
	}

	return sum

}

// taskTwo addresses https://adventofcode.com/2024/day/3#part2
// There are two new instructions you'll need to handle:
// The do() instruction enables future mul instructions.
// The don't() instruction disables future mul instructions.
// Only the most recent do() or don't() instruction applies. At the beginning of the program, mul instructions are enabled.
func taskTwo(expressions []expression) int {
	var sum int
	do := true

	for _, expression := range expressions {
		if expression.operand == "do" {
			do = true
		}
		if expression.operand == "don't" {
			do = false
		}

		if do {
			sum += multiply(expression.x, expression.y)
		}
	}
	return sum
}

func main() {
	file := file.Get(filePath)
	defer file.Close()
	expressions := buildExpressions(file)

	taskOneResult := taskOne(expressions)
	taskTwoResult := taskTwo(expressions)

	fmt.Println("task 1: ", taskOneResult)
	fmt.Println("task 2: ", taskTwoResult)
}
