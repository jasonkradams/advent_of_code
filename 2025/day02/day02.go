package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

type Answer struct {
	PartOne PartOne
	PartTwo PartTwo
}

type PartOne struct {
	invalidIDs []uint64
	Answer     uint64
}

type PartTwo struct {
	invalidIDs []uint64
	Answer     uint64
}

func Day02(idRanges []string) *Answer {
	answer := &Answer{}

	// Part 1
	for _, r := range idRanges {
		rangeStart, _ := strconv.ParseUint(strings.Split(r, "-")[0], 10, 64)
		rangeEnd, _ := strconv.ParseUint(strings.Split(r, "-")[1], 10, 64)
		// fmt.Printf("%d - %d\n", l, r)

		var currentID uint64
		for currentID = rangeStart; currentID <= rangeEnd; currentID++ {
			// fmt.Print("", i)
			idLower, idUpper := idRange(currentID)
			if idLower == idUpper {
				answer.PartOne.invalidIDs = append(answer.PartOne.invalidIDs, currentID)
				// fmt.Print("  !!!!! +1 !!!!! ")
				continue
			}

			// Part 2
			partTwo(currentID, answer)
		}
	}

	answer.PartOne.Answer = InvalidIDsTotal(answer.PartOne.invalidIDs)
	answer.PartTwo.Answer = InvalidIDsTotal(answer.PartTwo.invalidIDs) + answer.PartOne.Answer

	return answer
}

func partTwo(id uint64, answer *Answer) {
	idStr := idToString(id)
	left := 0
	for right := 0; right < len(idStr)/2; right++ {
		windowLength := right - left + 1
		sep := idStr[left : right+1]
		// fmt.Printf("id:%d,win:%d,sep:%q\n", id, windowLength, sep)

		n := (len(idStr) / windowLength)
		s := strings.SplitAfterN(idStr, sep, n)
		// fmt.Printf("id:%d,win:%d,sep:%q,%d,%#v\n", id, windowLength, sep, n, s)

		currentSlice := s[0]
		for _, slice := range s {
			// fmt.Println(id, "Hello", currentSlice, slice)
			if currentSlice != slice {
				goto skip
			}
			currentSlice = slice
		}

		// fmt.Printf("id:%d,slices: %v\n", id, s)
		answer.PartTwo.invalidIDs = append(answer.PartTwo.invalidIDs, id)

	skip:
	}
}

func InvalidIDsTotal(ids []uint64) uint64 {
	var total uint64
	for _, id := range ids {
		total += id
	}
	return total
}

func SolveDayTwo() *Answer {
	fileName := "input.txt"
	input := file.Get(fileName)

	idRanges := file.ReadLines(input)
	idRanges = strings.Split(idRanges[0], ",")
	// fmt.Printf("content: %#v\n", idRanges)

	answer := Day02(idRanges)
	fmt.Printf("Part1 Answer: %d\n", answer.PartOne.Answer)
	fmt.Printf("Part2 Answer: %d\n", answer.PartTwo.Answer)

	return answer
}

func idToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}

func idRange(id uint64) (lower, upper string) {
	idStr := idToString(id)
	lower = idStr[:len(idStr)/2]
	upper = idStr[len(idStr)/2:]

	return lower, upper
}
