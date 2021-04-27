package d1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/d1/d1typ"
)

type ItemSet struct {
	Id    int
	Name  string
	Bonus [][]d1typ.Effect
}

// TODO: maybe should validate by iterating through bonus
func (s ItemSet) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Bonus),
	)
}
