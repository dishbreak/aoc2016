package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, "18f47a30", part1("abc"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "05ace8e3", part2("abc"))
}
