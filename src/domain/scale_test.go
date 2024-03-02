package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScale_AllNotes(t *testing.T) {
	c := Note{Name: C, Accidental: NoAccidental, Octave: 0}
	e := Note{Name: E, Accidental: NoAccidental, Octave: 0}
	gFlat := Note{Name: G, Accidental: Flat, Octave: 0}

	cMajor := Scale{Root: c, Kind: Major}
	eMinorPentatonic := Scale{Root: e, Kind: MinorPentatonic}
	gFlatMinor := Scale{Root: gFlat, Kind: Minor}

	expectedCMajorNotes := []Note{
		{Name: C, Accidental: NoAccidental, Octave: 0},
		{Name: D, Accidental: NoAccidental, Octave: 0},
		{Name: E, Accidental: NoAccidental, Octave: 0},
		{Name: F, Accidental: NoAccidental, Octave: 0},
		{Name: G, Accidental: NoAccidental, Octave: 0},
		{Name: A, Accidental: NoAccidental, Octave: 1},
		{Name: B, Accidental: NoAccidental, Octave: 1},
	}

	expectedEMinorPentatonicNotes := []Note{
		{Name: E, Accidental: NoAccidental, Octave: 0},
		{Name: G, Accidental: NoAccidental, Octave: 0},
		{Name: A, Accidental: NoAccidental, Octave: 1},
		{Name: B, Accidental: NoAccidental, Octave: 1},
		{Name: D, Accidental: NoAccidental, Octave: 1},
	}

	expectedGFlatMinorNotes := []Note{
		{Name: G, Accidental: Flat, Octave: 0},
		{Name: A, Accidental: Flat, Octave: 1},
		{Name: B, Accidental: FlatFlat, Octave: 1},
		{Name: C, Accidental: Flat, Octave: 1},
		{Name: D, Accidental: Flat, Octave: 1},
		{Name: E, Accidental: FlatFlat, Octave: 1},
		{Name: F, Accidental: Flat, Octave: 1},
	}

	assert.Equal(t, expectedCMajorNotes, cMajor.AllNotes(), "Cmaj should have: C, D, E, F, G, A, B")
	assert.Equal(t, expectedEMinorPentatonicNotes, eMinorPentatonic.AllNotes(), "EminP should have: E, G, A, B, D")
	assert.Equal(t, expectedGFlatMinorNotes, gFlatMinor.AllNotes(), "Gmin should have: Gb, Ab, Bbb, Cb, Db, Ebb, Fb")
}
