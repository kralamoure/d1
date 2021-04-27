package d1

import "github.com/kralamoure/d1/d1typ"

type NPCResponse struct {
	Id         int
	Text       string
	Action     d1typ.NPCResponseAction
	Arguments  []string
	Conditions []string // TODO
}
