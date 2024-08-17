package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	salt := "ihaygndm"
	fmt.Printf("Part 1: %d\n", part1(salt))
}

var tripleChar = regexp.MustCompile(`(000|111|222|333|444|555|666|777|888|999|aaa|bbb|ccc|ddd|eee|fff)`)

func part1(salt string) int {
	var keys []int
	keysSeen := make(map[int]bool)
	idx := 0

	type checkRecord struct {
		r      *regexp.Regexp
		ttl    int
		keyIdx int
	}

	var checks []checkRecord

	for ; ; idx++ {
		val := fmt.Sprint(salt, idx)
		hash := md5.Sum([]byte(val))
		key := hex.EncodeToString(hash[:])
		for _, check := range checks {
			if check.ttl <= idx {
				continue
			}
			if keysSeen[check.keyIdx] {
				continue
			}
			if check.r.FindString(key) == "" {
				continue
			}
			keys = append(keys, check.keyIdx)
			if len(keys) == 64 {
				sort.Ints(keys)
				return keys[63]
			}
			keysSeen[check.keyIdx] = true
			check.ttl = -1 // don't run this check anymore
		}

		matches := tripleChar.FindStringSubmatch(key)
		if matches == nil {
			continue
		}

		check := checkRecord{
			ttl:    idx + 1000,
			keyIdx: idx,
			r:      regexp.MustCompile(strings.Repeat(string(matches[0][0]), 5)),
		}
		checks = append(checks, check)

	}

}
