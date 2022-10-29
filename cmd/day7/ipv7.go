package main

import (
	"context"
	"strings"
	"sync"
)

type IPV7Address struct {
	raw   string
	outer []string
	hnet  []string
}

func IPV7AddressFromString(input string) IPV7Address {
	i := IPV7Address{
		raw:   input,
		outer: make([]string, 0),
		hnet:  make([]string, 0),
	}

	var sb strings.Builder
	for _, c := range input {
		switch c {
		case '[':
			i.outer = append(i.outer, sb.String())
			sb.Reset()
		case ']':
			i.hnet = append(i.hnet, sb.String())
			sb.Reset()
		default:
			sb.WriteRune(c)
		}
	}
	// we assume a well formed address with no open brackets.
	// as such, the contents of the buffer are the last outer sequence.
	if sb.Len() > 0 {
		i.outer = append(i.outer, sb.String())
	}
	return i
}

func (i IPV7Address) SupportsTLS() bool {
	for _, s := range i.hnet {
		if abbaDetected(s) {
			return false
		}
	}

	for _, s := range i.outer {
		if abbaDetected(s) {
			return true
		}
	}

	return false
}

func (i IPV7Address) SupportsSSL() bool {
	babStream := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(i.outer))
	for _, s := range i.outer {
		go func(s string) {
			defer wg.Done()
			matchingBab(ctx, s, babStream)
		}(s)
	}

	go func() {
		wg.Wait()
		close(babStream)
	}()

	for {
		select {
		case <-ctx.Done():
			return false
		case b, ok := <-babStream:
			if !ok {
				return false
			}
			for _, s := range i.hnet {
				if strings.Contains(s, b) {
					return true
				}
			}
		}
	}
}

func abbaDetected(input string) bool {
	if len(input) < 4 {
		return false
	}

	w := make([]byte, 3)
	for i := 0; i < 3; i++ {
		w[i] = input[i]
	}
	for i := 3; i < len(input); i++ {
		w = append(w, input[i])
		if w[0] == w[3] && w[0] != w[1] && w[1] == w[2] {
			return true
		}
		w = w[1:]
	}
	return false
}

func matchingBab(ctx context.Context, input string, output chan<- string) {
	if len(input) < 3 {
		return
	}

	w := []byte(input[:2])
	for i := 2; i < len(input); i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		w = append(w, input[i])
		if w[0] == w[2] && w[0] != w[1] {
			output <- string([]byte{w[1], w[0], w[1]})
		}
		w = w[1:]
	}

}
