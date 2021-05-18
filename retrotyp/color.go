package retrotyp

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Color string

func (c Color) Validate() error {
	v := string(c)

	return validation.Validate(v,
		validation.Match(regexp.MustCompile(`^[0-9a-fA-F]{6}$`)),
		is.LowerCase,
	)
}
