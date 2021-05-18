package retro

import "github.com/kralamoure/retro/retrotyp"

type CharacterItem struct {
	Item
	Position    retrotyp.CharacterItemPosition
	CharacterId int
}
