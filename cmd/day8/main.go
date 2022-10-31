package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day8.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	scr := NewTinyScreen()

	for _, line := range input {
		if line == "" {
			continue
		}
		scr.Inst(line)
	}

	return scr.PixelsLit()
}
