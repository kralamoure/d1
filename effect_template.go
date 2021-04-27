package d1

import "github.com/kralamoure/d1/d1typ"

type EffectTemplate struct {
	Id               int
	Description      string
	Dice             bool
	Operator         d1typ.EffectOperator
	CharacteristicId d1typ.CharacteristicId
	Element          d1typ.EffectElement
}
