package main

import (
	"fmt"

	"github.com/garrickisgross/advent_of_code_2025/util"
)

func main() {
	part1()
	part2()
}

func canAccess(y int, x int, grid [][]string) bool {
	total := 0

	// up-left
	if x-1 >= 0 && y-1 >= 0 && grid[y-1][x-1] == "@" {
		total += 1
	}
	// up
	if y-1 >= 0 && grid[y-1][x] == "@" {
		total += 1
	}
	// up-right
	if y-1 >= 0 && x+1 < len(grid[0]) && grid[y-1][x+1] == "@" {
		total += 1
	}
	// right
	if x+1 < len(grid[0]) && grid[y][x+1] == "@" {
		total += 1
	}
	// down-right
	if y+1 < len(grid) && x+1 < len(grid[0]) && grid[y+1][x+1] == "@" {
		total += 1
	}
	// down
	if y+1 < len(grid) && grid[y+1][x] == "@" {
		total += 1
	}
	// down-left
	if y+1 < len(grid) && x-1 >= 0 && grid[y+1][x-1] == "@" {
		total += 1
	}
	// left
	if x-1 >= 0 && grid[y][x-1] == "@" {
		total += 1
	}
	return total < 4
}

func part1() {
	input_grid := util.ReadInput2dArray("input.txt")
	acc := 0
	for i := 0; i < len(input_grid); i++ {
		for j := 0; j < len(input_grid[0]); j++ {
			if input_grid[i][j] == "@" && canAccess(i, j, input_grid) {
				acc += 1
			}
		}
	}

	fmt.Printf("Answer Part 1: %d\n", acc)
}

func part2() {
	input_grid := util.ReadInput2dArray("input.txt")
	acc := 0
	done := false

	for !done {
		done = true

		for i := 0; i < len(input_grid); i++ {
			for j := 0; j < len(input_grid[0]); j++ {
				if input_grid[i][j] == "@" && canAccess(i, j, input_grid) {
					input_grid[i][j] = "."
					done = false
					acc += 1
					continue
				}

			}
		}

	}

	fmt.Printf("Answer Part 2: %d\n", acc)
}
