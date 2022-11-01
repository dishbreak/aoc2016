package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFsmParse(t *testing.T) {
	type testCase struct {
		input    string
		expected int
	}

	testCases := []testCase{
		{"(3x3)XYZ", 9},
		{"A(2x2)BCD(2x2)EFG", 11},
		{"(6x1)(1x3)A", 6},
		{`X(8x2)(3x3)ABCY`, 18},
	}

	for i, tc := range testCases {
		tcName := fmt.Sprintf("test case %d", i)
		t.Run(tcName, func(t *testing.T) {
			fsm := NewFsmParser()
			buf := strings.NewReader(tc.input)
			fsm.Parse(buf)
			assert.Equal(t, tc.expected, fsm.CharCount())
		})
	}
}
