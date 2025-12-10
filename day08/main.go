package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/garrickisgross/advent_of_code_2025/util"
)

type box struct {
	x    float64
	y    float64
	z    float64
	conn []*box
}

type pair struct {
	i    int
	j    int
	dist float64
}

func main() {
	part1(1000)
	part2()
}

func readInput(path string) []box {
	input, _ := util.ReadInputLines(path)
	var boxes []box
	for _, line := range input {
		values := strings.Split(line, ",")
		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])

		box := box{x: float64(x), y: float64(y), z: float64(z)}

		boxes = append(boxes, box)
	}

	return boxes
}

func cmp_distance(a pair, b pair) bool {
	// returns true if a < b
	return a.dist < b.dist
}

func getdistance(b1 box, b2 box) float64 {
	dx := b1.x - b2.x
	dy := b1.y - b2.y
	dz := b1.z - b2.z

	result := math.Sqrt((dx * dx) + (dy * dy) + (dz * dz))

	return result
}

func part1(conn int) {
	boxes := readInput("input.txt")
	var pairs []pair

	for i := range boxes {
		for j := i + 1; j <= len(boxes)-1; j++ { // ignore current box
			dist := getdistance(boxes[i], boxes[j])
			pairs = append(pairs, pair{i: i, j: j, dist: dist})
		}
	}

	fmt.Println(len(pairs))
	// sort the pairs
	quicksort(&pairs, 0, len(pairs)-1)
	// add connections.
	parent := make([]int, len(boxes))
	size := make([]int, len(boxes))

	for i := range boxes {
		parent[i] = i
		size[i] = 1
	}

	proccessed := 0
	for proccessed < conn {
		i := pairs[proccessed].i
		j := pairs[proccessed].j

		union(i, j, &parent, &size)

		proccessed += 1
	}

	var circuitSizes []int
	for i := range boxes {
		if parent[i] == i {
			circuitSizes = append(circuitSizes, size[i])
		}
	}

	sort.Ints(circuitSizes)
	first := circuitSizes[len(circuitSizes)-1]
	second := circuitSizes[len(circuitSizes)-2]
	third := circuitSizes[len(circuitSizes)-3]

	fmt.Printf("Answer Part 1: %d\n", first*second*third)

}

func part2() {
	boxes := readInput("input.txt")
	var pairs []pair

	for i := range boxes {
		for j := i + 1; j <= len(boxes)-1; j++ { // ignore current box
			dist := getdistance(boxes[i], boxes[j])
			pairs = append(pairs, pair{i: i, j: j, dist: dist})
		}
	}

	fmt.Println(len(pairs))
	// sort the pairs
	quicksort(&pairs, 0, len(pairs)-1)
	// add connections.
	parent := make([]int, len(boxes))
	size := make([]int, len(boxes))

	for i := range boxes {
		parent[i] = i
		size[i] = 1
	}

	proccessed := len(boxes)
	idx := 0
	result := 0.0
	for {

		i := pairs[idx].i
		j := pairs[idx].j

		madeMerge := union(i, j, &parent, &size)

		if madeMerge {
			proccessed -= 1
		}

		if proccessed == 1 {
			result = boxes[pairs[idx].i].x * boxes[pairs[idx].j].x
			break
		}

		idx += 1
	}

	fmt.Printf("Answer Part 2: %d\n", int(result))

}

func find(i int, parent []int) int {
	cur := i
	for parent[cur] != cur {
		cur = parent[cur]
	}
	return cur
}

func union(a int, b int, parent *[]int, size *[]int) bool {
	ra := find(a, *parent)
	rb := find(b, *parent)

	if ra == rb {
		return false
	}

	if (*size)[ra] > (*size)[rb] {
		(*parent)[rb] = ra
		(*size)[ra] += (*size)[rb]
	} else {
		(*parent)[ra] = rb
		(*size)[rb] += (*size)[ra]
	}

	return true

}

// Quick sort implementation
func partition(arr *[]pair, low int, high int) int {
	pivot := (*arr)[high]

	i := low - 1

	for j := low; j < high; j++ {
		if (*arr)[j].dist < pivot.dist {
			i += 1
			swap(arr, i, j)
		}
	}

	swap(arr, i+1, high)
	return i + 1
}

func swap(arr *[]pair, i int, j int) {
	temp := (*arr)[j]

	(*arr)[j] = (*arr)[i]

	(*arr)[i] = temp
}

func quicksort(arr *[]pair, low int, high int) {

	if low < high {
		pi := partition(arr, low, high)

		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}

}
