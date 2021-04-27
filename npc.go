package d1

import "github.com/kralamoure/d1/d1typ"

type NPC struct {
	Id            string
	GameServerId  int
	MapId         int
	CellId        int
	Direction     int
	TemplateId    int
	Sex           d1typ.Sex
	GFX           int
	ScaleX        int
	ScaleY        int
	Color1        d1typ.Color
	Color2        d1typ.Color
	Color3        d1typ.Color
	Accessories   string // TODO
	ExtraClip     int    // -1 for no extra clip.
	CustomArtwork int
	DialogId      int
	MarketId      string
}
