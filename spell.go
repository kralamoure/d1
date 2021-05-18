package retro

import "github.com/kralamoure/retro/retrotyp"

type Spell struct {
	Id          int
	Name        string
	Description string
	Levels      []retrotyp.SpellLevel
}
