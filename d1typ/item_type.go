package d1typ

const (
	ItemTypeAmulet                           ItemType = 1
	ItemTypeBow                              ItemType = 2
	ItemTypeWand                             ItemType = 3
	ItemTypeStaff                            ItemType = 4
	ItemTypeDagger                           ItemType = 5
	ItemTypeSword                            ItemType = 6
	ItemTypeHammer                           ItemType = 7
	ItemTypeShovel                           ItemType = 8
	ItemTypeRing                             ItemType = 9
	ItemTypeBelt                             ItemType = 10
	ItemTypeBoots                            ItemType = 11
	ItemTypePotion                           ItemType = 12
	ItemTypeExperienceScroll                 ItemType = 13
	ItemTypeGiftItem                         ItemType = 14
	ItemTypeResource                         ItemType = 15
	ItemTypeHat                              ItemType = 16
	ItemTypeCloak                            ItemType = 17
	ItemTypePet                              ItemType = 18
	ItemTypeAxe                              ItemType = 19
	ItemTypeTool                             ItemType = 20
	ItemTypePickaxe                          ItemType = 21
	ItemTypeScythe                           ItemType = 22
	ItemTypeDofus                            ItemType = 23
	ItemTypeQuestItem                        ItemType = 24
	ItemTypeDocument                         ItemType = 25
	ItemTypeSmithmagicPotion                 ItemType = 26
	ItemTypeMutationItem                     ItemType = 27
	ItemTypeBoostFood                        ItemType = 28
	ItemTypeBlessing                         ItemType = 29
	ItemTypeCurse                            ItemType = 30
	ItemTypeRoleplayBuff                     ItemType = 31
	ItemTypeFollowingCharacter               ItemType = 32
	ItemTypeBread                            ItemType = 33
	ItemTypeCereal                           ItemType = 34
	ItemTypeFlower                           ItemType = 35
	ItemTypePlant                            ItemType = 36
	ItemTypeBeer                             ItemType = 37
	ItemTypeWood                             ItemType = 38
	ItemTypeOre                              ItemType = 39
	ItemTypeAlloy                            ItemType = 40
	ItemTypeFish                             ItemType = 41
	ItemTypeCandy                            ItemType = 42
	ItemTypeForgetfulnessPotion              ItemType = 43
	ItemTypeUnlearningPotionForProfession    ItemType = 44
	ItemTypeSkillLossPotion                  ItemType = 45
	ItemTypeFruit                            ItemType = 46
	ItemTypeBone                             ItemType = 47
	ItemTypePowder                           ItemType = 48
	ItemTypeEdibleFish                       ItemType = 49
	ItemTypePreciousStone                    ItemType = 50
	ItemTypeStone                            ItemType = 51
	ItemTypeFlour                            ItemType = 52
	ItemTypeFeather                          ItemType = 53
	ItemTypeHair                             ItemType = 54
	ItemTypeFabric                           ItemType = 55
	ItemTypeLeather                          ItemType = 56
	ItemTypeWool                             ItemType = 57
	ItemTypeSeed                             ItemType = 58
	ItemTypeSkin                             ItemType = 59
	ItemTypeOil                              ItemType = 60
	ItemTypeStuffedAnimal                    ItemType = 61
	ItemTypeGuttedFish                       ItemType = 62
	ItemTypeMeat                             ItemType = 63
	ItemTypePreservedMeat                    ItemType = 64
	ItemTypeTail                             ItemType = 65
	ItemTypeMetaria                          ItemType = 66
	ItemTypeVegetable                        ItemType = 68
	ItemTypeEdibleMeat                       ItemType = 69
	ItemTypeDyeing                           ItemType = 70
	ItemTypeAlchemyEquipment                 ItemType = 71
	ItemTypePetEgg                           ItemType = 72
	ItemTypeSkill                            ItemType = 73
	ItemTypeFairywork                        ItemType = 74
	ItemTypeSpellScroll                      ItemType = 75
	ItemTypeCharacteristicScroll             ItemType = 76
	ItemTypeBowKennelCertificate             ItemType = 77
	ItemTypeSmithmagicRune                   ItemType = 78
	ItemTypeDrink                            ItemType = 79
	ItemTypeMissionItem                      ItemType = 80
	ItemTypeBackpack                         ItemType = 81
	ItemTypeShield                           ItemType = 82
	ItemTypeSoulStone                        ItemType = 83
	ItemTypeKeys                             ItemType = 84
	ItemTypeFullSoulStone                    ItemType = 85
	ItemTypeForgetfulnessPotionForCollectors ItemType = 86
	ItemTypeSeekerScroll                     ItemType = 87
	ItemTypeMagicStone                       ItemType = 88
	ItemTypeGifts                            ItemType = 89
	ItemTypePetGhost                         ItemType = 90
	ItemTypeDragoturkey                      ItemType = 91
	ItemTypeGobball                          ItemType = 92
	ItemTypeBreedingItem                     ItemType = 93
	ItemTypeUsableItem                       ItemType = 94
	ItemTypePlank                            ItemType = 95
	ItemTypeBark                             ItemType = 96
	ItemTypeMountCertificate                 ItemType = 97
	ItemTypeRoot                             ItemType = 98
	ItemTypeCapturingNet                     ItemType = 99
	ItemTypeBagOfResources                   ItemType = 100
	ItemTypeCrossbow                         ItemType = 102
	ItemTypeLeg                              ItemType = 103
	ItemTypeWing                             ItemType = 104
	ItemTypeEgg                              ItemType = 105
	ItemTypeEar                              ItemType = 106
	ItemTypeCarapace                         ItemType = 107
	ItemTypeBud                              ItemType = 108
	ItemTypeEye                              ItemType = 109
	ItemTypeJelly                            ItemType = 110
	ItemTypeShell                            ItemType = 111
	ItemTypePrism                            ItemType = 112
	ItemTypeLivingObjects                    ItemType = 113
	ItemTypeMagicWeapon                      ItemType = 114
	ItemTypeFragmentOfShushuSoul             ItemType = 115
	ItemTypePetPotion                        ItemType = 116
)

type ItemType int

var ItemTypes = map[ItemType]string{
	ItemTypeAmulet:                           "Amulet",
	ItemTypeBow:                              "Bow",
	ItemTypeWand:                             "Wand",
	ItemTypeStaff:                            "Staff",
	ItemTypeDagger:                           "Dagger",
	ItemTypeSword:                            "Sword",
	ItemTypeHammer:                           "Hammer",
	ItemTypeShovel:                           "Shovel",
	ItemTypeRing:                             "Ring",
	ItemTypeBelt:                             "Belt",
	ItemTypeBoots:                            "Boots",
	ItemTypePotion:                           "Potion",
	ItemTypeExperienceScroll:                 "Experience Scroll",
	ItemTypeGiftItem:                         "Gift Item",
	ItemTypeResource:                         "Resource",
	ItemTypeHat:                              "Hat",
	ItemTypeCloak:                            "Cloak",
	ItemTypePet:                              "Pet",
	ItemTypeAxe:                              "Axe",
	ItemTypeTool:                             "Tool",
	ItemTypePickaxe:                          "Pickaxe",
	ItemTypeScythe:                           "Scythe",
	ItemTypeDofus:                            "Dofus",
	ItemTypeQuestItem:                        "Quest Item",
	ItemTypeDocument:                         "Document",
	ItemTypeSmithmagicPotion:                 "Smithmagic Potion",
	ItemTypeMutationItem:                     "Mutation Item",
	ItemTypeBoostFood:                        "Boost Food",
	ItemTypeBlessing:                         "Blessing",
	ItemTypeCurse:                            "Curse",
	ItemTypeRoleplayBuff:                     "Roleplay Buff",
	ItemTypeFollowingCharacter:               "Following Character",
	ItemTypeBread:                            "Bread",
	ItemTypeCereal:                           "Cereal",
	ItemTypeFlower:                           "Flower",
	ItemTypePlant:                            "Plant",
	ItemTypeBeer:                             "Beer",
	ItemTypeWood:                             "Wood",
	ItemTypeOre:                              "Ore",
	ItemTypeAlloy:                            "Alloy",
	ItemTypeFish:                             "Fish",
	ItemTypeCandy:                            "Candy",
	ItemTypeForgetfulnessPotion:              "Forgetfulness Potion",
	ItemTypeUnlearningPotionForProfession:    "Unlearning Potion for Profession",
	ItemTypeSkillLossPotion:                  "Skill Loss Potion",
	ItemTypeFruit:                            "Fruit",
	ItemTypeBone:                             "Bone",
	ItemTypePowder:                           "Powder",
	ItemTypeEdibleFish:                       "Edible Fish",
	ItemTypePreciousStone:                    "Precious Stone",
	ItemTypeStone:                            "Stone",
	ItemTypeFlour:                            "Flour",
	ItemTypeFeather:                          "Feather",
	ItemTypeHair:                             "Hair",
	ItemTypeFabric:                           "Fabric",
	ItemTypeLeather:                          "Leather",
	ItemTypeWool:                             "Wool",
	ItemTypeSeed:                             "Seed",
	ItemTypeSkin:                             "Skin",
	ItemTypeOil:                              "Oil",
	ItemTypeStuffedAnimal:                    "Stuffed Animal",
	ItemTypeGuttedFish:                       "Gutted Fish",
	ItemTypeMeat:                             "Meat",
	ItemTypePreservedMeat:                    "Preserved Meat",
	ItemTypeTail:                             "Tail",
	ItemTypeMetaria:                          "Metaria",
	ItemTypeVegetable:                        "Vegetable",
	ItemTypeEdibleMeat:                       "Edible Meat",
	ItemTypeDyeing:                           "Dyeing",
	ItemTypeAlchemyEquipment:                 "Alchemy Equipment",
	ItemTypePetEgg:                           "PetEgg",
	ItemTypeSkill:                            "Skill",
	ItemTypeFairywork:                        "Fairywork",
	ItemTypeSpellScroll:                      "Spell Scroll",
	ItemTypeCharacteristicScroll:             "Characteristic Scroll",
	ItemTypeBowKennelCertificate:             "Bow Kennel Certificate",
	ItemTypeSmithmagicRune:                   "Smithmagic Rune",
	ItemTypeDrink:                            "Drink",
	ItemTypeMissionItem:                      "Mission Item",
	ItemTypeBackpack:                         "Backpack",
	ItemTypeShield:                           "Shield",
	ItemTypeSoulStone:                        "Soul Stone",
	ItemTypeKeys:                             "Keys",
	ItemTypeFullSoulStone:                    "Full Soul Stone",
	ItemTypeForgetfulnessPotionForCollectors: "Forgetfulness Potion for Collectors",
	ItemTypeSeekerScroll:                     "Seeker Scroll",
	ItemTypeMagicStone:                       "Magic Stone",
	ItemTypeGifts:                            "Gifts",
	ItemTypePetGhost:                         "Pet Ghost",
	ItemTypeDragoturkey:                      "Dragoturkey",
	ItemTypeGobball:                          "Gobball",
	ItemTypeBreedingItem:                     "Breeding Item",
	ItemTypeUsableItem:                       "Usable Item",
	ItemTypePlank:                            "Plank",
	ItemTypeBark:                             "Bark",
	ItemTypeMountCertificate:                 "Mount Certificate",
	ItemTypeRoot:                             "Root",
	ItemTypeCapturingNet:                     "Capturing Net",
	ItemTypeBagOfResources:                   "Bag of Resources",
	ItemTypeCrossbow:                         "Crossbow",
	ItemTypeLeg:                              "Leg",
	ItemTypeWing:                             "Wing",
	ItemTypeEgg:                              "Egg",
	ItemTypeEar:                              "Ear",
	ItemTypeCarapace:                         "Carapace",
	ItemTypeBud:                              "Bud",
	ItemTypeEye:                              "Eye",
	ItemTypeJelly:                            "Jelly",
	ItemTypeShell:                            "Shell",
	ItemTypePrism:                            "Prism",
	ItemTypeLivingObjects:                    "Living Objects",
	ItemTypeMagicWeapon:                      "Magic Weapon",
	ItemTypeFragmentOfShushuSoul:             "Fragment of Shushu Soul",
	ItemTypePetPotion:                        "Pet Potion",
}

func (t ItemType) Validate() error {
	_, ok := ItemTypes[t]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (t ItemType) String() string {
	name, ok := ItemTypes[t]
	if ok {
		return name
	}

	return unknownStr
}
