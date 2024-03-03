package domain

import (
	"fmt"
	"strconv"
)

type NoteName struct {
	Name      string
	Semitones Semitone
	Order     int
}

type Accidental struct {
	Symbol    string
	Semitones Semitone
}

var (
	A = NoteName{Name: "A", Semitones: 1, Order: 1}
	B = NoteName{Name: "B", Semitones: 3, Order: 2}
	C = NoteName{Name: "C", Semitones: 4, Order: 3}
	D = NoteName{Name: "D", Semitones: 6, Order: 4}
	E = NoteName{Name: "E", Semitones: 8, Order: 5}
	F = NoteName{Name: "F", Semitones: 9, Order: 6}
	G = NoteName{Name: "G", Semitones: 11, Order: 7}
)

var (
	NoAccidental = Accidental{Symbol: "", Semitones: 0}
	Sharp        = Accidental{Symbol: "#", Semitones: 1}
	SharpSharp   = Accidental{Symbol: "##", Semitones: 2}
	Flat         = Accidental{Symbol: "b", Semitones: -1}
	FlatFlat     = Accidental{Symbol: "bb", Semitones: -2}
)

type Note struct {
	Name       NoteName
	Accidental Accidental
	Octave     int
}

var NoteNames = []NoteName{A, B, C, D, E, F, G}
var Accidentals = []Accidental{NoAccidental, Sharp, SharpSharp, Flat, FlatFlat}

var NoteNamesByOrder = map[int]NoteName{
	A.Order: A,
	B.Order: B,
	C.Order: C,
	D.Order: D,
	E.Order: E,
	F.Order: F,
	G.Order: G,
}

var NoteNamesByString = map[string]NoteName{
	A.Name: A,
	B.Name: B,
	C.Name: C,
	D.Name: D,
	E.Name: E,
	F.Name: F,
	G.Name: G,
}

var AccidentalByString = map[string]Accidental{
	NoAccidental.Symbol: NoAccidental,
	Sharp.Symbol:        Sharp,
	Flat.Symbol:         Flat,
	SharpSharp.Symbol:   SharpSharp,
	FlatFlat.Symbol:     FlatFlat,
}

func (note Note) String() string {
	return fmt.Sprintf("%s%s%d", note.Name.Name, note.Accidental.Symbol, note.Octave)
}

func DecodeNote(str string) Note {
	octave, _ := strconv.Atoi(str[len(str)-1:])

	var accidental Accidental

	if len(str)-1 >= 1 {
		accidental = AccidentalByString[str[1:len(str)-1]]
	} else {
		accidental = NoAccidental
	}

	return Note{
		Name:       NoteNamesByString[str[0:1]],
		Accidental: accidental,
		Octave:     octave,
	}
}

func (note Note) Semitones() Semitone {
	return note.Name.Semitones + note.Accidental.Semitones + Semitone(12*note.Octave)
}

func (note Note) Add(interval IntervalKind) Note {
	newNoteOrder := note.Name.Order + interval.NoteOffset
	if newNoteOrder > G.Order {
		newNoteOrder = newNoteOrder % G.Order
	}

	name := NoteNamesByOrder[newNoteOrder]
	semitones := note.Semitones() + interval.Semitones
	newOctave := int(semitones / PerfectOctave.Semitones)
	octaveOffset := newOctave - note.Octave

	if octaveOffset >= 1 {
		semitones = semitones % PerfectOctave.Semitones
	}

	var accidental Accidental
	semitonesLeft := semitones - name.Semitones

	switch semitonesLeft {
	case Sharp.Semitones:
		accidental = Sharp
	case Flat.Semitones:
		accidental = Flat
	case SharpSharp.Semitones:
		accidental = SharpSharp
	case FlatFlat.Semitones:
		accidental = FlatFlat
	default:
		accidental = NoAccidental
	}

	return Note{
		Name:       name,
		Accidental: accidental,
		Octave:     newOctave,
	}
}
