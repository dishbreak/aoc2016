package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %s\n", part1(input[0]))
	fmt.Printf("Part 2: %s\n", part2(input[0]))
}

func part1(input string) string {
	return bruteForce(input, checkForMatchingHashPt1, &pt1Assembler{
		buf: make([]hashRecord, 0),
	})
}

func part2(input string) string {
	return bruteForce(input, checkForMatchingHashPt2, &pt2Assembler{})
}

func bruteForce(input string, hr hashRunner, asm passwordAssembler) string {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	index := countStream(ctx)
	hashHits := runHash(ctx, index, input, hr)
	return <-assemblePassword(ctx, hashHits, asm)
}

func countStream(ctx context.Context) <-chan int {
	output := make(chan int)

	go func() {
		for val := 0; ; val++ {
			select {
			case <-ctx.Done():
				close(output)
				return
			default:
				output <- val
			}
		}
	}()

	return output
}

type hashRecord struct {
	idx int
	pos int
	c   byte
}

type hashRunner func(context.Context, string, int, chan<- hashRecord)

func runHash(ctx context.Context, idxInput <-chan int, doorId string, cb hashRunner) <-chan hashRecord {
	output := make(chan hashRecord)

	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer close(output)
		for {
			select {
			case <-ctx.Done():
				return
			case idx, ok := <-idxInput:
				if !ok {
					return
				}
				go cb(ctx, doorId, idx, output)
			}
		}
	}()

	return output
}

func checkForMatchingHashPt1(ctx context.Context, doorId string, idx int, output chan<- hashRecord) {
	h := md5.New()
	fmt.Fprintf(h, "%s%d", doorId, idx)
	hashRepr := fmt.Sprintf("%x", h.Sum(nil))
	if strings.HasPrefix(hashRepr, "00000") {
		output <- hashRecord{
			idx: idx,
			c:   hashRepr[5],
		}
	}
}

var pt2InterestingHash *regexp.Regexp = regexp.MustCompile(`^00000([0-7])(.)`)

func checkForMatchingHashPt2(ctx context.Context, doorId string, idx int, output chan<- hashRecord) {
	h := md5.New()
	fmt.Fprintf(h, "%s%d", doorId, idx)
	hashRepr := fmt.Sprintf("%x", h.Sum(nil))
	if m := pt2InterestingHash.FindStringSubmatch(hashRepr); m != nil {
		pos, _ := strconv.Atoi(m[1])
		val := m[2][0]
		output <- hashRecord{
			idx: idx,
			c:   val,
			pos: pos,
		}
	}
}

type passwordAssembler interface {
	AddReport(hashRecord)
	Ready() bool
	Assemble() string
}

type pt1Assembler struct {
	buf []hashRecord
}

func (p *pt1Assembler) AddReport(h hashRecord) {
	p.buf = append(p.buf, h)
}

func (p *pt1Assembler) Ready() bool {
	return len(p.buf) == 8
}

func (p *pt1Assembler) Assemble() string {
	sort.Slice(p.buf, func(i, j int) bool {
		return p.buf[i].idx < p.buf[j].idx
	})
	var sb strings.Builder
	for _, r := range p.buf {
		sb.WriteByte(r.c)
	}

	return sb.String()
}

type pt2Assembler struct {
	buf  [8]*hashRecord
	seen int
}

func (p *pt2Assembler) AddReport(h hashRecord) {
	if p.buf[h.pos] == nil {
		p.buf[h.pos] = &h
		p.seen++
	}
}

func (p *pt2Assembler) Ready() bool {
	return p.seen == 8
}

func (p *pt2Assembler) Assemble() string {
	var sb strings.Builder
	for _, r := range p.buf {
		sb.WriteByte(r.c)
	}
	return sb.String()
}

func assemblePassword(ctx context.Context, input <-chan hashRecord, asm passwordAssembler) <-chan string {
	output := make(chan string)

	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer close(output)

		for {
			select {
			case <-ctx.Done():
				fmt.Println("failed to assemble password!ß")
				return
			case r, ok := <-input:
				if !ok {
					return
				}
				asm.AddReport(r)
				if !asm.Ready() {
					continue
				}
				output <- asm.Assemble()
				return
			}
		}
	}()

	return output
}
