package d1typ

const (
	CharacteristicIdAP CharacteristicId = 1
	CharacteristicIdMP CharacteristicId = 23

	CharacteristicIdInitiative  CharacteristicId = 44
	CharacteristicIdProspecting CharacteristicId = 48

	CharacteristicIdVitality CharacteristicId = 11
	CharacteristicIdWisdom   CharacteristicId = 12

	CharacteristicIdStrength     CharacteristicId = 10
	CharacteristicIdIntelligence CharacteristicId = 15
	CharacteristicIdChance       CharacteristicId = 13
	CharacteristicIdAgility      CharacteristicId = 14

	CharacteristicIdRange                     CharacteristicId = 19
	CharacteristicIdMaxSummonedCreaturesBoost CharacteristicId = 26

	CharacteristicIdDamages              CharacteristicId = 16
	CharacteristicIdPhysicalDamages      CharacteristicId = 32
	CharacteristicIdWeaponDamagesPercent CharacteristicId = 31
	CharacteristicIdDamagesPercent       CharacteristicId = 25
	CharacteristicIdTrapDamages          CharacteristicId = 70
	CharacteristicIdTrapDamagesPercent   CharacteristicId = 69
	CharacteristicIdHeals                CharacteristicId = 49
	CharacteristicIdDamagesReflection    CharacteristicId = 50

	CharacteristicIdCriticalHits     CharacteristicId = 18
	CharacteristicIdCriticalFailures CharacteristicId = 39

	CharacteristicIdDodgeAP CharacteristicId = 27
	CharacteristicIdDodgeMP CharacteristicId = 28

	CharacteristicIdNeutralResistance CharacteristicId = 58
	CharacteristicIdEarthResistance   CharacteristicId = 54
	CharacteristicIdFireResistance    CharacteristicId = 55
	CharacteristicIdWaterResistance   CharacteristicId = 56
	CharacteristicIdAirResistance     CharacteristicId = 57

	CharacteristicIdNeutralResistancePercent CharacteristicId = 37
	CharacteristicIdEarthResistancePercent   CharacteristicId = 33
	CharacteristicIdFireResistancePercent    CharacteristicId = 34
	CharacteristicIdWaterResistancePercent   CharacteristicId = 35
	CharacteristicIdAirResistancePercent     CharacteristicId = 36

	CharacteristicIdNeutralResistancePVP CharacteristicId = 68
	CharacteristicIdEarthResistancePVP   CharacteristicId = 64
	CharacteristicIdFireResistancePVP    CharacteristicId = 65
	CharacteristicIdWaterResistancePVP   CharacteristicId = 66
	CharacteristicIdAirResistancePVP     CharacteristicId = 67

	CharacteristicIdNeutralResistancePercentPVP CharacteristicId = 63
	CharacteristicIdEarthResistancePercentPVP   CharacteristicId = 59
	CharacteristicIdFireResistancePercentPVP    CharacteristicId = 60
	CharacteristicIdWaterResistancePercentPVP   CharacteristicId = 61
	CharacteristicIdAirResistancePercentPVP     CharacteristicId = 62

	CharacteristicIdDamagesFactor      CharacteristicId = 17
	CharacteristicIdMagicalResistance  CharacteristicId = 20
	CharacteristicIdPhysicalResistance CharacteristicId = 21
	// CharacteristicIdExperienceBoost    CharacteristicId = 22
	// CharacteristicIdInvisibility       CharacteristicId = 24
	// CharacteristicIdEnergyPoints       CharacteristicId = 29
	// CharacteristicIdAlignment          CharacteristicId = 30
	// CharacteristicIdGFX         CharacteristicId = 38
	// CharacteristicIdState       CharacteristicId = 71

	CharacteristicIdMaxWeight CharacteristicId = 100 // Custom
)

type CharacteristicId int

var CharacteristicIds = map[CharacteristicId]string{
	CharacteristicIdAP:                          "AP",
	CharacteristicIdMP:                          "MP",
	CharacteristicIdInitiative:                  "Initiative",
	CharacteristicIdProspecting:                 "Prospecting",
	CharacteristicIdVitality:                    "Vitality",
	CharacteristicIdWisdom:                      "Wisdom",
	CharacteristicIdStrength:                    "Strength",
	CharacteristicIdIntelligence:                "Intelligence",
	CharacteristicIdChance:                      "Chance",
	CharacteristicIdAgility:                     "Agility",
	CharacteristicIdRange:                       "Range",
	CharacteristicIdMaxSummonedCreaturesBoost:   "Max Summoned Creatures Boost",
	CharacteristicIdDamages:                     "Damages",
	CharacteristicIdPhysicalDamages:             "Physical Damages",
	CharacteristicIdWeaponDamagesPercent:        "Weapon Damages Percent",
	CharacteristicIdDamagesPercent:              "Damages Percent",
	CharacteristicIdTrapDamages:                 "Trap Damages",
	CharacteristicIdTrapDamagesPercent:          "Trap Damages Percent",
	CharacteristicIdHeals:                       "Heals",
	CharacteristicIdDamagesReflection:           "Damages Reflection",
	CharacteristicIdCriticalHits:                "Critical Hits",
	CharacteristicIdCriticalFailures:            "Critical Failures",
	CharacteristicIdDodgeAP:                     "Dodge AP",
	CharacteristicIdDodgeMP:                     "Dodge MP",
	CharacteristicIdNeutralResistance:           "Neutral Resistance",
	CharacteristicIdEarthResistance:             "Earth Resistance",
	CharacteristicIdFireResistance:              "Fire Resistance",
	CharacteristicIdWaterResistance:             "Water Resistance",
	CharacteristicIdAirResistance:               "Air Resistance",
	CharacteristicIdNeutralResistancePercent:    "Neutral Resistance Percent",
	CharacteristicIdEarthResistancePercent:      "Earth Resistance Percent",
	CharacteristicIdFireResistancePercent:       "Fire Resistance Percent",
	CharacteristicIdWaterResistancePercent:      "Water Resistance Percent",
	CharacteristicIdAirResistancePercent:        "Air Resistance Percent",
	CharacteristicIdNeutralResistancePVP:        "Neutral Resistance PVP",
	CharacteristicIdEarthResistancePVP:          "Earth Resistance PVP",
	CharacteristicIdFireResistancePVP:           "Fire Resistance PVP",
	CharacteristicIdWaterResistancePVP:          "Water Resistance PVP",
	CharacteristicIdAirResistancePVP:            "Air Resistance PVP",
	CharacteristicIdNeutralResistancePercentPVP: "Neutral Resistance Percent PVP",
	CharacteristicIdEarthResistancePercentPVP:   "Earth Resistance Percent PVP",
	CharacteristicIdFireResistancePercentPVP:    "Fire Resistance Percent PVP",
	CharacteristicIdWaterResistancePercentPVP:   "Water Resistance Percent PVP",
	CharacteristicIdAirResistancePercentPVP:     "Air Resistance Percent PVP",
	CharacteristicIdDamagesFactor:               "Damages Factor",
	CharacteristicIdMagicalResistance:           "Magical Resistance",
	CharacteristicIdPhysicalResistance:          "Physical Resistance",
	// CharacteristicIdExperienceBoost: "Experience Boost",
	// CharacteristicIdInvisibility: "Invisibility",
	// CharacteristicIdEnergyPoints: "Energy Points",
	// CharacteristicIdAlignment: "Alignment",
	// CharacteristicIdGFX: "GFX",
	// CharacteristicIdState: "State",
	CharacteristicIdMaxWeight: "Max Weight",
}

func (s CharacteristicId) Validate() error {
	_, ok := CharacteristicIds[s]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (s CharacteristicId) String() string {
	name, ok := CharacteristicIds[s]
	if ok {
		return name
	}

	return unknownStr
}
