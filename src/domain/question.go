package domain

type ScaleNotesQuestion struct {
	Scale Scale
}

func (question ScaleNotesQuestion) Answer() []Note {
	return question.Scale.AllNotes()
}

type IntervalKindQuestion struct {
	interval Interval
}

func (question IntervalKindQuestion) Answer() IntervalKind {
	return question.interval.IntervalKind()
}

type NoteIntervalQuestion struct {
	root         Note
	intervalKind IntervalKind
}

func (question NoteIntervalQuestion) Answer() Note {
	return question.root.Add(question.intervalKind)
}

type ScaleIntervalsQuestion struct {
	scaleKind ScaleKind
}

func (question ScaleIntervalsQuestion) Answer() []IntervalKind {
	return question.scaleKind.Intervals
}
