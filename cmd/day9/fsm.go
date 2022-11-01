package main

import (
	"errors"
	"io"
)

type CharCtr struct {
	subseqLoad *SubseqLoad
	ctr        *int
}

type NumParser struct {
	val int
}

type SubseqLoad struct {
	*NumParser
	subseqCtr *SubseqCtr
	multLoad  *MultLoad
}
type MultLoad struct {
	*NumParser
	subseqCtr *SubseqCtr
}
type SubseqCtr struct {
	seqCt     int
	increment int
	ctr       *int
	charCtr   *CharCtr
}

type FsmState interface {
	Update(byte) FsmState
}

func (n *NumParser) AddDigit(b byte) {
	offset := b - '0'
	if offset > 9 {
		panic(errors.New("parsing non-numberic character"))
	}

	n.val = n.val * 10
	n.val = n.val + int(offset)
}

func (n *NumParser) GetValue() int {
	return n.val
}

func (n *NumParser) Reset() {
	n.val = 0
}

func (c *CharCtr) Update(b byte) FsmState {
	if b == '(' {
		return c.subseqLoad
	}
	(*c.ctr)++
	return c
}

func (s *SubseqCtr) Update(b byte) FsmState {
	(*s.ctr) += s.increment
	s.seqCt--
	if s.seqCt == 0 {
		return s.charCtr
	}
	return s
}

func (s *SubseqLoad) Update(b byte) FsmState {
	if b == 'x' {
		s.subseqCtr.seqCt = s.GetValue()
		s.Reset()
		return s.multLoad
	}
	s.AddDigit(b)
	return s
}

func (m *MultLoad) Update(b byte) FsmState {
	if b == ')' {
		m.subseqCtr.increment = m.GetValue()
		m.Reset()
		return m.subseqCtr
	}
	m.AddDigit(b)
	return m
}

type FsmParser struct {
	ctr        int
	charCtr    *CharCtr
	subseqLoad *SubseqLoad
	multLoad   *MultLoad
	subseqCtr  *SubseqCtr
	state      FsmState
}

func NewFsmParser() *FsmParser {
	f := &FsmParser{
		charCtr: &CharCtr{},
		subseqLoad: &SubseqLoad{
			NumParser: &NumParser{},
		},
		multLoad: &MultLoad{
			NumParser: &NumParser{},
		},
		subseqCtr: &SubseqCtr{},
	}

	f.charCtr.subseqLoad = f.subseqLoad
	f.subseqLoad.multLoad = f.multLoad
	f.subseqLoad.subseqCtr = f.subseqCtr
	f.multLoad.subseqCtr = f.subseqCtr
	f.subseqCtr.charCtr = f.charCtr

	f.state = f.charCtr
	f.charCtr.ctr = &f.ctr
	f.subseqCtr.ctr = &f.ctr

	return f
}

func (f *FsmParser) Parse(b io.ByteReader) {
	f.ctr = 0

	for n, err := b.ReadByte(); err == nil; n, err = b.ReadByte() {
		if n == '\n' {
			continue
		}
		f.state = f.state.Update(n)
	}
}

func (f *FsmParser) CharCount() int {
	return f.ctr
}
