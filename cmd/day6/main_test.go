package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input []string = []string{
	"eedadn",
	"drvtee",
	"eandsr",
	"raavrd",
	"atevrs",
	"tsrnev",
	"sdttsa",
	"rasrtv",
	"nssdts",
	"ntnada",
	"svetve",
	"tesnvt",
	"vntsnd",
	"vrdear",
	"dvrsen",
	"enarar",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, "easter", part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "advent", part2(input))
}
