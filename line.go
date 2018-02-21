package tightywhities

import (
	"io"
	"math"
	"strconv"
)

// Line represents a line chart. It tries as best to draw a connective chart using an "aliased" (as in antialiasing) method
type Line struct {
	Xs       []string // nil OK
	Ys       []float64
	min      float64
	max      float64
	interval float64

	Mode Mode
}

// NewLine creates a new *Line
func NewLine(Xs []string, Ys []float64) (cinema *Line) {
	// validation
	if len(Xs) == 0 {
		Xs = make([]string, 0, len(Ys))
		for i := range Ys {
			Xs = append(Xs, strconv.FormatInt(int64(i), 10))
		}
	}
	if len(Xs) != len(Ys) {
		panic("Expected Xs and Ys to have the same length")
	}

	min := math.Inf(1)
	max := math.Inf(-1)
	for _, v := range Ys {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	interval := math.Abs(max - min)
	return &Line{
		Xs:       Xs,
		Ys:       Ys,
		min:      min,
		max:      max,
		interval: interval,
	}
}

// Plot draws the plot according to the desired W and H
//
// BUG(anyone): currently the plot doesn't scale for X. I'm too sleepy at this point to figure out scaling
func (l *Line) Plot(pltW, pltH int, w io.Writer) (n int64, err error) {
	var spacing = 2
	// scaleX := int(math.Ceil(float64(pltW) / float64(len(l.Ys))))
	scaleY := float64(pltH) / l.interval
	min := math.Floor(l.min * scaleY)
	max := math.Ceil(l.max * scaleY)

	height := int(math.Abs(max - min))
	// width := pltW + spacing // spacing
	width := len(l.Ys) + spacing

	s := newState(width+1, height)
	s.o = spacing
	fstIdx := int(l.Ys[0]*scaleY - min)
	s.SetAt(-spacing, height-fstIdx, '+')
	switch l.Mode {
	case Connective:
		l.plotLine(s, scaleY, min)
	case Point:
		l.plotPoint(s, scaleY, min)
	}

	return s.WriteTo(w)
}

func (l *Line) plotLine(s *state, scaleY, min float64) {
	for i, y := range l.Ys[:len(l.Ys)-1] {
		this := int(y*scaleY - min)
		next := int((l.Ys[i+1] * scaleY) - min)
		x := i
		if this == next {
			s.SetAt(x, this, '─')
		} else {
			if this > next {
				s.SetAt(x, this, '╮')
				s.SetAt(x, next, '╰')
			} else {
				s.SetAt(x, this, '╯')
				s.SetAt(x, next, '╭')
			}

			start := minInt(this, next)
			end := maxInt(this, next)
			for j := start + 1; j < end; j++ {
				s.SetAt(x, j, '│')
			}
		}
	}
}

func (l *Line) plotPoint(s *state, scaleY, min float64) {
	for i, y := range l.Ys[:len(l.Ys)-1] {
		this := int(y*scaleY - min)
		next := int((l.Ys[i+1] * scaleY) - min)
		x := i
		if this == next {
			s.SetAt(x, this, '⠤')
		} else {
			if this > next {
				s.SetAt(x, this, '⠢')
				s.SetAt(x, next, '⠑')
			} else {
				s.SetAt(x, this, '⠊')
				s.SetAt(x, next, '⡠')
			}

			start := minInt(this, next)
			end := maxInt(this, next)
			for j := start + 1; j < end; j++ {
				s.SetAt(x, j, '⡇')
			}
		}
	}
}
