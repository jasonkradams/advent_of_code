package day01

import (
	"fmt"
	"strconv"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

const (
	zeroPoint  = 0
	startPoint = 50
	maxPoint   = 99
)

func rotateLeft(point, distance int) int {
	distance = distance % 100
	point -= distance
	if point < zeroPoint {
		point = (maxPoint + 1 + point)
	}
	return point
}

func rotateRight(point, distance int) int {
	distance = distance % 100
	point += distance
	if point > maxPoint {
		return (point - maxPoint - 1)
	}
	return point

}

func Day01(rotations []string) int {
	currentPoint := startPoint
	var code int
	for _, r := range rotations {
		direction := string(r[0])
		distanceS := r[1:]
		distance, _ := strconv.Atoi(distanceS)

		// fmt.Println(direction, distance)

		switch direction {
		case "L":
			fmt.Printf("%qB, %d, distance: %d\n", direction, currentPoint, distance)

			currentPoint = rotateLeft(currentPoint, distance)

			fmt.Printf("%qA, %d, distance: %d\n", direction, currentPoint, distance)
		case "R":
			fmt.Printf("%qB, %d, distance: %d\n", direction, currentPoint, distance)
			currentPoint = rotateRight(currentPoint, distance)
			fmt.Printf("%qA, %d, distance: %d\n", direction, currentPoint, distance)
		default:
			fmt.Printf("Unknown direction %q", direction)
		}

		if currentPoint == 0 {
			code++
		}
	}

	return code
}

func PartOne() int {
	fileName := "input.txt"
	input := file.Get(fileName)
	rotations := file.ReadLines(input)
	answer := Day01(rotations)
	fmt.Println("Answer:", answer)
	return answer
}
