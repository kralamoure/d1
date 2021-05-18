package retro

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/retro/retrotyp"
)

type ItemSet struct {
	Id    int
	Name  string
	Bonus [][]retrotyp.Effect
}

// TODO: maybe should validate by iterating through bonus
func (s ItemSet) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Bonus),
	)
}
