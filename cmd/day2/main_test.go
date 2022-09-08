package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []string{
		"ULL",
		"RRDDD",
		"LURDL",
		"UUUUD",
	}
	assert.Equal(t, "1985", part1(input))
}
