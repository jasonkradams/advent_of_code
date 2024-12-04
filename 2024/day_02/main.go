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

// isSafe marks a report safe if:
// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
func isSafe(report Report) bool {
	var lastReading int = report.levels[0]

	for i := 1; i < len(report.levels); i++ {
		reading := report.levels[i]

		checkIncrementing(&report, reading, lastReading)
		checkDecrementing(&report, reading, lastReading)

		if reading == lastReading {
			return false
		}

		if report.incrementing && report.decrementing {
			return false
		}
		if safeDifference(lastReading, reading) != true {
			return false
		}
		lastReading = reading
	}
	return true
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

func taskOne(reports []Report) int {
	// Identify and mark all safe reports.
	for i := range reports {
		reports[i].safe = isSafe(reports[i])
	}

	// Count reports with Safe = true
	safeCount := 0
	for _, report := range reports {
		if report.safe {
			safeCount++
		}
	}
	return safeCount
}

func taskTwo() {

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
		reports = append(reports, Report{levels: readings})
	}
	return reports
}

func main() {
	// Open the file
	file := file.Get(filePath)
	defer file.Close()

	reports := buildReports(file)
	safeCount := taskOne(reports)

	fmt.Println("task 1: ", safeCount)
}
