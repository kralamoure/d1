package retro

import "github.com/kralamoure/retro/retrotyp"

type NPCTemplate struct {
	Id      int
	Name    string
	Actions []retrotyp.NPCAction
}
