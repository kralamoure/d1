package retrotyp

const (
	NPCResponseActionLeaveDialog NPCResponseAction = iota
	NPCResponseActionCreateDialog
)

type NPCResponseAction int

var NPCResponseActions = map[NPCResponseAction]string{
	NPCResponseActionLeaveDialog:  "Leave dialog",
	NPCResponseActionCreateDialog: "Create dialog",
}

func (a NPCResponseAction) Validate() error {
	_, ok := NPCResponseActions[a]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (a NPCResponseAction) String() string {
	name, ok := NPCResponseActions[a]
	if ok {
		return name
	}

	return unknownStr
}
