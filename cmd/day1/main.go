package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input[0]))
	fmt.Printf("Part 2: %d\n", part2(input[0]))
}

func manhattanDist(p image.Point) int {
	abs := func(i int) int {
		if i < 0 {
			return -1 * i
		}
		return i
	}

	return abs(p.X) + abs(p.Y)
}

func part1(input string) int {
	dirs := []image.Point{
		image.Pt(0, 1),
		image.Pt(1, 0),
		image.Pt(0, -1),
		image.Pt(-1, 0),
	}

	vIdx := 0
	pt := image.Point{}

	for _, inst := range strings.Split(input, ", ") {
		dir := inst[0]
		steps, _ := strconv.Atoi(inst[1:])

		if dir == 'L' {
			vIdx--
			if vIdx < 0 {
				vIdx = 3
			}
		} else {
			vIdx = (vIdx + 1) % len(dirs)
		}

		pt = pt.Add(dirs[vIdx].Mul(steps))
	}

	return manhattanDist(pt)
}

func part2(input string) int {
	dirs := []image.Point{
		image.Pt(0, 1),
		image.Pt(1, 0),
		image.Pt(0, -1),
		image.Pt(-1, 0),
	}

	vIdx := 0
	pt := image.Point{}

	hits := make(map[image.Point]bool)

	for _, inst := range strings.Split(input, ", ") {
		dir := inst[0]
		steps, _ := strconv.Atoi(inst[1:])

		if dir == 'L' {
			vIdx--
			if vIdx < 0 {
				vIdx = 3
			}
		} else {
			vIdx = (vIdx + 1) % len(dirs)
		}

		for i := 0; i < steps; i++ {
			pt = pt.Add(dirs[vIdx])
			if hits[pt] {
				return manhattanDist(pt)
			}

			hits[pt] = true
		}
	}

	return -1
}
