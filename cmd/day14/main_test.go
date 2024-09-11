package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 22728, part1("abc"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 22551, part2("abc"))
}
