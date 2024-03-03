package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotesSemitones(t *testing.T) {
	bSharp := Note{Name: B, Accidental: Sharp, Octave: 0}
	c := Note{Name: C, Accidental: NoAccidental, Octave: 0}

	assert.Equal(t, Semitone(4), bSharp.Semitones(), "B# should have 4 Semitones")
	assert.Equal(t, c.Semitones(), bSharp.Semitones(), "B# and C should have the same number of Semitones")
}

func TestAddNotes(t *testing.T) {
	b := Note{Name: B, Accidental: NoAccidental, Octave: 0}
	c := Note{Name: C, Accidental: NoAccidental, Octave: 0}
	cSharp := Note{Name: C, Accidental: Sharp, Octave: 0}
	eFlat := Note{Name: E, Accidental: Flat, Octave: 0}
	fFlat := Note{Name: F, Accidental: Flat, Octave: 0}
	f := Note{Name: F, Accidental: NoAccidental, Octave: 0}
	g := Note{Name: G, Accidental: NoAccidental, Octave: 0}
	aFlat := Note{Name: A, Accidental: Flat, Octave: 1}
	bFlat := Note{Name: B, Accidental: Flat, Octave: 1}
	bFlatUp := Note{Name: B, Accidental: Flat, Octave: 2}

	assert.Equal(t, cSharp, b.Add(MajorSecond), "B + M2 should be C#")
	assert.Equal(t, c, b.Add(MinorSecond), "B + m2 should be C")
	assert.Equal(t, fFlat, eFlat.Add(MinorSecond), "Eb + m2 should be Fb")
	assert.Equal(t, f, eFlat.Add(MajorSecond), "Eb + m2 should be F")
	assert.Equal(t, aFlat, g.Add(MinorSecond), "Gb + m2 should be Ab an octave up")
	assert.Equal(t, aFlat, eFlat.Add(PerfectFourth), "Eb + P4 should be Ab an octave up")
	assert.Equal(t, bFlat, eFlat.Add(PerfectFifth), "Eb + P5 should be Bb an octave up")
	assert.Equal(t, bFlatUp, bFlat.Add(PerfectOctave), "Bb + P12 should be Bb an octave up")
}

func TestDecodeNote(t *testing.T) {
	dSharpStr := "D#1"
	gFlatFlatStr := "Gbb3"
	aStr := "A4"
	bSharpSharpStr := "B##4"
	cFlatStr := "Cb2"

	dSharp := Note{Name: D, Accidental: Sharp, Octave: 1}
	gFlatFlat := Note{Name: G, Accidental: FlatFlat, Octave: 3}
	a := Note{Name: A, Accidental: NoAccidental, Octave: 4}
	bSharpSharp := Note{Name: B, Accidental: SharpSharp, Octave: 4}
	cFlat := Note{Name: C, Accidental: Flat, Octave: 2}

	// todo: add failing cases
	assert.Equal(t, dSharp, DecodeNote(dSharpStr), "D#1 should be decoded")
	assert.Equal(t, gFlatFlat, DecodeNote(gFlatFlatStr), "Gbb3 should be decoded")
	assert.Equal(t, a, DecodeNote(aStr), "A4 should be decoded")
	assert.Equal(t, bSharpSharp, DecodeNote(bSharpSharpStr), "4 should be decoded")
	assert.Equal(t, cFlat, DecodeNote(cFlatStr), "Cb2 should be decoded")
}
