package d1

import "github.com/kralamoure/d1/d1typ"

type CharacterItem struct {
	Item
	Position    d1typ.CharacterItemPosition
	CharacterId int
}
