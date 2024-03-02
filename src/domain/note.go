package domain

type NoteName struct {
	Name      string
	Semitones Semitone
	Order     int
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

const (
	NoAccidental Semitone = 0
	Sharp        Semitone = 1
	SharpSharp   Semitone = 2
	Flat         Semitone = -1
	FlatFlat     Semitone = -2
)

type Note struct {
	Name       NoteName
	Accidental Semitone
	Octave     int
}

var NoteNames = []NoteName{A, B, C, D, E, F, G}
var Accidentals = []Semitone{NoAccidental, Sharp, SharpSharp, Flat, FlatFlat}

var NoteNamesByOrder = map[int]NoteName{
	A.Order: A,
	B.Order: B,
	C.Order: C,
	D.Order: D,
	E.Order: E,
	F.Order: F,
	G.Order: G,
}

func (note Note) Semitones() Semitone {
	return note.Name.Semitones + note.Accidental + Semitone(12*note.Octave)
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

	var accidental Semitone
	semitonesLeft := semitones - name.Semitones

	switch semitonesLeft {
	case Sharp:
		accidental = Sharp
	case Flat:
		accidental = Flat
	case SharpSharp:
		accidental = SharpSharp
	case FlatFlat:
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
