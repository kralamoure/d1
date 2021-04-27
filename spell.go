package d1

import "github.com/kralamoure/d1/d1typ"

type Spell struct {
	Id          int
	Name        string
	Description string
	Levels      []d1typ.SpellLevel
}
