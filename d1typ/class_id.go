package d1typ

const (
	ClassIdFeca ClassId = 1 + iota
	ClassIdOsamodas
	ClassIdEnutrof
	ClassIdSram
	ClassIdXelor
	ClassIdEcaflip
	ClassIdEniripsa
	ClassIdIop
	ClassIdCra
	ClassIdSadida
	ClassIdSacrier
	ClassIdPandawa
)

type ClassId int

var ClassIds = map[ClassId]string{
	ClassIdFeca:     "Feca",
	ClassIdOsamodas: "Osamodas",
	ClassIdEnutrof:  "Enutrof",
	ClassIdSram:     "Sram",
	ClassIdXelor:    "Xelor",
	ClassIdEcaflip:  "Ecaflip",
	ClassIdEniripsa: "Eniripsa",
	ClassIdIop:      "Iop",
	ClassIdCra:      "Cra",
	ClassIdSadida:   "Sadida",
	ClassIdSacrier:  "Sacrier",
	ClassIdPandawa:  "Pandawa",
}

func (c ClassId) Validate() error {
	_, ok := ClassIds[c]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (c ClassId) String() string {
	name, ok := ClassIds[c]
	if ok {
		return name
	}

	return unknownStr
}
