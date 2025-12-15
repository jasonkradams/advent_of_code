package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

type State struct {
	x, y, dir int
}

const (
	filePath string = "../input.txt"
)

/*
readGridFromFile reads the grid from the given file handle and returns it as a 2D byte slice.

- Parameters:
  - file *os.File: The open file handle to read the grid from.

- Returns:
  - [][]byte: A 2D array representing the grid layout.
*/
func readGridFromFile(file *os.File) [][]byte {
	grid := make([][]byte, 0)
	for line := bufio.NewScanner(file); line.Scan(); {
		gridLine := []byte(line.Text())
		grid = append(grid, gridLine)
	}
	return grid
}

/*
taskTwo simulates the guard's movement for each potential obstruction position to detect loops.

- Parameters:
  - grid [][]byte: The original grid layout.

- Returns:
  - int: The count of valid obstruction positions that cause the guard to loop indefinitely.
*/
func taskTwo(grid [][]byte) int {
	var x0, y0 int
	found := false
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '^' {
				x0, y0 = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	directions := []struct {
		dx int // delta x (row index)
		dy int // delta y (column index)
	}{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	count := 0

	for i := range grid {
		for j := range grid[i] {
			if i == x0 && j == y0 {
				continue // Skip the starting position
			}

			newGrid := copyGrid(grid, i, j)
			loopDetected := simulateGuardMovement(newGrid, x0, y0, directions)

			if loopDetected {
				count++
			}
		}
	}

	return count
}

/*
copyGrid creates a deep copy of the original grid and adds an obstruction at the specified position.

- Parameters:
  - original [][]byte: The original grid layout.
  - i int: Row index where the obstruction is to be added.
  - j int: Column index where the obstruction is to be added.

- Returns:
  - [][]byte: A new grid with the obstruction added.
*/
func copyGrid(original [][]byte, i, j int) [][]byte {
	newGrid := make([][]byte, len(original))
	for row := 0; row < len(original); row++ {
		newRow := make([]byte, len(original[row]))
		copy(newRow, original[row])
		newGrid[row] = newRow
	}
	if newGrid[i][j] != '#' {
		newGrid[i][j] = '#'
	}
	return newGrid
}

/*
simulateGuardMovement simulates the guard's movement on a given grid starting from a specific position and detects loops.

- Parameters:
  - grid [][]byte: The grid layout with an obstruction added.
  - x0 int: Initial row position of the guard.
  - y0 int: Initial column position of the guard.
  - directions []struct{dx, dy}: Possible movement directions for the guard.

- Returns:
  - bool: True if a loop is detected, False otherwise.
*/
func simulateGuardMovement(grid [][]byte, x0, y0 int, directions []struct{ dx, dy int }) bool {
	visited := make(map[State]bool)
	current_i, current_j := x0, y0
	dir := 0 // Start facing up

	for {
		stateKey := State{current_i, current_j, dir}
		if visited[stateKey] {
			return true // Loop detected
		}
		visited[stateKey] = true

		dx := directions[dir].dx
		dy := directions[dir].dy

		front_i := current_i + dx
		front_j := current_j + dy

		rows := len(grid)
		cols := len(grid[0])
		front_in_bounds := front_i >= 0 && front_i < rows && front_j >= 0 && front_j < cols

		if !front_in_bounds {
			current_i += dx
			current_j += dy

			if current_i < 0 || current_i >= rows || current_j < 0 || current_j >= cols {
				break // Guard leaves the area
			}
		} else {
			if grid[front_i][front_j] == '#' {
				dir = (dir + 1) % 4 // Blocked, turn right
			} else {
				current_i += dx
				current_j += dy

				if current_i < 0 || current_i >= rows || current_j < 0 || current_j >= cols {
					break // Guard leaves the area
				}
			}
		}
	}

	return false // No loop detected
}

/*
main is the entry point of the program. It reads the input file, identifies the guard's starting position, and delegates simulation to taskTwo.

- Steps:
 1. Read the grid from the input file.
 2. Identify the guard's starting position (^).
 3. Delegate loop detection to taskTwo(grid) for each potential obstruction.
 4. Print the result.

- Output:
  - Prints the count of valid obstruction positions that cause the guard to loop indefinitely.
*/
func main() {
	file := file.Get(filePath)
	defer file.Close()

	grid := readGridFromFile(file)

	taskTwoResult := taskTwo(grid)

	fmt.Printf("task2: %d\n", taskTwoResult)
}
