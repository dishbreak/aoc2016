package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	acc := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		rid := parseRoomID(line)
		if rid.IsValid() {
			acc += rid.SectorID
		}
	}

	return acc
}

type roomID struct {
	raw      string
	SectorID int
	checksum string
	parts    []string
}

var sectorSum *regexp.Regexp = regexp.MustCompile(`(\d+)\[([a-z]+)]`)

func parseRoomID(identifier string) roomID {
	r := roomID{
		raw: identifier,
	}

	parts := strings.Split(identifier, "-")
	r.parts = parts[:len(parts)-1]

	matches := sectorSum.FindSubmatch([]byte(parts[len(parts)-1]))
	r.SectorID, _ = strconv.Atoi(string(matches[1]))
	r.checksum = string(matches[2])

	return r
}

func (r roomID) IsValid() bool {
	hitCount := make(map[rune]int)

	for _, part := range r.parts {
		for _, c := range part {
			hitCount[c]++
		}
	}

	type hitRecord struct {
		r    rune
		hits int
	}

	hits := make([]hitRecord, 0)

	for r, c := range hitCount {
		hits = append(hits, hitRecord{r, c})
	}

	sort.Slice(hits, func(i, j int) bool {
		if hits[i].hits > hits[j].hits {
			return true
		}
		if hits[i].hits == hits[j].hits && hits[i].r < hits[j].r {
			return true
		}
		return false
	})

	for i, r := range r.checksum {
		if hits[i].r != r {
			return false
		}
	}

	return true
}

func (r roomID) Decrypt() string {
	plaintext := make([]string, len(r.parts))

	for i, part := range r.parts {
		plaintext[i] = decrypt(part, r.SectorID)
	}

	return strings.Join(plaintext, " ")
}

func part2(input []string) int {
	for _, line := range input {
		if line == "" {
			continue
		}

		rid := parseRoomID(line)
		if !rid.IsValid() {
			continue
		}

		name := rid.Decrypt()
		if strings.HasPrefix(name, "northpole") {
			return rid.SectorID
		}
	}
	return 0
}

func decrypt(word string, shift int) string {
	buf := make([]byte, len(word))

	for i, c := range word {
		effective := byte(((int(c) - 'a') + shift) % 26)
		buf[i] = 'a' + effective
	}

	return string(buf)
}
