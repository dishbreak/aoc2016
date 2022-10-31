package main

import (
	"image"
	"regexp"
	"strconv"
	"strings"
)

type TinyScreen struct {
	pixels map[image.Point]bool
	maxDim image.Point
}

var rotateInst *regexp.Regexp = regexp.MustCompile(`^rotate (row|column) [xy]=(\d+) by (\d+)$`)
var rectInst *regexp.Regexp = regexp.MustCompile(`^rect (\d+)x(\d+)$`)

type TinyScreenOption func(t *TinyScreen)

func WithInitialValue(input []string) TinyScreenOption {
	return func(t *TinyScreen) {
		if input == nil {
			return
		}
		for y, line := range input {
			for x, c := range line {
				t.pixels[image.Pt(x, y)] = c == '#'
			}
		}
	}
}

func WithDimensions(dims image.Point) TinyScreenOption {
	return func(t *TinyScreen) {
		t.maxDim = dims
	}
}

func NewTinyScreen(opts ...TinyScreenOption) *TinyScreen {
	t := &TinyScreen{
		pixels: make(map[image.Point]bool),
		maxDim: image.Pt(50, 6),
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func (t *TinyScreen) Inst(instruction string) {
	switch {
	case rectInst.MatchString(instruction):
		t.rect(instruction)
	case rotateInst.MatchString(instruction):
		t.rotate(instruction)
	}
}

func min(one, other int) int {
	if one < other {
		return one
	}
	return other
}

func (t *TinyScreen) rect(instruction string) {
	matches := rectInst.FindStringSubmatch(instruction)

	var dims image.Point
	dims.X, _ = strconv.Atoi(matches[1])
	dims.Y, _ = strconv.Atoi(matches[2])

	dims.X, dims.Y = min(dims.X, t.maxDim.X), min(dims.Y, t.maxDim.Y)

	for y := 0; y < dims.Y; y++ {
		for x := 0; x < dims.X; x++ {
			t.pixels[image.Pt(x, y)] = true
		}
	}
}

func (t *TinyScreen) rotate(instruction string) {
	matches := rotateInst.FindStringSubmatch(instruction)

	var start, end, direction image.Point
	coord, _ := strconv.Atoi(matches[2])
	offset, _ := strconv.Atoi(matches[3])

	switch matches[1] {
	case "row":
		start.Y = coord
		end.Y = coord
		end.X = t.maxDim.X
		direction.X = 1
	case "column":
		start.X = coord
		end.X = coord
		end.Y = t.maxDim.Y
		direction.Y = 1
	}

	shift := direction.Mul(offset)

	buf := make(map[image.Point]bool)
	for n := start; n != end; n = n.Add(direction) {
		buf[n] = t.pixels[n]
	}

	for n, val := range buf {
		p := n.Add(shift)
		p.X, p.Y = p.X%t.maxDim.X, p.Y%t.maxDim.Y
		t.pixels[p] = val
	}
}

func (t *TinyScreen) String() string {
	var sb strings.Builder
	for y := 0; y < t.maxDim.Y; y++ {
		for x := 0; x < t.maxDim.X; x++ {
			if !t.pixels[image.Pt(x, y)] {
				sb.WriteRune(' ')
				continue
			}
			sb.WriteRune('#')
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (t *TinyScreen) PixelsLit() int {
	acc := 0
	for _, val := range t.pixels {
		if val {
			acc++
		}
	}

	return acc
}
