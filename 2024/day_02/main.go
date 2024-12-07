package main

import (
	"aoc/file"

	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	atLeast = 1
	atMost  = 3

	filePath string = "input.txt"
)

type Report struct {
	levels       []int
	incrementing bool
	decrementing bool
	safe         bool
}

func NewReport() Report {
	return Report{
		safe: true,
	}
}

func checkIncrementing(report *Report, reading, lastReading int) {
	if reading > lastReading {
		report.incrementing = true
	}
}

func checkDecrementing(report *Report, reading, lastReading int) {
	if reading < lastReading {
		report.decrementing = true
	}
}

// checkSafety marks a report safe if:
// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func checkSafety(report Report) bool {
	var lastReading int = report.levels[0]

	for i := 1; i < len(report.levels); i++ {

		reading := report.levels[i]

		checkIncrementing(&report, reading, lastReading)
		checkDecrementing(&report, reading, lastReading)

		if reading == lastReading {
			report.safe = false
			break
		}
		if report.incrementing && report.decrementing {
			report.safe = false
			break
		}
		if safeDifference(lastReading, reading) != true {
			report.safe = false
			break
		}

		lastReading = reading
	}

	return report.safe
}

func safeDifference(a, b int) bool {
	difference := int(math.Abs(float64(a - b)))

	if difference < atLeast {
		return false
	}

	if difference > atMost {
		return false
	}
	return true
}

func safeCount(reports []Report) int {
	// Count reports with Safe = true
	safeCount := 0
	for _, report := range reports {
		if report.safe {
			safeCount++
		}
	}
	return safeCount
}

func taskOne(reports []Report) int {
	// Identify and mark all safe reports.
	for i := range reports {
		reports[i].safe = checkSafety(reports[i])
	}

	safeCount := safeCount(reports)
	return safeCount
}

func deleteLevel(levels []int, i int) []int {
	// Create a copy of the slice to avoid mutating the original
	copyLevels := append([]int{}, levels...)
	return append(copyLevels[:i], copyLevels[i+1:]...)
}

func taskTwo(reports []Report) int {
	for i, report := range reports {
		if safe := checkSafety(report); safe {
			reports[i].safe = true
			break
		}

		// reset safety to true to give us another chance.
		report.safe = true

		// Try every combination of removing a single reading from the report
		// i.e. does any combination of full report, or full report -1 result in a pass?
		// logger.Printf("%d:\treport: %v", i, report.levels)
		for j := range report.levels {
			newReport := report
			// Always work on the original levels, not mutated ones
			newReport.levels = deleteLevel(newReport.levels, j)

			if safe := checkSafety(newReport); safe {
				reports[i].safe = true
				break
			}
		}
	}

	safeCount := safeCount(reports)
	return safeCount
}

func buildReports(file *os.File) []Report {
	var reports []Report
	// Read from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var readings []int
		for _, reading := range strings.Fields(line) {
			reading, err := strconv.Atoi(reading)
			if err != nil {
				fmt.Printf("Invalid line: %s\n", line)
			}

			readings = append(readings, reading)
		}

		report := NewReport()
		report.levels = readings

		reports = append(reports, report)
	}
	return reports
}

func main() {

	// Open the file
	file := file.Get(filePath)
	defer file.Close()
	reports := buildReports(file)

	taskOneResult := taskOne(reports)
	taskTwoResult := taskTwo(reports)

	fmt.Println("task 1: ", taskOneResult)
	fmt.Println("task 2: ", taskTwoResult)
}
