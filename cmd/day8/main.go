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
	fmt.Printf("Part 2: \n%s\n", part2(input))
}

func part1(input []string) int {
	return loadScreen(input).PixelsLit()
}

func part2(input []string) string {
	return loadScreen(input).String()
}

func loadScreen(input []string) *TinyScreen {
	scr := NewTinyScreen()

	for _, line := range input {
		if line == "" {
			continue
		}
		scr.Inst(line)
	}

	return scr
}
