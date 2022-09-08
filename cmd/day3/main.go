package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	acc := 0
	for _, item := range input {
		if item == "" {
			continue
		}
		if isValid(item) {
			acc++
		}
	}
	return acc
}

func isValid(line string) bool {
	var sides [3]int
	for i, pt := range strings.Fields(line) {
		sides[i], _ = strconv.Atoi(pt)
	}

	combos := [][]int{
		{0, 1, 2},
		{1, 2, 0},
		{2, 0, 1},
	}

	for _, combo := range combos {
		if (sides[combo[0]] + sides[combo[1]]) <= sides[combo[2]] {
			return false
		}
	}

	return true
}
