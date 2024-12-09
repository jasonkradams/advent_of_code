package file

import (
	"fmt"
	"os"
)

// Get opens the specified file and returns a pointer to the opened *os.File.
// If the file cannot be opened, it logs the error message and panics.
//
// Parameters:
//   - file: A string representing the path to the file to be opened.
//
// Returns:
//   - A pointer to the opened *os.File.
//
// Panics:
//   - If the file cannot be opened, the function logs the error message and terminates the program.
//
// Example:
//
//	file := file.Get("input.txt")
//	defer file.Close()
//
//	// Use the file for reading...
func Get(file string) *os.File {
	contents, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	return contents
}
