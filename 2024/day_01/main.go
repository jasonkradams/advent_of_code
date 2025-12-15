package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	file "github.com/jasonkradams/advent_of_code/2024/file"
)

var (
	sum int
)

const (
	filePath string = "input.txt"
)

type Left struct {
	number    int
	peerCount int
}

type Right struct {
	number int
}

func main() {
	// Specify the file path (update this path as needed)

	// Open the file
	file := file.Get(filePath)
	defer file.Close()

	// Initialize a slice to hold the pairs
	// var keys []int
	var right []int
	var left []Left

	// Read from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			first, err1 := strconv.Atoi(parts[0])
			second, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				left = append(left, Left{first, 0})
				right = append(right, second)
			} else {
				fmt.Printf("Skipping invalid line: %s\n", line)
			}
		} else {
			fmt.Printf("Skipping malformed line: %s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Sort keys and values
	sort.Slice(left, func(i, j int) bool {
		return left[i].number < left[j].number
	})
	sort.Ints(right)

	taskOne := taskOne(left, right)
	fmt.Println("task 1:", taskOne)

	// task 2
	taskTwo := taskTwo(&left, right)
	fmt.Println("task 2:", taskTwo)

}

func getPeerCount(left *[]Left, right []int) {
	for i, target := range *left {
		count := 0
		for _, num := range right {
			if num == target.number {
				count++
			}
			(*left)[i].peerCount = count
		}
	}
}

func taskOne(left []Left, right []int) int {
	// Print the sorted pairs
	for i := range len(right) {
		first := left[i].number
		second := right[i]
		// fmt.Printf("%d %d\n", first, second)
		if first > second {
			sum += first - second
		} else {
			sum += second - first
		}
	}
	return sum
}

// getSimilarityScore calculates a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.
func getSimilarityScore(left *[]Left) int {
	sum := 0
	for _, num := range *left {
		if num.peerCount == 0 {
			continue
		}
		sum += num.number * num.peerCount
	}
	return sum
}

func taskTwo(left *[]Left, right []int) int {
	getPeerCount(left, right)
	similarityScore := getSimilarityScore(left)

	return similarityScore
}
