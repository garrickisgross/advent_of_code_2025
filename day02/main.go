package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/garrickisgross/advent_of_code_2025/util"
)

func main() {
	part1()
	part2()
}

func processRange(in string, fn func(int) bool) []int {
	temp := strings.Split(in, "-")
	min, _ := strconv.Atoi(temp[0])
	max, _ := strconv.Atoi(temp[1])

	var result []int

	for i := min; i <= max; i++ {
		if fn(i) {
			result = append(result, i)
		}
	}

	return result

}

func repeatingSeq(i int) bool {
	//will probably be easier as a str
	str := strconv.Itoa(i)

	return str[len(str)/2:] == str[:len(str)/2]

}

func repeatAnyAtLeastTwice(i int) bool {
	str := strconv.Itoa(i)

	seq := ""
	for _, v := range str {
		seq = seq + string(v)

		if str == strings.Repeat(seq, len(str)/len(seq)) && seq != str {
			return true
		}

	}

	return false
}

func part1() {
	id_ranges, err := util.ReadInputDelimit("input.txt", ",")
	if err != nil {
		fmt.Printf("Error Reading input file: %s", err)
		return
	}

	var result []int

	acc := 0

	for i := range id_ranges {
		result = processRange(id_ranges[i], repeatingSeq)
		for _, val := range result {
			acc += val
		}
	}

	fmt.Printf("Answer for Part 1: %d\n", acc)

}

func part2() {
	id_ranges, err := util.ReadInputDelimit("input.txt", ",")
	if err != nil {
		fmt.Printf("Error Reading input file: %s", err)
		return
	}

	var result []int

	acc := 0

	for i := range id_ranges {
		result = processRange(id_ranges[i], repeatAnyAtLeastTwice)
		for _, val := range result {
			acc += val
		}
	}

	fmt.Printf("Answer for Part 2: %d\n", acc)

}
