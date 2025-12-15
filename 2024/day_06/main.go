package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

type Direction int

const (
	filePath string = "input.txt"
)

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

	// Set of visited cells
	visited := make(map[[2]int]bool)
	visited[[2]int{startX, startY}] = true

	// Function to check if a cell is within bounds and not an obstruction
	isValid := func(x, y int) bool {
		return !(x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '#')
	}

	x, y := startX, startY
	for {
		// Move until obstruction
		for {
			nextX, nextY := x+directions[dir][0], y+directions[dir][1]
			if !isValid(nextX, nextY) {
				break
			}
			x += directions[dir][0]
			y += directions[dir][1]
			visited[[2]int{x, y}] = true
		}

		// Try turning right (update direction)
		dir = (dir + 1) % 4

		// Check if we can move in the new direction
		nextX, nextY := x+directions[dir][0], y+directions[dir][1]
		if !isValid(nextX, nextY) {
			break // No valid moves; exit the loop
		}
	}

	return len(visited)
}

const (
	Right Direction = iota
	Down
	Left
	Up
)

func newDirection(d Direction) Direction {
	if d == Right {
		return Down
	}
	if d == Down {
		return Left
	}
	if d == Left {
		return Up
	}
	return Right
}

var directionStrings = map[Direction]string{
	Right: "Right",
	Down:  "Down",
	Left:  "Left",
	Up:    "Up",
}

func countInfiniteLoops(grid [][]rune) int {
	infiniteLoopCount := 0

	// Function to check if a cell is within bounds and not an obstruction
	isValid := func(x, y int, grid [][]rune) bool {
		return (x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '#')
	}

	// Function to check for infinite loop by detecting repeated states
	containsInfiniteLoop := func(startX, startY int, grid [][]rune) bool {
		for _, d := range []Direction{Right, Down, Left, Up} {
			var movements int

			visitedStates := make(map[[3]any]bool)
			x, y, dir := startX, startY, d

			for {
				movements += 1

				state := [3]any{x, y, dir}
				fmt.Printf("%v, ", state)
				fmt.Printf("(%d, %d) visiting - %s", x, y, directionStrings[dir])
				if visitedStates[state] {
					fmt.Printf("\nvisitedState%v", state)
					return true
				}
				visitedStates[state] = true

				// Move until obstruction
				for {
					nextX, nextY := x+directions[dir][0], y+directions[dir][1]
					if !isValid(nextX, nextY, grid) {
						fmt.Printf(" - is invalid\n")
						break
					}
					x, y = nextX, nextY
				}

				// Turn right
				dir = (dir + 1) % 4

				// Check if next move is valid
				nextX, nextY := x+directions[dir][0], y+directions[dir][1]
				if !isValid(nextX, nextY, grid) {
					break
				}
			}
		}
		return false
	}

	// Iterate over the grid to try placing an obstruction at every valid cell
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' { // Place obstruction only on empty cells
				grid[i][j] = '#'
				fmt.Printf("\n(%d, %d)\n", i, j)
				if containsInfiniteLoop(i, j, grid) {
					fmt.Printf(" is infinite\n")
					infiniteLoopCount++
				} else {
					fmt.Printf("")
				}
				grid[i][j] = '.' // Restore cell
			}
		}
	}

	return infiniteLoopCount
}

func taskOne(grid [][]rune) int {
	return findRoute(grid)
}

func taskTwo(grid [][]rune) int {
	return countInfiniteLoops(grid)
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
