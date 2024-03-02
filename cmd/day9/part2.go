package main

import (
	"bufio"
	"io"
	"strconv"
)

func part2(r io.Reader) int {
	b := bufio.NewScanner(r)
	b.Split(bufio.ScanRunes)

	c := &CountScanner{
		Scanner: b,
	}

	count := 0
	for c.Scan() {
		t := c.Text()
		if t == "(" {
			count += readSubseq(c)
			continue
		}
		count++
	}

	return count
}

type CountScanner struct {
	*bufio.Scanner
	idx int
}

func (c *CountScanner) Scan() bool {
	c.idx++
	return c.Scanner.Scan()
}

func (c *CountScanner) Pos() int {
	return c.idx
}

func readSubseq(c *CountScanner) (count int) {
	seq, mult := parseHeader(c)
	limit := c.Pos() + seq

	for c.Pos() < limit && c.Scan() {
		t := c.Text()
		if t == "(" {
			count += readSubseq(c)
			continue
		}
		count++
	}

	count *= mult
	return
}

func parseHeader(b *CountScanner) (seq, mult int) {
	buf := 0

	for b.Scan() {
		c := b.Text()
		switch c {
		case "x":
			seq = buf
			buf = 0
		case ")":
			mult = buf
			return
		default:
			digit, _ := strconv.Atoi(c)
			buf = buf*10 + digit
		}
	}
	return
}
