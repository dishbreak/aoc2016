package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	input = input[:len(input)-1]

	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %s\n", part2(input))
}

func part1(input []string) string {
	keypad := map[image.Point]rune{
		image.Pt(0, 0): '1',
		image.Pt(1, 0): '2',
		image.Pt(2, 0): '3',
		image.Pt(0, 1): '4',
		image.Pt(1, 1): '5',
		image.Pt(2, 1): '6',
		image.Pt(0, 2): '7',
		image.Pt(1, 2): '8',
		image.Pt(2, 2): '9',
	}

	return decode(input, image.Pt(1, 1), keypad)
}

func part2(input []string) string {
	/*
		  0 1 2 3 4
		0     1
		1   2 3 4
		2 5 6 7 8 9
		3   A B C
		4     D
	*/
	keypad := map[image.Point]rune{
		image.Pt(2, 0): '1',
		image.Pt(1, 1): '2',
		image.Pt(2, 1): '3',
		image.Pt(3, 1): '4',
		image.Pt(0, 2): '5',
		image.Pt(1, 2): '6',
		image.Pt(2, 2): '7',
		image.Pt(3, 2): '8',
		image.Pt(4, 2): '9',
		image.Pt(1, 3): 'A',
		image.Pt(2, 3): 'B',
		image.Pt(3, 3): 'C',
		image.Pt(2, 4): 'D',
	}

	return decode(input, image.Pt(0, 2), keypad)
}

func decode(input []string, start image.Point, keypad map[image.Point]rune) string {
	var sb strings.Builder

	pt := start

	for _, line := range input {
		pt = trace(pt, line, keypad)
		sb.WriteRune(keypad[pt])
	}

	return sb.String()
}

func trace(start image.Point, inst string, keypad map[image.Point]rune) image.Point {
	dirs := map[rune]image.Point{
		'U': image.Pt(0, -1),
		'D': image.Pt(0, 1),
		'L': image.Pt(-1, 0),
		'R': image.Pt(1, 0),
	}

	pt := start

	for _, c := range inst {
		v := dirs[c]
		nextPt := pt.Add(v)
		if _, ok := keypad[nextPt]; !ok {
			continue
		}
		pt = nextPt
	}

	return pt
}
