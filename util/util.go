package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInputLines(path string) ([]string, error) {
	content, err := os.Open(path)
	if err != nil {
		fmt.Print("Error opening file")
		return nil, err
	}

	defer content.Close()

	var result []string

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

func ReadInputDelimit(path string, delim string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error opening file")
		return nil, err
	}

	result := strings.Split(string(content), delim)

	return result, nil

}

func ReadInput2dArray(path string) [][]string {
	lines, _ := ReadInputLines(path)
	var result [][]string
	for _, v := range lines {
		split := strings.Split(v, "")
		result = append(result, split)
	}

	return result
}
