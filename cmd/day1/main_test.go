package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	assert.Equal(t, 4, part2("R8, R4, R4, R8"))
}
