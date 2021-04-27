package d1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/d1/d1typ"
)

type ItemTemplate struct {
	Id            int
	Name          string
	Description   string
	Type          d1typ.ItemType
	Enhanceable   bool
	TwoHands      bool
	Ethereal      bool
	Hidden        bool
	ItemSetId     int
	CanUse        bool
	CanTarget     bool
	Level         int
	GFX           int
	Price         int
	Weight        int
	Cursed        bool
	Conditions    string // TODO
	Effects       []d1typ.Effect
	WeaponEffects WeaponEffects
}

type WeaponEffects struct {
	CriticalHitBonus int
	APCost           int
	RangeMin         int
	RangeMax         int
	CriticalHit      int
	CriticalFailure  int
	LineOnly         bool
	LineOfSight      bool
}

// TODO: maybe should validate by iterating through effects
func (t ItemTemplate) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Type),
		validation.Field(&t.Level, validation.Min(1)),
		validation.Field(&t.Weight, validation.Min(0)),
		validation.Field(&t.Price, validation.Min(0)),
	)
}
