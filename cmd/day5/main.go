package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"sort"
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
}

func part1(input string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	index := countStream(ctx)
	hashHits := runHash(ctx, index, input)
	return <-assemblePassword(ctx, hashHits)
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
	c   byte
}

func runHash(ctx context.Context, idxInput <-chan int, doorId string) <-chan hashRecord {
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
				go checkForMatchingHash(ctx, doorId, idx, output)
			}
		}
	}()

	return output
}

func checkForMatchingHash(ctx context.Context, doorId string, idx int, output chan<- hashRecord) {
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

func assemblePassword(ctx context.Context, input <-chan hashRecord) <-chan string {
	output := make(chan string)

	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer close(output)
		buf := make([]hashRecord, 0)

		for {
			select {
			case <-ctx.Done():
				fmt.Println("failed to assemble password!ß")
				return
			case r, ok := <-input:
				if !ok {
					return
				}
				buf = append(buf, r)
				if len(buf) < 8 {
					continue
				}

				sort.Slice(buf, func(i, j int) bool {
					return buf[i].idx < buf[j].idx
				})

				var sb strings.Builder
				for _, r := range buf {
					sb.WriteByte(r.c)
				}

				output <- sb.String()
				return
			}
		}
	}()

	return output
}
