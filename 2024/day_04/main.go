package main

import (
	"aoc/file"
	"bufio"
	"fmt"
	"os"
)

const filePath string = "input.txt"

// Read the grid from the file
func readGridFromFile(file *os.File) [][]rune {
	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid
}

// Task 1: Count all "XMAS" occurrences
func findXMAS(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])

	directions := [][2]int{
		{0, 1}, {0, -1}, // Horizontal
		{1, 0}, {-1, 0}, // Vertical
		{1, 1}, {-1, -1}, // Diagonal
		{1, -1}, {-1, 1}, // Anti-diagonal
	}

	word := "XMAS"
	wordLen := len(word)
	matchCount := 0

	checkDirection := func(x, y, dx, dy int) bool {
		for i := 0; i < wordLen; i++ {
			nx, ny := x+dx*i, y+dy*i
			if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == 'X' {
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

// Task 2: Count all "X-MAS" patterns
func findXMASPatterns(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	matchCount := 0

	// Check for "MAS" in the shape of an X
	isMAS := func(x, y, dx, dy int) bool {
		midX, midY := x+dx, y+dy
		endX, endY := x+2*dx, y+2*dy

		if x < 0 || y < 0 || midX < 0 || midY < 0 || endX < 0 || endY < 0 ||
			x >= rows || y >= cols || midX >= rows || midY >= cols || endX >= rows || endY >= cols {
			return false
		}

		return (grid[x][y] == 'M' && grid[midX][midY] == 'A' && grid[endX][endY] == 'S') ||
			(grid[x][y] == 'S' && grid[midX][midY] == 'A' && grid[endX][endY] == 'M')
	}

	// Scan for X-MAS patterns
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			// Check if the center of the "X" is valid
			if grid[x][y] == 'A' {
				if isMAS(x-1, y-1, 1, 1) && isMAS(x-1, y+1, 1, -1) { // Top-left to bottom-right and top-right to bottom-left
					matchCount++
				}
			}
		}
	}

	return matchCount
}

// Task 1: Call findXMAS
func taskOne(grid [][]rune) int {
	return findXMAS(grid)
}

// Task 2: Call findXMASPatterns
func taskTwo(grid [][]rune) int {
	return findXMASPatterns(grid)
}

// Main function
func main() {
	file := file.Get(filePath)
	defer file.Close()

	grid := readGridFromFile(file)

	taskOneResult := taskOne(grid)
	taskTwoResult := taskTwo(grid)

	fmt.Println("task 1:", taskOneResult)
	fmt.Println("task 2:", taskTwoResult)
}
