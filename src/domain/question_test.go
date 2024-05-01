package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScaleNotesQuestion(t *testing.T) {
	c := Note{Name: C, Accidental: NoAccidental, Octave: 0}
	cMajor := Scale{Root: c, Kind: Major}
	cMajorQuestion := ScaleNotesQuestion{Scale: cMajor}

	assert.Equal(t, cMajorQuestion.Answer(), cMajor.AllNotes(), "Question should have scale notes as answer")
}

func TestIntervalKindQuestion(t *testing.T) {
	b := Note{Name: B, Accidental: NoAccidental, Octave: 0}
	c := Note{Name: C, Accidental: NoAccidental, Octave: 0}
	bcInterval := Interval{Note1: b, Note2: c}
	question := IntervalKindQuestion{interval: bcInterval}

	assert.Equal(t, question.Answer(), bcInterval.IntervalKind(), "Question should have interval kind as answer")
}

func TestNoteIntervalQuestion(t *testing.T) {
	dSharp := Note{Name: D, Accidental: Sharp, Octave: 1}
	minorSixth := MinorSixth
	question := NoteIntervalQuestion{root: dSharp, intervalKind: minorSixth}

	assert.Equal(t, question.Answer(), dSharp.Add(minorSixth), "Question should have proper note as answer")
}

func TestScaleIntervalsQuestion(t *testing.T) {
	kind := MajorPentatonic
	question := ScaleIntervalsQuestion{scaleKind: kind}

	assert.Equal(t, question.Answer(), kind.Intervals, "Question should have proper intervals as answer")
}
