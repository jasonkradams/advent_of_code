package day01

import (
	"fmt"
	"strconv"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

const (
	zeroPoint   = 0
	startPoint  = 50
	maxPoint    = 99
	totalPoints = 100
)

type Answer struct {
	LandOnZero     int
	ClicksPastZero int
}

type Rotation struct {
	point    int
	distance int
}

func NewAnswer(rotations []string) *Answer {
	return &Answer{
		LandOnZero:     0,
		ClicksPastZero: 0,
	}
}

func (a *Answer) rotateLeft(r Rotation) int {
	a.ClicksPastZero += r.distance / totalPoints
	r.distance = r.distance % totalPoints

	point := r.point - r.distance

	if point < zeroPoint {
		if r.point != 0 {
			a.ClicksPastZero++
		}
		point = (maxPoint + 1 + point)
	}
	if point == 0 {
		a.ClicksPastZero++
	}
	return point
}

func (a *Answer) rotateRight(r Rotation) int {
	a.ClicksPastZero += r.distance / totalPoints
	distance := r.distance % totalPoints

	point := r.point + distance
	if point > maxPoint {
		if r.point != 0 {
			a.ClicksPastZero++
		}
		return (point - maxPoint - 1)
	}
	if point == 0 {
		a.ClicksPastZero++
	}
	return point
}

func Day01(rotations []string) Answer {
	answer := NewAnswer(rotations)
	currentPoint := startPoint
	for _, r := range rotations {
		direction := string(r[0])
		distanceS := r[1:]
		distance, _ := strconv.Atoi(distanceS)

		switch direction {
		case "L":
			b := currentPoint
			r := answer.ClicksPastZero
			currentPoint = answer.rotateLeft(Rotation{currentPoint, distance})
			fmt.Printf("%q: %3.1d < %3.1d < %3.1d (%2.1d)\n", direction, b, distance, currentPoint, (answer.ClicksPastZero - r))
		case "R":
			b := currentPoint
			r := answer.ClicksPastZero
			currentPoint = answer.rotateRight(Rotation{currentPoint, distance})
			fmt.Printf("%q: %3.1d > %3.1d > %3.1d (%2.1d)\n", direction, b, distance, currentPoint, (answer.ClicksPastZero - r))
		default:
			fmt.Printf("Unknown direction %q", direction)
		}

		if currentPoint == 0 {
			answer.LandOnZero++
		}
	}
	fmt.Printf("%#v\n", answer)

	return *answer
}

func PartOne() Answer {
	fileName := "input.txt"
	input := file.Get(fileName)
	rotations := file.ReadLines(input)
	answer := Day01(rotations)
	fmt.Printf("Answer: %#v\n", answer)
	fmt.Printf("Total: %d\n", answer.ClicksPastZero+answer.LandOnZero)
	return answer
}
