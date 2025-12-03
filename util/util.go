package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(file string) ([]string, error) {
	stuff, err := os.Open(file)
	if err != nil {
		fmt.Print("Error opening file \n")
		return nil, err
	}

	defer stuff.Close()

	var result []string

	scanner := bufio.NewScanner(stuff)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

func Abs(x int) int {
	if x < 0 {
		x = x * -1
	}

	return x
}
