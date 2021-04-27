package d1typ

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CharacterName string

func (n CharacterName) Validate() error {
	v := string(n)
	return validation.Validate(v,
		validation.Length(2, 20),
		validation.Match(regexp.MustCompile(`^[A-Z][a-z]+(-[a-zA-Z][a-z]+)?$`)),
		validation.By(func(value interface{}) error {
			if !checkMaxRepeatedContinuousRunes(v, 2) {
				return errors.New("too many repeated continuous characters")
			}
			return nil
		}),
	)
}

func checkMaxRepeatedContinuousRunes(s string, n int) bool {
	var last rune
	qty := 1
	for _, v := range s {
		if last == v {
			qty++
			if qty > n {
				return false
			}
		} else {
			qty = 1
		}
		last = v
	}
	return true
}
