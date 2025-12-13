package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sb strings.Builder

var IDs map[int]int

/*
1. Convert from dense number format to expanded format.
2. Move last file block into first available free space (.)
*/

func solveTask1(m string) int {
	IDs = make(map[int]int)
	// fmt.Println("initial map: " + m)

	var fileID = 0
	IDs[fileID] = 0
	// create new map

	// initial conversion
	var lastType string
	for i, c := range m {
		num, _ := strconv.Atoi(string(c))

		// fmt.Printf("c: %c\n", c)
		// A file is even. Freespace will be odd
		if i%2 == 0 {
			if lastType == "freeSpace" {
				IDs[fileID] = 0
				lastType = "file"
			}

			// File is marked with an index
			for range num {
				IDs[fileID]++
				// fmt.Printf("In loop, fileindex: %d\n", fileIndex)
				sb.WriteString(strconv.Itoa(fileID))
			}
			fileID++
		} else {
			if lastType == "file" {
				lastType = "freeSpace"
			}

			for range num {
				sb.WriteString(".")
			}
		}
	}

	// move each file block from the end of the drive to the first available space
	result := sb.String()
	for {
		oldSub := strconv.Itoa(fileID)
		newSub := "."
		for range IDs[fileID] {

			replaceIndex := strings.LastIndex(result, oldSub)
			if replaceIndex != -1 {
				// Rebuild the string: everything up to replaceIndex, then the replacement,
				// then everything after replaceIndex+len(oldSub).
				result = result[:replaceIndex] + newSub + result[replaceIndex+len(oldSub):]
				result = strings.Replace(result, newSub, oldSub, 1)
			}
		}
		// fmt.Println(result)

		fileID--
		if fileID < 0 {
			break
		}
	}
	fmt.Println(result)

	return 0
}

func solveTask2(m string) int {
	diskMap := m
	fmt.Println("initial map: " + diskMap)
	return 0
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Could not read input file: %v", err)
	}

	diskMap := strings.TrimSpace(string(data))

	// PART 1 output
	resultPart1 := solveTask1(diskMap)
	fmt.Println("Part 1 result:", resultPart1)

	// sb.Reset()
	// // PART 2 output
	// resultPart2 := solveTask2(diskMap)
	// fmt.Println("Part 2 result:", resultPart2)
}
