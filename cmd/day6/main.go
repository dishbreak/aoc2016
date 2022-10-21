package main

import (
	"fmt"
	"sync"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day6.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %s\n", part1(input))
}

func part1(input []string) string {
	slots := len(input[0])
	letterStreams := make([]chan rune, slots)

	type letterReport struct {
		pos int
		r   rune
	}

	reports := make(chan letterReport)

	var wg sync.WaitGroup
	wg.Add(slots)
	go func() {
		wg.Wait()
		close(reports)
	}()

	for i := range letterStreams {
		letterStreams[i] = make(chan rune, len(input))
		go func(input <-chan rune, pos int) {
			defer wg.Done()

			hits := make(map[rune]int, 26)
			for b := range input {
				hits[b]++
			}

			max := -1
			var r rune
			for b, ct := range hits {
				if ct > max {
					max = ct
					r = b
				}
			}

			reports <- letterReport{
				r:   r,
				pos: pos,
			}

		}(letterStreams[i], i)
	}

	result := make(chan string)
	go func() {
		defer close(result)
		msg := make([]byte, slots)

		for report := range reports {
			msg[report.pos] = byte(report.r)
		}

		result <- string(msg)
	}()

	for _, val := range input {
		if val == "" {
			continue
		}

		for i, c := range val {
			letterStreams[i] <- c
		}
	}

	for i := range letterStreams {
		close(letterStreams[i])
	}

	return <-result
}
