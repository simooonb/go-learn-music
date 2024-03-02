package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindInterval(t *testing.T) {
	bSharp := Note{Name: B, Accidental: Sharp, Octave: 0}
	c := Note{Name: C, Accidental: NoAccidental, Octave: 0}
	d := Note{Name: D, Accidental: NoAccidental, Octave: 0}
	g := Note{Name: G, Accidental: NoAccidental, Octave: 0}

	interval1 := Interval{Note1: bSharp, Note2: c}
	interval2 := Interval{Note1: c, Note2: g}
	interval3 := Interval{Note1: d, Note2: g}

	assert.Equal(t, PerfectUnison, interval1.IntervalKind(), "B# to C should be a perfect unison")
	assert.Equal(t, PerfectFifth, interval2.IntervalKind(), "C to G should be a perfect fifth")
	assert.Equal(t, PerfectFourth, interval3.IntervalKind(), "D to G should be a perfect fourth")
}
