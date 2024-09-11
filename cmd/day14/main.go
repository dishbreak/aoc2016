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
	fmt.Printf("Part 2: %d\n", part2(salt))
}

func part1(salt string) int {
	return findKey(salt, standardMd5)
}

func part2(salt string) int {
	return findKey(salt, stretchedMd5)
}

var tripleChar = regexp.MustCompile(`(000|111|222|333|444|555|666|777|888|999|aaa|bbb|ccc|ddd|eee|fff)`)

type hashFunc func(salt string, idx int) string

func standardMd5(salt string, idx int) string {
	val := fmt.Sprint(salt, idx)
	hash := md5.Sum([]byte(val))
	return hex.EncodeToString(hash[:])
}

func stretchedMd5(salt string, idx int) string {
	h := standardMd5(salt, idx)
	for i := 0; i < 2016; i++ {
		b := md5.Sum([]byte(h))
		h = hex.EncodeToString(b[:])
	}
	return h
}

func findKey(salt string, hasher hashFunc) int {
	var keys []int
	idx := 0

	type checkRecord struct {
		r      *regexp.Regexp
		ttl    int
		keyIdx int
	}

	var checks []checkRecord

	for ; ; idx++ {
		key := hasher(salt, idx)
		for _, check := range checks {
			if check.ttl <= idx {
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
