package main

import (
	"fmt"
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScreen(t *testing.T) {
	type testCase struct {
		input        []string
		instructions []string
		result       string
	}

	testCases := []testCase{
		{
			instructions: []string{
				"rect 3x3",
			},
			result: `###   
###   
###   
`,
		},
		{
			instructions: []string{
				"rect 3x4",
			},
			result: `###   
###   
###   
`,
		},
		{
			input: []string{
				"##    ",
				"##    ",
				"      ",
			},
			instructions: []string{
				"rotate column y=1 by 2",
			},
			result: `##    
#     
 #    
`,
		},
		{
			input: []string{
				"##    ",
				"##    ",
				"      ",
			},
			instructions: []string{
				"rotate row x=1 by 8",
			},
			result: `##    
  ##  
      
`,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			scr := NewTinyScreen(WithInitialValue(tc.input), WithDimensions(image.Pt(6, 3)))
			for _, line := range tc.instructions {
				scr.Inst(line)
			}
			assert.Equal(t, tc.result, scr.String())
		})
	}
}
