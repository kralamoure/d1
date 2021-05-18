package retro

import "github.com/kralamoure/retro/retrotyp"

type NPC struct {
	Id            string
	GameServerId  int
	MapId         int
	CellId        int
	Direction     int
	TemplateId    int
	Sex           retrotyp.Sex
	GFX           int
	ScaleX        int
	ScaleY        int
	Color1        retrotyp.Color
	Color2        retrotyp.Color
	Color3        retrotyp.Color
	Accessories   string // TODO
	ExtraClip     int    // -1 for no extra clip.
	CustomArtwork int
	DialogId      int
	MarketId      string
}
