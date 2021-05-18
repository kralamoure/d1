package retrotyp

const (
	ItemConditionMP ItemCondition = "CM"

	ItemConditionVitality     ItemCondition = "CV"
	ItemConditionWisdom       ItemCondition = "CW"
	ItemConditionStrength     ItemCondition = "CS"
	ItemConditionIntelligence ItemCondition = "CI"
	ItemConditionChance       ItemCondition = "CC"
	ItemConditionAgility      ItemCondition = "CA"

	ItemConditionVitalityBase     ItemCondition = "Cv"
	ItemConditionWisdomBase       ItemCondition = "Cw"
	ItemConditionStrengthBase     ItemCondition = "Cs"
	ItemConditionIntelligenceBase ItemCondition = "Ci"
	ItemConditionChanceBase       ItemCondition = "Cc"
	ItemConditionAgilityBase      ItemCondition = "Ca"

	ItemConditionName       ItemCondition = "PN"
	ItemConditionSex        ItemCondition = "PS"
	ItemConditionClass      ItemCondition = "PG"
	ItemConditionAppearance ItemCondition = "PM"
	ItemConditionSubscriber ItemCondition = "PZ"
	ItemConditionKamas      ItemCondition = "PK"
	ItemConditionItem       ItemCondition = "PO"
	ItemConditionEmote      ItemCondition = "PE"
	ItemConditionLevel      ItemCondition = "PL"
	ItemConditionWedding    ItemCondition = "PR"

	ItemConditionAlignment               ItemCondition = "Ps"
	ItemConditionAlignmentLevel          ItemCondition = "Pa"
	ItemConditionAlignmentGift           ItemCondition = "Pg"
	ItemConditionAlignmentSpecialization ItemCondition = "Pr"
	ItemConditionAlignmentGrade          ItemCondition = "PP"

	ItemConditionProfession        ItemCondition = "PJ"
	ItemConditionCurrentProfession ItemCondition = "Pj"

	ItemConditionUnusable ItemCondition = "BI"
)

type ItemCondition string

var ItemConditions = map[ItemCondition]string{
	ItemConditionMP: "MP",

	ItemConditionVitality:     "Vitality",
	ItemConditionWisdom:       "Wisdom",
	ItemConditionStrength:     "Strength",
	ItemConditionIntelligence: "Intelligence",
	ItemConditionChance:       "Chance",
	ItemConditionAgility:      "Agility",

	ItemConditionVitalityBase:     "Base Vitality",
	ItemConditionWisdomBase:       "Base Wisdom",
	ItemConditionStrengthBase:     "Base Strength",
	ItemConditionIntelligenceBase: "Base Intelligence",
	ItemConditionChanceBase:       "Base Chance",
	ItemConditionAgilityBase:      "Base Agility",

	ItemConditionName:       "Name",
	ItemConditionSex:        "Sex",
	ItemConditionClass:      "Class",
	ItemConditionAppearance: "Appearance",
	ItemConditionSubscriber: "Subscriber",
	ItemConditionKamas:      "Kamas",
	ItemConditionItem:       "Item",
	ItemConditionEmote:      "Emote",
	ItemConditionLevel:      "Level",
	ItemConditionWedding:    "Wedding",

	ItemConditionAlignment:               "Alignment",
	ItemConditionAlignmentLevel:          "Alignment Level",
	ItemConditionAlignmentGift:           "Alignment Gift",
	ItemConditionAlignmentSpecialization: "Alignment Specialization",
	ItemConditionAlignmentGrade:          "Alignment Grade",

	ItemConditionProfession:        "Profession",
	ItemConditionCurrentProfession: "Current Profession",

	ItemConditionUnusable: "Unusable",
}

func (id ItemCondition) Validate() error {
	_, ok := ItemConditions[id]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (id ItemCondition) String() string {
	name, ok := ItemConditions[id]
	if ok {
		return name
	}

	return unknownStr
}
