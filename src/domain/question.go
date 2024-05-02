package domain

const (
	ScaleNotesQuestionKind     = "ScaleNotesQuestionKind"
	IntervalKindQuestionKind   = "IntervalKindQuestionKind"
	NoteIntervalQuestionKind   = "NoteIntervalQuestionKind"
	ScaleIntervalsQuestionKind = "ScaleIntervalsQuestionKind"
)

var QuestionKinds = []string{
	ScaleNotesQuestionKind,
	IntervalKindQuestionKind,
	NoteIntervalQuestionKind,
	ScaleIntervalsQuestionKind,
}

type Question interface {
	Kind() string
}

type ScaleNotesQuestion struct {
	Scale Scale
}

func (question ScaleNotesQuestion) Kind() string {
	return ScaleNotesQuestionKind
}

func (question ScaleNotesQuestion) Answer() []Note {
	return question.Scale.AllNotes()
}

type IntervalKindQuestion struct {
	Interval Interval
}

func (question IntervalKindQuestion) Kind() string {
	return IntervalKindQuestionKind
}

func (question IntervalKindQuestion) Answer() IntervalKind {
	return question.Interval.IntervalKind()
}

type NoteIntervalQuestion struct {
	Root         Note
	IntervalKind IntervalKind
}

func (question NoteIntervalQuestion) Kind() string {
	return NoteIntervalQuestionKind
}

func (question NoteIntervalQuestion) Answer() Note {
	return question.Root.Add(question.IntervalKind)
}

type ScaleIntervalsQuestion struct {
	ScaleKind ScaleKind
}

func (question ScaleIntervalsQuestion) Kind() string {
	return ScaleIntervalsQuestionKind
}

func (question ScaleIntervalsQuestion) Answer() []IntervalKind {
	return question.ScaleKind.Intervals
}
