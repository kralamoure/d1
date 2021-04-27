package d1

import "github.com/kralamoure/d1/d1typ"

type Class struct {
	Id               d1typ.ClassId
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
