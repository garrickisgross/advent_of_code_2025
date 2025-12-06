package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/garrickisgross/advent_of_code_2025/util"
)

func main() {
	part1()
	part2()
}

func resolveOverlap(x []int, res *[][]int, j int) {
	(*res)[j][0] = int(math.Min(float64((x[0])), float64((*res)[j][0])))
	(*res)[j][1] = int(math.Max(float64((x[1])), float64((*res)[j][1])))
}

func mergeRanges(ranges [][]int) [][]int {
	var result [][]int
	edited := false
	min := 0
	max := 1

	result = append(result, ranges[0])
	ranges = ranges[1:]

	for i := range ranges {
		edited = false
		for j := range result {
			if ranges[i][max] < result[j][min] || ranges[i][min] > result[j][max] {
				continue
			} else {
				resolveOverlap(ranges[i], &result, j)
				edited = true
			}

		}

		if edited {
			result = mergeRanges(result)
		} else {
			result = append(result, ranges[i])
		}

	}

	return result
}

func checkRanges(n int, ranges [][]int) bool {
	for i := range ranges {
		if ranges[i][0] <= n && ranges[i][1] >= n {
			return true
		}
	}

	return false
}

func part1() {
	lines, _ := util.ReadInputLines("input.txt")
	result := 0
	var ranges [][]int
	var numbers []int
	for _, v := range lines {
		if strings.Contains(v, "-") {
			num_range := strings.Split(v, "-")
			var temp []int
			min, _ := strconv.Atoi(num_range[0])
			max, _ := strconv.Atoi(num_range[1])
			temp = append(temp, min, max)
			ranges = append(ranges, temp)
			continue
		}

		if v == "" {
			continue
		}
		num, _ := strconv.Atoi(v)
		numbers = append(numbers, num)
	}

	for _, v := range numbers {
		if checkRanges(v, ranges) {
			result += 1
		}
	}

	fmt.Printf("Answer Part 1: %d\n", result)
}

func part2() {
	lines, _ := util.ReadInputLines("input.txt")
	result := 0
	var ranges [][]int
	for _, v := range lines {
		if strings.Contains(v, "-") {
			num_range := strings.Split(v, "-")
			var temp []int
			min, _ := strconv.Atoi(num_range[0])
			max, _ := strconv.Atoi(num_range[1])
			temp = append(temp, min, max)
			ranges = append(ranges, temp)
			continue
		}

		if v == "" {
			continue
		}
	}
	ranges = mergeRanges(ranges)
	for i := range ranges {
		result += ranges[i][1] - ranges[i][0] + 1
	}

	fmt.Printf("Answer Part 2: %d\n", result)
}
