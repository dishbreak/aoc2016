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
	f, err := os.Open("inputs/day10.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

type bot struct {
	id       int
	targetId [2]int
	chips    [2]int
	isOutput [2]bool
	full     bool
}

func (b *bot) add(val int) {
	if b.full {
		panic(fmt.Errorf("bot %d can't take value %d -- out of space", b.id, val))
	}

	if b.chips[0] == 0 {
		b.chips[0] = val
		return
	}

	b.chips[1] = val
	if b.chips[0] > b.chips[1] {
		b.chips[0], b.chips[1] = b.chips[1], b.chips[0]
	}
	b.full = true
}

func (b *bot) flush() (r []addInst) {
	r = make([]addInst, 2)
	for i := range r {
		r[i] = addInst{targetId: b.targetId[i], value: b.chips[i], outputBin: b.isOutput[i]}
	}
	b.chips = [2]int{0, 0}
	b.full = false
	return
}

type addInst struct {
	value, targetId int
	outputBin       bool
}

func parse(r io.Reader) (bots map[int]*bot, q []addInst) {
	bots = make(map[int]*bot)
	q = make([]addInst, 0)

	s := bufio.NewScanner(r)
	for s.Scan() {
		pts := strings.Fields(s.Text())

		switch pts[0] {
		case "value":
			targetId, _ := strconv.Atoi(pts[5])
			value, _ := strconv.Atoi(pts[1])
			q = append(q, addInst{value: value, targetId: targetId})
		case "bot":
			b := &bot{}
			b.id, _ = strconv.Atoi(pts[1])
			b.targetId[0], _ = strconv.Atoi(pts[6])
			b.targetId[1], _ = strconv.Atoi(pts[11])
			b.isOutput[0] = pts[5] == "output"
			b.isOutput[1] = pts[10] == "output"
			bots[b.id] = b
		}
	}
	return
}

func part1(r io.Reader) int {

	bots, q := parse(r)

	for len(q) != 0 {
		inst := q[0]
		q = q[1:]
		bot := bots[inst.targetId]
		bot.add(inst.value)
		if bot.full {
			if bot.chips == [2]int{17, 61} {
				return bot.id
			}
			q = append(q, bot.flush()...)
		}
	}

	return -1
}

func part2(r io.Reader) int {
	bots, q := parse(r)
	outputs := make(map[int]int, 0)

	for len(q) != 0 {
		inst := q[0]
		q = q[1:]
		if inst.outputBin {
			outputs[inst.targetId] = inst.value
			continue
		}

		b := bots[inst.targetId]
		b.add(inst.value)
		if b.full {
			q = append(q, b.flush()...)
		}
	}

	acc := 1

	for i := 0; i < 3; i++ {
		acc *= outputs[i]
	}

	return acc
}
