package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

//
// ------------------------- PART 1 LOGIC (original) -------------------------
//

// solve (Part 1) computes the number of unique antinode locations under
// the "1:2 distance" rule.
func solve(lines []string) int {
	// Collect antenna positions by frequency
	freqMap := make(map[rune][][2]int) // freq -> list of (row,col)

	// Dimensions of the grid
	rows := len(lines)
	if rows == 0 {
		return 0
	}
	cols := len(lines[0])

	// Parse the grid
	for r, line := range lines {
		for c, ch := range line {
			if ch != '.' {
				freqMap[ch] = append(freqMap[ch], [2]int{r, c})
			}
		}
	}

	// A set to collect all unique antinodes
	antinodes := make(map[[2]int]bool)

	// For each frequency, consider all pairs of antennas
	for _, coords := range freqMap {
		if len(coords) < 2 {
			continue
		}
		for i := 0; i < len(coords); i++ {
			A := coords[i]
			for j := i + 1; j < len(coords); j++ {
				B := coords[j]

				// P1 = 2A - B
				P1 := [2]int{2*A[0] - B[0], 2*A[1] - B[1]}
				// P2 = 2B - A
				P2 := [2]int{2*B[0] - A[0], 2*B[1] - A[1]}

				// Check if P1 is within bounds
				if P1[0] >= 0 && P1[0] < rows && P1[1] >= 0 && P1[1] < cols {
					antinodes[P1] = true
				}
				// Check if P2 is within bounds
				if P2[0] >= 0 && P2[0] < rows && P2[1] >= 0 && P2[1] < cols {
					antinodes[P2] = true
				}
			}
		}
	}

	return len(antinodes)
}

//
// ------------------------- PART 2 LOGIC (new model) -------------------------
//

// solvePart2 computes the number of antinode locations under the updated rule:
// “An antinode occurs at ANY grid position that is exactly in line with at least
// two antennas of the same frequency, regardless of distance.”
//
// Key steps:
//
//  1. Parse antennas by frequency, same as Part 1.
//
//  2. For each frequency, consider each distinct pair of antennas (A,B).
//
//  3. Find *all grid points collinear* with A,B (infinite line) that lie inside
//     the grid, and mark them as antinodes.
//
//     Note: Because we have a set of valid positions, the same point is only counted once,
//     even if it is collinear with multiple pairs.
func solvePart2(lines []string) int {
	freqMap := make(map[rune][][2]int)

	rows := len(lines)
	if rows == 0 {
		return 0
	}
	cols := len(lines[0])

	// Parse the grid
	for r, line := range lines {
		for c, ch := range line {
			if ch != '.' {
				freqMap[ch] = append(freqMap[ch], [2]int{r, c})
			}
		}
	}

	// A set of unique antinode positions
	antinodes := make(map[[2]int]bool)

	// For each frequency, consider all pairs of antennas
	for _, coords := range freqMap {
		// If there's only one antenna, it can't form a line with others
		if len(coords) < 2 {
			continue
		}

		for i := 0; i < len(coords); i++ {
			A := coords[i]
			for j := i + 1; j < len(coords); j++ {
				B := coords[j]

				dx := B[0] - A[0]
				dy := B[1] - A[1]

				// If dx == 0 and dy == 0, A and B are the same point
				// (which shouldn't happen if each grid square is unique),
				// but let's just guard anyway.
				if dx == 0 && dy == 0 {
					continue
				}

				// Reduce (dx, dy) to the smallest integer step (the slope).
				g := gcd(abs(dx), abs(dy))
				stepR := dx / g
				stepC := dy / g

				//
				// Walk from A forward in the slope direction
				// until we go out of bounds, adding each collinear point.
				//
				rCur, cCur := A[0], A[1]
				for {
					rCur += stepR
					cCur += stepC
					if rCur < 0 || rCur >= rows || cCur < 0 || cCur >= cols {
						break
					}
					antinodes[[2]int{rCur, cCur}] = true
				}

				//
				// Walk from A backward in the slope direction
				// until we go out of bounds.
				//
				rCur, cCur = A[0], A[1]
				for {
					rCur -= stepR
					cCur -= stepC
					if rCur < 0 || rCur >= rows || cCur < 0 || cCur >= cols {
						break
					}
					antinodes[[2]int{rCur, cCur}] = true
				}

				// Also mark A and B themselves as antinodes,
				// because each antenna in line with at least one other
				// (same freq) is an antinode.
				antinodes[A] = true
				antinodes[B] = true
			}
		}
	}

	return len(antinodes)
}

//
// ------------------------- HELPER FUNCTIONS -------------------------
//

// gcd computes the greatest common divisor of a and b (Euclid's algorithm).
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//
// ------------------------- MAIN -------------------------
//

// For TDD, we currently only output the Part 1 result.
// The updated solvePart2 is implemented above, but not invoked here.
func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	// PART 1 output
	resultPart1 := solve(lines)
	fmt.Println("Part 1 result:", resultPart1)

	// If you want to see PART 2 result, uncomment the following lines:
	resultPart2 := solvePart2(lines)
	fmt.Println("Part 2 result:", resultPart2)
}
