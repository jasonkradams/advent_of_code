// This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words.
// It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them.
// Here are a few ways XMAS might appear, where irrelevant characters have been replaced with x:
// xxXxxx
// xSAMXx
// xAxxAx
// XMASxS
// xXxxxx

package main

import (
	"aoc/file"
	"bufio"
	"fmt"
	"os"
)

const filePath string = "input.txt"

func findXMAS(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	// Directions to search in
	directions := [][2]int{
		{0, 1}, {0, -1}, // Horizontal
		{1, 0}, {-1, 0}, // Vertical
		{1, 1}, {-1, -1}, // Diagonal
		{1, -1}, {-1, 1}, // Anti-diagonal
	}

	word := "XMAS"
	wordLen := len(word)
	matchCount := 0

	// Helper function to check if a word matches in a given direction
	checkDirection := func(x, y, dx, dy int) bool {
		for i := 0; i < wordLen; i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	// Scan the grid
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == 'X' { // Start with the first character
				for _, d := range directions {
					if checkDirection(x, y, d[0], d[1]) {
						matchCount++
					}
				}
			}
		}
	}
	return matchCount
}

func readGridFromFile(file *os.File) [][]rune {
	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid
}

func taskOne(grid [][]rune) int {
	matchCount := findXMAS(grid)
	return matchCount
}

func taskTwo() bool {
	return false
}

func main() {
	file := file.Get(filePath)
	defer file.Close()

	grid := readGridFromFile(file)

	taskOneResult := taskOne(grid)
	taskTwoResult := taskTwo()

	fmt.Println("task 1: ", taskOneResult)
	fmt.Println("task 2: ", taskTwoResult)
}
