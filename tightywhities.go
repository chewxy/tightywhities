package tightywhities

import (
	"fmt"
	"io"
	"strings"
)

// Mode indicates the drawing mode of
type Mode int

const (
	// Connective tries to daraw ASCII lines
	Connective Mode = iota

	// Point draws dots instead. Inspired directly from TermUI
	Point
)

// Plot is an interface that represents a buffer for the plot
type Plot interface {
	// SetAt sets the value of point(x,y) to the char
	SetAt(x, y int, char rune)

	// Height returns the height of the plot (for scaling calculations)
	Height() int

	// Width returns the width of the plot (for scaling calculations)
	Width() int

	// WriterTo has to be implemented for flushing purposes
	io.WriterTo
}

type state struct {
	buf     [][]rune
	w, h, o int // width, height, x-offset
}

func newState(width, height int) *state {
	buf := make([][]rune, height+1)
	for i := range buf {
		buf[i] = make([]rune, width+1)
		for j := range buf[i] {
			buf[i][j] = ' '
		}
	}
	return &state{
		buf: buf,
		w:   width,
		h:   height,
	}
}

func (s *state) SetAt(x, y int, char rune) {
	s.buf[s.h-y][x+s.o] = char
}

func (s *state) WriteTo(w io.Writer) (n int64, err error) {
	for i := range s.buf {
		var nn int
		if nn, err = fmt.Fprintf(w, "│%s\n", strings.TrimRight(string(s.buf[i]), " ")); err != nil {
			return int64(nn), err
		}
		n += int64(nn)
	}
	return
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
