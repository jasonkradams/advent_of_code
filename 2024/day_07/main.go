package main

import (
	"aoc/file"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	filePath string = "input.txt"
)

type Calibration struct {
	result     int
	expression []int
}

func stringToCalibration(str string) (Calibration, error) {
	parts := strings.Split(str, ": ")
	if len(parts) != 2 {
		return Calibration{}, fmt.Errorf("invalid input string format")
	}

	result, err := strconv.Atoi(parts[0])
	if err != nil {
		return Calibration{}, fmt.Errorf("error converting result to integer: %w", err)
	}

	expressionStrs := strings.Fields(parts[1])
	expression := make([]int, len(expressionStrs))
	for i, str := range expressionStrs {
		num, err := strconv.Atoi(str)
		if err != nil {
			return Calibration{}, fmt.Errorf("error converting expression element to integer: %w", err)
		}
		expression[i] = num
	}

	return Calibration{result, expression}, nil
}

func readCalibrationsFromFile(file *os.File) ([]Calibration, error) {
	var calibrations []Calibration
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		calibration, err := stringToCalibration(line)
		if err != nil {
			return []Calibration{}, fmt.Errorf("error converting line to calibration")
		}
		calibrations = append(calibrations, calibration)
	}
	return calibrations, nil
}

func generateExpressions(calibration Calibration) int {
	nums := calibration.expression

	for i := 0; i < (1 << len(nums)); i++ {
		expression := ""
		currentNumber := nums[0]

		// Iterate through each bit in the binary representation of i
		for j := 1; j < len(nums); j++ {
			bit := (i >> (j - 1)) & 1
			if bit == 1 {
				expression += fmt.Sprintf("*%d", nums[j])
				currentNumber *= nums[j]

			} else {
				expression += fmt.Sprintf("+%d", nums[j])
				currentNumber += nums[j]

			}
		}
		// this is a possible match and will be added to the sum
		if currentNumber == calibration.result {
			return currentNumber
		}
	}

	// no possible matches found
	return 0
}

func taskOne(calibrations []Calibration) int {
	sum := 0

	for _, calibration := range calibrations {
		sum += generateExpressions(calibration)
	}

	return sum
}

func taskTwo(calibrations []Calibration) int {
	return 0
}

func main() {
	file := file.Get(filePath)
	defer file.Close()

	calibrations, err := readCalibrationsFromFile(file)
	if err != nil {
		panic("unable to convert lines to calibrations")
	}

	taskOneResult := taskOne(calibrations)
	taskTwoResult := taskTwo(calibrations)

	fmt.Println("task 1:", taskOneResult)
	fmt.Println("task 2:", taskTwoResult)
}
