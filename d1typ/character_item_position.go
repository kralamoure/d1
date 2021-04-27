package d1typ

const (
	CharacterItemPositionInventory   CharacterItemPosition = -1
	CharacterItemPositionAmulet      CharacterItemPosition = 0
	CharacterItemPositionWeapon      CharacterItemPosition = 1
	CharacterItemPositionRingRight   CharacterItemPosition = 2
	CharacterItemPositionBelt        CharacterItemPosition = 3
	CharacterItemPositionRingLeft    CharacterItemPosition = 4
	CharacterItemPositionBoots       CharacterItemPosition = 5
	CharacterItemPositionHat         CharacterItemPosition = 6
	CharacterItemPositionCloak       CharacterItemPosition = 7
	CharacterItemPositionPet         CharacterItemPosition = 8
	CharacterItemPositionDofus1      CharacterItemPosition = 9
	CharacterItemPositionDofus2      CharacterItemPosition = 10
	CharacterItemPositionDofus3      CharacterItemPosition = 11
	CharacterItemPositionDofus4      CharacterItemPosition = 12
	CharacterItemPositionDofus5      CharacterItemPosition = 13
	CharacterItemPositionDofus6      CharacterItemPosition = 14
	CharacterItemPositionShield      CharacterItemPosition = 15
	CharacterItemPositionDragoturkey CharacterItemPosition = 16

	CharacterItemPositionMutationItem       CharacterItemPosition = 20
	CharacterItemPositionBoostFood          CharacterItemPosition = 21
	CharacterItemPositionBlessing1          CharacterItemPosition = 22
	CharacterItemPositionBlessing2          CharacterItemPosition = 23
	CharacterItemPositionCurse1             CharacterItemPosition = 24
	CharacterItemPositionCurse2             CharacterItemPosition = 25
	CharacterItemPositionRoleplayBuff       CharacterItemPosition = 26
	CharacterItemPositionFollowingCharacter CharacterItemPosition = 27
)

type CharacterItemPosition int

var CharacterItemPositions = map[CharacterItemPosition]string{
	CharacterItemPositionInventory:   "Inventory",
	CharacterItemPositionAmulet:      "Amulet",
	CharacterItemPositionWeapon:      "Weapon",
	CharacterItemPositionRingRight:   "Right Ring",
	CharacterItemPositionBelt:        "Belt",
	CharacterItemPositionRingLeft:    "Left Ring",
	CharacterItemPositionBoots:       "Boots",
	CharacterItemPositionHat:         "Hat",
	CharacterItemPositionCloak:       "Cloak",
	CharacterItemPositionPet:         "Pet",
	CharacterItemPositionDofus1:      "Dofus 1",
	CharacterItemPositionDofus2:      "Dofus 2",
	CharacterItemPositionDofus3:      "Dofus 3",
	CharacterItemPositionDofus4:      "Dofus 4",
	CharacterItemPositionDofus5:      "Dofus 5",
	CharacterItemPositionDofus6:      "Dofus 6",
	CharacterItemPositionShield:      "Shield",
	CharacterItemPositionDragoturkey: "Dragoturkey",

	CharacterItemPositionMutationItem:       "Mutation Item",
	CharacterItemPositionBoostFood:          "Boost Food",
	CharacterItemPositionBlessing1:          "Blessing 1",
	CharacterItemPositionBlessing2:          "Blessing 2",
	CharacterItemPositionCurse1:             "Curse 1",
	CharacterItemPositionCurse2:             "Curse 2",
	CharacterItemPositionRoleplayBuff:       "Roleplay Buff",
	CharacterItemPositionFollowingCharacter: "Following Character",
}

func (t CharacterItemPosition) Validate() error {
	_, ok := CharacterItemPositions[t]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (t CharacterItemPosition) String() string {
	name, ok := CharacterItemPositions[t]
	if ok {
		return name
	}

	return unknownStr
}
