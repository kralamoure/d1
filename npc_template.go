package d1

import "github.com/kralamoure/d1/d1typ"

type NPCTemplate struct {
	Id      int
	Name    string
	Actions []d1typ.NPCAction
}
