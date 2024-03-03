package domain

import (
	"slices"
)

type Interval struct {
	Note1 Note
	Note2 Note
}

type IntervalKind struct {
	Name       string
	Semitones  Semitone
	NoteOffset int
}

var (
	PerfectUnison     = IntervalKind{Name: "Perfect unison", Semitones: 0, NoteOffset: 0}
	MinorSecond       = IntervalKind{Name: "Minor second", Semitones: 1, NoteOffset: 1}
	MajorSecond       = IntervalKind{Name: "Major second", Semitones: 2, NoteOffset: 1}
	MinorThird        = IntervalKind{Name: "Minor third", Semitones: 3, NoteOffset: 2}
	MajorThird        = IntervalKind{Name: "Major third", Semitones: 4, NoteOffset: 2}
	PerfectFourth     = IntervalKind{Name: "Perfect fourth", Semitones: 5, NoteOffset: 3}
	PerfectFifth      = IntervalKind{Name: "Perfect fifth", Semitones: 7, NoteOffset: 4}
	MinorSixth        = IntervalKind{Name: "Minor sixth", Semitones: 8, NoteOffset: 5}
	MajorSixth        = IntervalKind{Name: "Major sixth", Semitones: 9, NoteOffset: 5}
	MinorSeventh      = IntervalKind{Name: "Minor seventh", Semitones: 10, NoteOffset: 6}
	MajorSeventh      = IntervalKind{Name: "Major seventh", Semitones: 11, NoteOffset: 6}
	PerfectOctave     = IntervalKind{Name: "Perfect octave", Semitones: 12, NoteOffset: 7}
	DiminishedSecond  = IntervalKind{Name: "Diminished second", Semitones: 0, NoteOffset: 1}
	AugmentedUnison   = IntervalKind{Name: "Augmented unison", Semitones: 1, NoteOffset: 0}
	DiminishedThird   = IntervalKind{Name: "Diminished third", Semitones: 2, NoteOffset: 2}
	AugmentedSecond   = IntervalKind{Name: "Augmented second", Semitones: 3, NoteOffset: 1}
	DiminishedFourth  = IntervalKind{Name: "Diminished fourth", Semitones: 4, NoteOffset: 3}
	AugmentedThird    = IntervalKind{Name: "Augmented third", Semitones: 5, NoteOffset: 2}
	DiminishedFifth   = IntervalKind{Name: "Diminished fifth", Semitones: 6, NoteOffset: 4}
	AugmentedFourth   = IntervalKind{Name: "Augmented fourth", Semitones: 6, NoteOffset: 3}
	DiminishedSixth   = IntervalKind{Name: "Diminished sixth", Semitones: 7, NoteOffset: 5}
	AugmentedFifth    = IntervalKind{Name: "Augmented fifth", Semitones: 8, NoteOffset: 4}
	DiminishedSeventh = IntervalKind{Name: "Diminished seventh", Semitones: 9, NoteOffset: 6}
	AugmentedSixth    = IntervalKind{Name: "Augmented sixth", Semitones: 10, NoteOffset: 5}
	DiminishedOctave  = IntervalKind{Name: "Diminished octave", Semitones: 11, NoteOffset: 7}
	AugmentedSeventh  = IntervalKind{Name: "Augmented seventh", Semitones: 12, NoteOffset: 6}
)

var PerfectMinorMajorIntervals = []IntervalKind{
	PerfectUnison,
	MinorSecond,
	MajorSecond,
	MinorThird,
	MajorThird,
	PerfectFourth,
	PerfectFifth,
	MinorSixth,
	MajorSixth,
	MinorSeventh,
	MajorSeventh,
	PerfectOctave,
}

var AugmentedDiminishedIntervals = []IntervalKind{
	DiminishedSecond,
	AugmentedUnison,
	DiminishedThird,
	AugmentedSecond,
	DiminishedFourth,
	AugmentedThird,
	DiminishedFifth,
	AugmentedFourth,
	DiminishedSixth,
	AugmentedFifth,
	DiminishedSeventh,
	AugmentedSixth,
	DiminishedOctave,
	AugmentedSeventh,
}

func (interval Interval) Semitones() Semitone {
	note1Semitones := interval.Note1.Semitones()
	note2Semitones := interval.Note2.Semitones()

	if note1Semitones > note2Semitones {
		return note1Semitones - note2Semitones
	} else {
		return note2Semitones - note1Semitones
	}
}

func (interval Interval) IntervalKind() IntervalKind {
	expectedSemitones := interval.Semitones()
	index := slices.IndexFunc(
		PerfectMinorMajorIntervals,
		func(i IntervalKind) bool { return i.Semitones == expectedSemitones },
	)

	return PerfectMinorMajorIntervals[index]
}
