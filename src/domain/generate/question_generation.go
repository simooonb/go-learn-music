package generate

import (
	"go-learn-music/src/domain"
	"math/rand/v2"
)

const MAX_OCTAVE int = 4

func RandomQuestion() domain.Question {
	var question domain.Question

	switch RandomQuestionKind() {
	case domain.ScaleNotesQuestionKind:
		question = RandomScaleNotesQuestion()
	case domain.IntervalKindQuestionKind:
		question = RandomIntervalKindQuestion()
	case domain.NoteIntervalQuestionKind:
		question = RandomNoteIntervalQuestion()
	case domain.ScaleIntervalsQuestionKind:
		question = RandomScaleNotesQuestion()
	}

	return question
}

func RandomQuestionKind() string {
	return domain.QuestionKinds[rand.IntN(len(domain.QuestionKinds))]
}

func RandomScaleNotesQuestion() domain.ScaleNotesQuestion {
	return domain.ScaleNotesQuestion{Scale: RandomScale()}
}

func RandomIntervalKindQuestion() domain.IntervalKindQuestion {
	return domain.IntervalKindQuestion{
		Interval: domain.Interval{Note1: RandomNote(0), Note2: RandomNote(0)},
	}
}

func RandomNoteIntervalQuestion() domain.NoteIntervalQuestion {
	return domain.NoteIntervalQuestion{
		Root:         RandomNote(0),
		IntervalKind: RandomIntervalKind(),
	}
}

func RandomScaleIntervalsQuestion() domain.ScaleIntervalsQuestion {
	return domain.ScaleIntervalsQuestion{
		ScaleKind: RandomScaleKind(),
	}
}

func RandomNote(octave int) domain.Note {
	return domain.Note{
		Name:       RandomNoteName(),
		Accidental: RandomAccidental(),
		Octave:     octave,
	}
}

func RandomScale() domain.Scale {
	return domain.Scale{Root: RandomNote(0), Kind: RandomScaleKind()}
}

func RandomScaleKind() domain.ScaleKind {
	return domain.Scales[rand.IntN(len(domain.Scales))]
}

func RandomIntervalKind() domain.IntervalKind {
	return domain.PerfectMinorMajorIntervals[rand.IntN(len(domain.PerfectMinorMajorIntervals))]
}

func RandomNoteName() domain.NoteName {
	return domain.NoteNames[rand.IntN(len(domain.NoteNames))]
}

func RandomAccidental() domain.Accidental {
	return domain.Accidentals[rand.IntN(len(domain.Accidentals))]
}

func RandomOctave() int {
	return rand.IntN(MAX_OCTAVE) + 1
}
