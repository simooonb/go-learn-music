package domain

type ScaleKind struct {
	Name      string
	Intervals []IntervalKind
}

type Scale struct {
	Root Note
	Kind ScaleKind
}

var (
	Major = ScaleKind{
		Name: "Major",
		Intervals: []IntervalKind{
			PerfectUnison, MajorSecond, MajorThird, PerfectFourth, PerfectFifth, MajorSixth, MajorSeventh,
		},
	}
	Minor = ScaleKind{
		Name: "Minor",
		Intervals: []IntervalKind{
			PerfectUnison, MajorSecond, MinorThird, PerfectFourth, PerfectFifth, MinorSixth, MinorSeventh,
		},
	}
	MajorPentatonic = ScaleKind{
		Name:      "Major pentatonic",
		Intervals: []IntervalKind{PerfectUnison, MajorSecond, MajorThird, PerfectFifth, MajorSixth},
	}
	MinorPentatonic = ScaleKind{
		Name:      "Minor pentatonic",
		Intervals: []IntervalKind{PerfectUnison, MinorThird, PerfectFourth, PerfectFifth, MinorSeventh},
	}
)

var Scales = []ScaleKind{Major, Minor, MajorPentatonic, MinorPentatonic}

func (scale Scale) AllNotes() []Note {
	scaleNotes := make([]Note, len(scale.Kind.Intervals))

	for i, interval := range scale.Kind.Intervals {
		scaleNotes[i] = scale.Root.Add(interval)
	}

	return scaleNotes
}
