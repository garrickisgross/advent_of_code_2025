package main

import (
	"fmt"
	"math/big"

	"github.com/garrickisgross/advent_of_code_2025/util"
)

func main() {
	grid := readInput()
	fmt.Printf("Answer Part 1: %d\n", countSplits(grid))
	fmt.Printf("Answer Part 2: %s\n", countTimelines(grid))
}

// countSplits simulates all beams moving downward and splitting at each '^'.
// It returns the total number of times any beam hits a splitter.
func countSplits(grid [][]string) int {
	startY, startX, ok := findStart(grid)
	if !ok {
		return 0
	}

	// Track active beams by column for the current row to merge duplicates.
	current := map[int]struct{}{startX: {}}
	splits := 0

	for y := startY + 1; y < len(grid); y++ {
		next := make(map[int]struct{})
		rowLen := len(grid[y])

		for x := range current {
			if x < 0 || x >= rowLen {
				continue
			}
			if grid[y][x] == "^" {
				splits++
				if x-1 >= 0 {
					next[x-1] = struct{}{}
				}
				if x+1 < rowLen {
					next[x+1] = struct{}{}
				}
			} else {
				next[x] = struct{}{}
			}
		}

		current = next

		if len(current) == 0 {
			break
		}
	}

	return splits
}

// countTimelines returns the number of possible timelines for a single particle
// that splits left/right at each '^'.
func countTimelines(grid [][]string) string {
	startY, startX, ok := findStart(grid)
	if !ok {
		return "0"
	}

	// Track timeline counts per column using big ints to avoid overflow.
	current := map[int]*big.Int{startX: big.NewInt(1)}

	for y := startY + 1; y < len(grid); y++ {
		next := make(map[int]*big.Int)
		rowLen := len(grid[y])

		for x, cnt := range current {
			if x < 0 || x >= rowLen {
				continue
			}
			if grid[y][x] == "^" {
				if x-1 >= 0 {
					addTo(next, x-1, cnt)
				}
				if x+1 < rowLen {
					addTo(next, x+1, cnt)
				}
			} else {
				addTo(next, x, cnt)
			}
		}

		current = next

		if len(current) == 0 {
			break
		}
	}

	total := big.NewInt(0)
	for _, cnt := range current {
		total.Add(total, cnt)
	}

	return total.String()
}

func readInput() [][]string {
	return util.ReadInput2dArray("input.txt", "")
}

func findStart(grid [][]string) (int, int, bool) {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "S" {
				return y, x, true
			}
		}
	}
	return 0, 0, false
}

func addTo(m map[int]*big.Int, x int, delta *big.Int) {
	if existing, ok := m[x]; ok {
		existing.Add(existing, delta)
		return
	}
	m[x] = new(big.Int).Set(delta)
}
