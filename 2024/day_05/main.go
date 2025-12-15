package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jasonkradams/advent_of_code/2024/file"
)

const (
	filePath               string = "input.txt"
	pageRuleRegexPattern   string = `(\d+)\|(\d+)`
	pageUpdateRegexPattern string = `(\d+,?\d+)+`
)

var (
	pageRuleRegex   = regexp.MustCompile(pageRuleRegexPattern)
	pageUpdateRegex = regexp.MustCompile(pageUpdateRegexPattern)
)

type pageUpdates struct {
	rules []orderingRule
	pages []pagesToProduce
}

type orderingRule struct {
	first int
	last  int
}

type pagesToProduce struct {
	pages   []int
	ordered bool
}

// convertStringsToInts converts a slice of strings to a slice of integers.
// If a string cannot be converted to an integer, it is skipped.
//
// Parameters:
//   - input: A slice of strings to be converted.
//
// Returns:
//   - A slice of integers obtained from the input strings.
func convertStringsToInts(input []string) []int {
	var result []int
	for _, str := range input {
		num, _ := strconv.Atoi(str)
		result = append(result, num)
	}
	return result
}

// readPageRulesFromFile reads page rules and updates from the provided file.
// It parses lines matching specific regular expressions and structures them into `pageUpdates`.
//
// Parameters:
//   - file: A pointer to the file to read from.
//
// Returns:
//   - A `pageUpdates` struct containing parsed rules and updates.
func readPageRulesFromFile(file *os.File) pageUpdates {
	var updates pageUpdates

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if pageRuleRegex.MatchString(line) {
			rule := pageRuleRegex.FindStringSubmatch(line)

			first, _ := strconv.Atoi(rule[1])
			last, _ := strconv.Atoi(rule[2])

			updates.rules = append(updates.rules, orderingRule{first, last})
			continue
		}

		if pageUpdateRegex.MatchString(line) {
			pagesString := strings.Split(line, ",")
			pagesInt := convertStringsToInts(pagesString)

			updates.pages = append(updates.pages, pagesToProduce{pages: pagesInt})
		}

	}
	return updates
}

// getIndex returns the index of the first occurrence of the target value in the slice.
// If the target is not found, it returns -1.
//
// Parameters:
//   - slice: A slice of integers to search through.
//   - target: The integer value to search for in the slice.
//
// Returns:
//   - The index of the first occurrence of the target value in the slice, or -1 if the value is not found.
//
// Example:
//
//	slice := []int{18, 43, 32, 19}
//	target := 32
//	index := firstIndex(slice, target)
//	// index will be 2
func getIndex(target int, slice []int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// getMiddle returns the middle value of a slice of integers.
//
// Parameters:
//   - slice: A slice of integers.
//
// Returns:
//   - The middle value of the slice.
func getMiddle(slice []int) int {
	return slice[len(slice)/2]
}

// validatePageRules checks if the provided page updates follow the ordering rules.
// It ensures that the index of the `first` value in each rule is less than the index of the `last` value.
//
// Parameters:
//   - rules: A slice of orderingRule structs defining the rules to validate against.
//   - pageUpdates: A slice of integers representing the pages to validate.
//
// Returns:
//   - A boolean indicating whether all rules are valid for the given page updates.
func validatePageRules(rules []orderingRule, pageUpdates []int) bool {
	for _, rule := range rules {
		firstIndex := getIndex(rule.first, pageUpdates)
		lastIndex := getIndex(rule.last, pageUpdates)

		if firstIndex == -1 || lastIndex == -1 {
			continue
		}

		if firstIndex > lastIndex {
			return false
		}
	}
	return true
}

// taskOne calculates the sum of the middle pages of valid page updates according to the rules.
//
// Parameters:
//   - updates: A pageUpdates struct containing rules and pages.
//
// Returns:
//   - The sum of the middle pages of valid updates.
func taskOne(updates pageUpdates) int {
	var sum int
	for _, pages := range updates.pages {
		validRule := validatePageRules(updates.rules, pages.pages)
		if validRule {
			sum += getMiddle(pages.pages)
		}
	}
	return sum
}

// orderPages rearranges the pages slice to satisfy the given ordering rules.
// If a rule specifies that a page `first` should come before a page `last`,
// the function swaps elements in the slice until all rules are satisfied.
//
// Parameters:
//   - rules: A slice of orderingRule structs defining the order constraints.
//   - pages: A slice of integers representing the pages to be ordered.
//
// Returns:
//   - A reordered slice of integers that satisfies all the provided rules.
//
// Example:
//
//	rules := []orderingRule{{first: 1, last: 3}, {first: 2, last: 4}}
//	pages := []int{3, 1, 4, 2}
//	result := orderPages(rules, pages)
//	// result will be [1, 2, 3, 4]
func orderPages(rules []orderingRule, pages []int) []int {
	// Keep track of changes to prevent an endless loop
	for validatePageRules(rules, pages) != true {
		modified := false // Flag to check if we made a change

		for _, rule := range rules {
			first := getIndex(rule.first, pages)
			last := getIndex(rule.last, pages)

			if first == -1 || last == -1 {
				// If either page is not found, skip this rule
				continue
			}

			if first > last {
				// Swap to correct the order
				pages[first], pages[last] = pages[last], pages[first]
				modified = true
			}
		}

		if !modified {
			// If no modifications were made, break to avoid infinite loop
			break
		}
	}
	return pages
}

// taskTwo calculates the sum of the middle pages of updates that initially do not
// satisfy the ordering rules. Pages are reordered using orderPages before calculating the sum.
//
// Parameters:
//   - updates: A pageUpdates struct containing rules and pages.
//
// Returns:
//   - An integer representing the sum of the middle pages of updates that
//     were corrected to satisfy the ordering rules.
//
// Example:
//
//	updates := pageUpdates{
//	    rules: []orderingRule{{first: 1, last: 3}, {first: 2, last: 5}},
//	    pages: []pagesToProduce{
//	        {pages: []int{3, 1, 2, 5, 4}},
//	        {pages: []int{1, 3, 2, 4, 5}},
//	    },
//	}
//	result := taskTwo(updates)
//	// result will be the sum of the middle values of the corrected pages.
func taskTwo(updates pageUpdates) int {
	var sum int
	for _, pages := range updates.pages {
		validRule := validatePageRules(updates.rules, pages.pages)
		if !validRule {
			orderedPages := orderPages(updates.rules, pages.pages)
			sum += getMiddle(orderedPages)
		}
	}
	return sum
}

func main() {
	file := file.Get(filePath)
	defer file.Close()

	updates := readPageRulesFromFile(file)

	taskOneResult := taskOne(updates)
	taskTwoResult := taskTwo(updates)

	fmt.Println("task 1:", taskOneResult)
	fmt.Println("task 2:", taskTwoResult)
}
