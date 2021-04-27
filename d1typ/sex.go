package d1typ

const (
	SexMale Sex = iota
	SexFemale
)

type Sex int

var Sexes = map[Sex]string{
	SexMale:   "Male",
	SexFemale: "Female",
}

func (g Sex) Validate() error {
	_, ok := Sexes[g]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (g Sex) String() string {
	name, ok := Sexes[g]
	if ok {
		return name
	}

	return unknownStr
}
