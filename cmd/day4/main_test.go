package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRoomID(t *testing.T) {
	input := "aaaaa-bbb-z-y-x-123[abxyz]"
	expected := roomID{
		raw:      "aaaaa-bbb-z-y-x-123[abxyz]",
		SectorID: 123,
		checksum: "abxyz",
		parts:    []string{"aaaaa", "bbb", "z", "y", "x"},
	}
	assert.Equal(t, expected, parseRoomID(input))
}

func TestRoomIDValid(t *testing.T) {
	type testCase struct {
		input    string
		expected bool
	}

	testCases := []testCase{
		{"aaaaa-bbb-z-y-x-123[abxyz]", true},
		{"a-b-c-d-e-f-g-h-987[abcde]", true},
		{"not-a-real-room-404[oarel]", true},
		{"totally-real-room-200[decoy]", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			rid := parseRoomID(tc.input)
			assert.Equal(t, tc.expected, rid.IsValid())
		})
	}
}

func TestDecryptWord(t *testing.T) {
	assert.Equal(t, "encrypted", decrypt("zixmtkozy", 343))
}
