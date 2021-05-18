package retrotyp

const (
	AlignmentNeutral Alignment = iota
	AlignmentBontarian
	AlignmentBrakmarian
	AlignmentMercenary
)

type Alignment int

var Alignments = map[Alignment]string{
	AlignmentNeutral:    "Neutral",
	AlignmentBontarian:  "Bontarian",
	AlignmentBrakmarian: "Brakmarian",
	AlignmentMercenary:  "Mercenary",
}

func (a Alignment) Validate() error {
	_, ok := Alignments[a]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (a Alignment) String() string {
	name, ok := Alignments[a]
	if ok {
		return name
	}

	return unknownStr
}
