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

type state struct {
	curr           int
	times_hit_zero int
}

func processRotation_1(rot string, ctx *state) error {
	dir := rot[0]
	num, err := strconv.Atoi(rot[1:])
	if err != nil {
		fmt.Println("Couldn't convert the rotation integer ")
		return err
	}

	if dir == 'L' {
		ctx.curr -= num
	}

	if dir == 'R' {
		ctx.curr += num
	}

	ctx.curr = ctx.curr % 100

	if ctx.curr == 0 {
		ctx.times_hit_zero += 1
	}

	return nil
}

func processRotation_2(rot string, ctx *state) error {
	dir := rot[0]
	num, err := strconv.Atoi(rot[1:])
	if err != nil {
		fmt.Println("Couldn't convert the rotation integer ")
		return err
	}

	ctx.times_hit_zero += num / 100

	for range num % 100 {
		if dir == 'L' {
			ctx.curr -= 1
			ctx.curr = ctx.curr % 100
		}

		if dir == 'R' {
			ctx.curr += 1
			ctx.curr = ctx.curr % 100
		}

		if ctx.curr == 0 {
			ctx.times_hit_zero += 1
		}
	}

	return nil
}

func part1() {
	ctx := state{curr: 50, times_hit_zero: 0}
	rots, err := util.ReadInputLines("input.txt")
	if err != nil {
		fmt.Printf("Couldn't read input.txt: %d\n", err)
		return
	}

	for i := 0; i < len(rots); i++ {
		processRotation_1(rots[i], &ctx)

	}

	fmt.Printf("Answer Part 1: %d\n", ctx.times_hit_zero)

}

func part2() {
	ctx := state{curr: 50, times_hit_zero: 0}

	rots, err := util.ReadInputLines("input.txt")
	if err != nil {
		fmt.Printf("Couldn't read input.txt: %d\n", err)
	}

	for i := 0; i < len(rots); i++ {
		processRotation_2(rots[i], &ctx)

	}

	fmt.Printf("Answer Part 2: %d\n", ctx.times_hit_zero)
}
