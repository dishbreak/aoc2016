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
	fmt.Printf("Part 2: %s\n", part2(input))
}

func makeLetterStreams(input []string) []chan rune {
	slots := len(input[0])
	letterStreams := make([]chan rune, slots)

	for i := range letterStreams {
		letterStreams[i] = make(chan rune, len(input))
	}

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

	return letterStreams
}

type analyzeLettersConfig struct {
	comparator func(int, int) bool
	start      int
}

type analyzeLettersOption func(*analyzeLettersConfig)

func UsingMostRepeatedCharacter() analyzeLettersOption {
	return func(alc *analyzeLettersConfig) {
		alc.comparator = func(i1, i2 int) bool {
			return i1 > i2
		}
		alc.start = -1
	}
}

func UsingLeastRepeatedCharacter() analyzeLettersOption {
	return func(alc *analyzeLettersConfig) {
		alc.comparator = func(i1, i2 int) bool {
			return i1 < i2
		}
		alc.start = 10000000000000
	}
}

type letterReport struct {
	pos int
	r   rune
}

func analyzeLetters(letterStreams []chan rune, opts ...analyzeLettersOption) <-chan letterReport {
	c := &analyzeLettersConfig{}

	for _, opt := range opts {
		opt(c)
	}

	if c.comparator == nil {
		UsingMostRepeatedCharacter()(c)
	}

	reports := make(chan letterReport)

	var wg sync.WaitGroup
	wg.Add(len(letterStreams))
	go func() {
		wg.Wait()
		close(reports)
	}()

	for i := range letterStreams {
		go func(input <-chan rune, pos int) {
			defer wg.Done()

			hits := make(map[rune]int, 26)
			for b := range input {
				hits[b]++
			}

			hc := c.start
			var r rune
			for b, ct := range hits {
				if c.comparator(ct, hc) {
					hc = ct
					r = b
				}
			}

			reports <- letterReport{
				r:   r,
				pos: pos,
			}

		}(letterStreams[i], i)
	}

	return reports
}

func generateResult(reports <-chan letterReport, slots int) <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		msg := make([]byte, slots)

		for report := range reports {
			msg[report.pos] = byte(report.r)
		}

		result <- string(msg)
	}()

	return result
}

func decode(input []string, opt analyzeLettersOption) string {
	letterStreams := makeLetterStreams(input)
	reports := analyzeLetters(letterStreams, opt)
	return <-generateResult(reports, len(letterStreams))
}

func part1(input []string) string {
	return decode(input, UsingMostRepeatedCharacter())
}

func part2(input []string) string {
	return decode(input, UsingLeastRepeatedCharacter())
}
