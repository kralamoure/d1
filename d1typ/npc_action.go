package d1typ

const (
	NPCActionBuyOrSell NPCAction = iota + 1
	NPCActionExchange
	NPCActionTalk
	NPCActionDropOrPickUpPet
	NPCActionSell
	NPCActionBuy
	NPCActionResurrectPet
	NPCActionTradeMount
	NPCActionViewHouses
	NPCActionViewPaddocks
)

type NPCAction int

var NPCActions = map[NPCAction]string{
	NPCActionBuyOrSell:       "Buy/Sell",
	NPCActionExchange:        "Exchange",
	NPCActionTalk:            "Talk",
	NPCActionDropOrPickUpPet: "Drop/Pick up a pet",
	NPCActionSell:            "Sell",
	NPCActionBuy:             "Buy",
	NPCActionResurrectPet:    "ResurrectPet",
	NPCActionTradeMount:      "Trade a mount",
	NPCActionViewHouses:      "View houses",
	NPCActionViewPaddocks:    "View paddocks",
}

func (a NPCAction) Validate() error {
	_, ok := NPCActions[a]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (a NPCAction) String() string {
	name, ok := NPCActions[a]
	if ok {
		return name
	}

	return unknownStr
}
