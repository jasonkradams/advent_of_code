package main

import (
	"aoc/file"
	"bufio"
	"fmt"
	"os"
)

const (
	filePath string = "input.txt"
)

// Directions: right, down, left, up
var directions = [][2]int{
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
	{-1, 0}, // Up
}

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

func findRoute(grid [][]rune) int {
	// Find starting position and direction
	var startX, startY, dir int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' {
				startX, startY = i, j
				dir = 3 // Initial direction is "up"
				break
			}
		}
	}
	fmt.Printf("Starting position: (%d, %d), Initial direction: %d\n", startX, startY, dir)

	// Set of visited cells
	visited := make(map[[2]int]bool)
	visited[[2]int{startX, startY}] = true
	fmt.Printf("Visited cells: %v\n", visited)

	// Function to check if a cell is within bounds and not an obstruction
	isValid := func(x, y int) (bool, string) {
		reason := ""
		if x < 0 || x >= len(grid) {
			reason += "out of bounds vertically; "
		}
		if y < 0 || y >= len(grid[0]) {
			reason += "out of bounds horizontally; "
		}
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == '#' {
			reason += "obstructed by a wall; "
		}
		valid := reason == ""
		return valid, reason
	}

	x, y := startX, startY
	for {
		// Move until obstruction
		var reason string
		for {
			var valid bool
			nextX, nextY := x+directions[dir][0], y+directions[dir][1]
			valid, reason = isValid(nextX, nextY)
			if !valid {
				fmt.Printf("Position (%d, %d) is not valid: %s\n", nextX, nextY, reason)
				break
			}

			x += directions[dir][0]
			y += directions[dir][1]
			visited[[2]int{x, y}] = true
			fmt.Printf("Moved to position: (%d, %d)\n", x, y)
		}

		if reason == "out of bounds horizontally; " || reason == "out of bounds vertically; " {
			break
		}

		// Try turning right (update direction)
		dir = (dir + 1) % 4
		fmt.Printf("Turned right, new direction: %d\n", dir)

		// Check if we can move in the new direction
		nextX, nextY := x+directions[dir][0], y+directions[dir][1]
		if valid, reason := isValid(nextX, nextY); !valid {
			fmt.Printf("No valid moves from position (%d, %d) in direction %d: %s\n", x, y, dir, reason)
			break // No valid moves; exit the loop
		}
	}

	fmt.Printf("Total unique visited cells: %d\n", len(visited))
	return len(visited)
}

func taskOne(grid [][]rune) int {
	return findRoute(grid)
}

func taskTwo(grid [][]rune) int {

	return 0
}

func main() {
	file := file.Get(filePath)
	defer file.Close()

	grid := readGridFromFile(file)

	taskOneResult := taskOne(grid)
	taskTwoResult := taskTwo(grid)

	fmt.Println("task 1:", taskOneResult)
	fmt.Println("task 2:", taskTwoResult)
}
