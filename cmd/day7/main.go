package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day7.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	acc := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		addr := IPV7AddressFromString(line)
		if addr.SupportsTLS() {
			acc++
		}
	}
	return acc
}

func part2(input []string) int {
	acc := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		addr := IPV7AddressFromString(line)
		if addr.SupportsSSL() {
			acc++
		}
	}
	return acc
}
