package retro

import "github.com/kralamoure/retro/retrotyp"

type Class struct {
	Id               retrotyp.ClassId
	Name             string
	Label            string
	ShortDescription string
	Description      string
	Spells           []int
	BoostCosts       ClassBoostCosts
}

type ClassBoostCosts struct {
	Vitality     []ClassBoostCost
	Wisdom       []ClassBoostCost
	Strength     []ClassBoostCost
	Intelligence []ClassBoostCost
	Chance       []ClassBoostCost
	Agility      []ClassBoostCost
}

type ClassBoostCost struct {
	Quantity int
	Cost     int
	Bonus    int
}
