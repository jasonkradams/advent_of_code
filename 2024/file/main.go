package file

import (
	"fmt"
	"os"
)

func Get(file string) *os.File {
	contents, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	return contents
}
