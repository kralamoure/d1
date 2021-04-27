package d1

import (
	"github.com/kralamoure/d1/d1typ"
)

type Market struct {
	Id            string
	GameServerId  int
	Quantity1     int
	Quantity2     int
	Quantity3     int
	Types         []d1typ.ItemType
	Fee           float32
	MaxLevel      int
	MaxPerAccount int
	MaxHours      int
}
