package main

import (
	"fmt"
	"math"
	"sort"
)

func countOnes(a int) (oneBits int) {
	for ; a != 0; a = a >> 1 {
		if a%2 == 1 {
			oneBits++
		}
	}
	return
}

type floorPlan struct {
	magicNumber int
}

func (f floorPlan) isWall(x, y int) bool {
	base := x*x + 3*x + 2*x*y + y + y*y
	base += f.magicNumber

	return countOnes(base)%2 == 1
}

type point struct {
	x, y int
}

func (p point) outOfBounds() bool {
	return p.x < 0 || p.y < 0
}

func (p point) add(o point) point {
	return point{p.x + o.x, p.y + o.y}
}

func (p point) eq(o point) bool {
	return p.x == o.x && p.y == o.y
}

var neighbors []point = []point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func dijkstra(start, end point, f floorPlan) int {
	visited := make(map[point]bool)

	unvisited := []point{start}

	cost := make(map[point]int)
	cost[start] = 0

	for len(unvisited) > 0 {
		c := unvisited[0]
		unvisited = unvisited[1:]

		if c.eq(end) {
			return cost[c]
		}

		visited[c] = true

		// check in each direction
		for _, n := range neighbors {
			nc := c.add(n)

			// if the neighbor is off the edge of the map...
			if nc.outOfBounds() {
				continue
			}
			// ...or we've already visited it...
			if visited[nc] {
				continue
			}
			// ...or it's actually a wall...
			if f.isWall(nc.x, nc.y) {
				continue
			}

			// ...skip it. otherwise, set its cost to be 1 more than the cost of the current node
			cost[nc] = cost[c] + 1
			unvisited = append(unvisited, nc)
		}

		// this isn't as nice as having a priority queue
		// but given the overall size of the unvisited list,
		// this is likely sufficient.
		sort.Slice(unvisited, func(i, j int) bool {
			return cost[unvisited[i]] < cost[unvisited[j]]
		})
	}

	return math.MaxInt
}

func part1(favNumber int) int {
	f := floorPlan{magicNumber: favNumber}
	start, end := point{1, 1}, point{31, 39}
	return dijkstra(start, end, f)
}

func main() {
	favNumber := 1352

	fmt.Printf("Part 1: %d\n", part1(favNumber))
}
