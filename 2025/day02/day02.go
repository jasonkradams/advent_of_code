package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

type Answer struct {
	A string
}

func Day02(idRanges []string) uint64 {
	invalidIDs := []uint64{}

	for _, r := range idRanges {
		l, _ := strconv.ParseUint(strings.Split(r, "-")[0], 10, 64)
		r, _ := strconv.ParseUint(strings.Split(r, "-")[1], 10, 64)
		// fmt.Printf("%d - %d\n", l, r)

		var i uint64
		for i = l; i <= r; i++ {
			fmt.Print("", i)
			iStr := strconv.FormatUint(i, 10)
			lStr := iStr[:len(iStr)/2]
			rStr := iStr[len(iStr)/2:]
			if lStr == rStr {
				invalidIDs = append(invalidIDs, i)
				// fmt.Print("  !!!!! +1 !!!!! ")
			}
			// fmt.Printf(" > %q:%q\n", lStr, rStr)
		}
	}
	return InvalidIDsTotal(invalidIDs)
}

func InvalidIDsTotal(ids []uint64) uint64 {
	var total uint64
	for _, id := range ids {
		total += id
	}
	return total
}

func PartOne() uint64 {
	fileName := "input.txt"
	input := file.Get(fileName)

	idRanges := file.ReadLines(input)
	idRanges = strings.Split(idRanges[0], ",")
	// fmt.Printf("content: %#v\n", idRanges)

	answer := Day02(idRanges)
	fmt.Printf("Answer: %v\n", answer)
	fmt.Printf("Answer: %d\n", answer)
	return answer
}
