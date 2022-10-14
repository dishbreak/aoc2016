package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	acc := 0
	for _, item := range input {
		if item == "" {
			continue
		}
		if validLIne(item) {
			acc++
		}
	}
	return acc
}

func part2(input []string) int {
	triangleStreams := make([]chan string, 3)
	for i := range triangleStreams {
		triangleStreams[i] = make(chan string, len(input))
	}

	parsed := make(chan [3]int)

	var wg sync.WaitGroup
	wg.Add(3)
	for i := range triangleStreams {
		go func(input chan string) {
			defer wg.Done()

			count := 0
			var buf [3]int
			for {
				s, ok := <-input
				if !ok {
					return
				}
				n, _ := strconv.Atoi(s)
				buf[count] = n
				count++

				if count == 3 {
					parsed <- buf
					buf = [3]int{}
					count = 0
				}
			}
		}(triangleStreams[i])
	}

	go func() {
		wg.Wait()
		close(parsed)
	}()

	result := make(chan int)
	defer close(result)

	go func() {
		acc := 0
		for triangle := range parsed {
			if isValid(triangle) {
				acc++
			}
		}
		result <- acc
	}()

	for _, line := range input {
		if line == "" {
			continue
		}

		for i, f := range strings.Fields(line) {
			triangleStreams[i] <- f
		}
	}

	for i := range triangleStreams {
		close(triangleStreams[i])
	}

	return <-result

}

func validLIne(line string) bool {
	var sides [3]int
	for i, pt := range strings.Fields(line) {
		sides[i], _ = strconv.Atoi(pt)
	}

	return isValid(sides)
}

func isValid(sides [3]int) bool {
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
