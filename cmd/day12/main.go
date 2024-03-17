package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	reg := make(map[string]int)
	execute(r, reg)
	return reg["a"]
}

func part2(r io.Reader) int {
	reg := map[string]int{
		"c": 1,
	}
	execute(r, reg)
	return reg["a"]
}

func execute(r io.Reader, reg map[string]int) {
	s := bufio.NewScanner(r)
	pgm := make([]string, 0)
	for s.Scan() {
		pgm = append(pgm, s.Text())
	}

	for pc := 0; pc < len(pgm); pc++ {
		pts := strings.Fields(pgm[pc])
		switch pts[0] {
		case "cpy":
			intVal, err := strconv.Atoi(pts[1])
			if err == nil {
				reg[pts[2]] = intVal
				continue
			}
			reg[pts[2]] = reg[pts[1]]
		case "inc":
			reg[pts[1]]++
		case "dec":
			reg[pts[1]]--
		case "jnz":
			if intVal, err := strconv.Atoi(pts[1]); err == nil {
				if intVal == 0 {
					continue
				}
			} else if reg[pts[1]] == 0 {
				continue
			}
			pc-- // counteract the for loop
			jump, _ := strconv.Atoi(pts[2])
			pc += jump
		}
	}
}
