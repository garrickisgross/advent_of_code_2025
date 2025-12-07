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

func doOperation(nums []int, op string) int {
	result := 0
	if op == "+" {
		for _, v := range nums {
			result += v
		}
	}

	if op == "*" {
		result = 1
		for _, v := range nums {
			result *= v
		}
	}

	return result

}

func doMath(arr [][]int, ops []string) int {
	result := 0

	for x := range ops {
		var nums []int
		for y := range arr {
			nums = append(nums, arr[y][x])
		}
		result += doOperation(nums, ops[x])
	}

	return result
}

func doMath2(nums [][]int, ops []string) int {
	result := 0

	for x := range ops {
		result += doOperation(nums[x], ops[x])
	}

	return result
}

func part1() {
	mathStuff := util.ReadInput2dArray("input.txt", " ")
	last := mathStuff[len(mathStuff)-1]
	mathStuff = mathStuff[:len(mathStuff)-1]
	var nums [][]int
	for _, v := range mathStuff {
		var temp []int
		for _, k := range v {
			num, _ := strconv.Atoi(k)
			if num != 0 {
				temp = append(temp, num)
			}

		}
		nums = append(nums, temp)
	}

	var ops []string

	for i := range last {
		if last[i] == "" || last[i] == " " {
			continue
		} else {
			ops = append(ops, last[i])
		}
	}

	result := doMath(nums, ops)

	fmt.Printf("Answer Part 1: %d\n", result)
}

func part2() {
	mathStuff, _ := util.ReadInputLines("input.txt")
	last := mathStuff[len(mathStuff)-1]
	mathStuff = mathStuff[:len(mathStuff)-1]

	var t []string
	for x := range mathStuff[0] {
		var temp string
		for y := range mathStuff {
			temp = temp + string(mathStuff[y][x])
		}

		t = append(t, temp)
	}

	var nums [][]int
	var operands []int
	for i := range t {
		if strings.TrimSpace(t[i]) != "" {
			num, _ := strconv.Atoi(strings.TrimSpace(t[i]))
			operands = append(operands, num)
		} else {
			clone := make([]int, len(operands))
			copy(clone, operands)
			nums = append(nums, clone)
			operands = operands[:0]
		}

	}

	if len(operands) > 0 {
		nums = append(nums, operands)
	}

	var ops []string

	for i := range last {
		if strings.TrimSpace(string(last[i])) == "" {
			continue
		} else {
			ops = append(ops, string(last[i]))
		}
	}

	result := doMath2(nums, ops)

	fmt.Printf("Answer Part 2: %d\n", result)

}
