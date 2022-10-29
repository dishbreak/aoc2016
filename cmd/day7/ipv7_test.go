package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPV7AddressFromString(t *testing.T) {
	type testCase struct {
		input    string
		expected IPV7Address
	}

	testCases := []testCase{
		{
			input: "pjdxcbutwijvtoftvw[zkqtzecoenkibees]llfxdbldntlydpvvn[uaweaigkebxceixszbh]xxlipjtlogbnxse",
			expected: IPV7Address{
				raw: "pjdxcbutwijvtoftvw[zkqtzecoenkibees]llfxdbldntlydpvvn[uaweaigkebxceixszbh]xxlipjtlogbnxse",
				outer: []string{
					"pjdxcbutwijvtoftvw",
					"llfxdbldntlydpvvn",
					"xxlipjtlogbnxse",
				},
				hnet: []string{
					"zkqtzecoenkibees",
					"uaweaigkebxceixszbh",
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, IPV7AddressFromString(tc.input))
		})
	}

}

func TestAbbaDetected(t *testing.T) {
	type testCase struct {
		input    string
		expected bool
	}

	testCases := []testCase{
		{"ghkhalfdsh", false},
		{"asfloolajf", true},
		{"abbadfadaf", true},
		{"afdfaffaaf", true},
		{"afdfafflaf", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, abbaDetected(tc.input))
		})
	}
}
