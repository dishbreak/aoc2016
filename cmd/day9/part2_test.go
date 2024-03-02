package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	type testCase struct {
		input  string
		result int
	}

	testCases := []testCase{
		{"(3x3)XYZ", 9},
		{"X(8x2)(3x3)ABCY", 20},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			r := strings.NewReader(tc.input)
			assert.Equal(t, tc.result, part2(r))
		})
	}
}
