package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("inputs/day9.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(f))
}

func part1(f *os.File) int {
	b := bufio.NewReader(f)
	fsm := NewFsmParser()
	fsm.Parse(b)
	return fsm.CharCount()
}
