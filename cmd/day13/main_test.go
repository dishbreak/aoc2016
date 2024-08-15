package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOnes(t *testing.T) {
	type testCase struct {
		a    int
		bits int
	}

	testCases := []testCase{
		{0b1000111010111, 8},
		{0, 0},
		{2, 1},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.bits, countOnes(tc.a))
		})
	}
}

func TestIsWall(t *testing.T) {
	type testCase struct {
		x, y   int
		isWall bool
	}

	/*
		  0123456789
		0 .#.####.##
		1 ..#..#...#
		2 #....##...
		3 ###.#.###.
		4 .##..#..#.
		5 ..##....#.
		6 #...##.###
	*/

	testCases := []testCase{
		{4, 0, true},
		{4, 2, false},
		{5, 3, false},
		{6, 2, true},
	}

	f := floorPlan{magicNumber: 10}
	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.isWall, f.isWall(tc.x, tc.y))
		})
	}
}

func TestDijkstra(t *testing.T) {
	start, end := point{1, 1}, point{7, 4}
	f := floorPlan{magicNumber: 10}

	assert.Equal(t, 11, dijkstra(start, end, f))
}
