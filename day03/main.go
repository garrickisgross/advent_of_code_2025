package main

import (
	"fmt"
	"strconv"

	"github.com/garrickisgross/advent_of_code_2025/util"
)

func main() {
	part1()
	part2()
}

func maxSubsequence(line string, k int) string {
	n := len(line)
	result := make([]byte, 0, k)

	start := 0
	for pos := 0; pos < k; pos++ {
		remaining := k - pos
		end := n - remaining

		maxCh := byte('0')
		maxIdx := start

		for i := start; i <= end; i++ {
			if line[i] > maxCh {
				maxCh = line[i]
				maxIdx = i
			}
		}

		result = append(result, maxCh)
		start = maxIdx + 1
	}

	return string(result)
}

func processHighestJoltage(s string, k int) (int, error) {

	return strconv.Atoi(maxSubsequence(s, k))

}

func part1() {
	lines, _ := util.ReadInputLines("input.txt")

	sum := 0

	for _, v := range lines {
		val, _ := processHighestJoltage(v, 2)
		sum += val
	}

	fmt.Printf("Answer for Part 2: %d\n", sum)
}

func part2() {
	lines, _ := util.ReadInputLines("input.txt")

	sum := 0

	for _, v := range lines {
		val, _ := processHighestJoltage(v, 12)
		sum += val
	}

	fmt.Printf("Answer for Part 2: %d\n", sum)
}
