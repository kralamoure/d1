package retrotyp

const (
	MountCapacityIdTireless MountCapacityId = iota + 1
	MountCapacityIdLoadBearer
	MountCapacityIdReproductive
	MountCapacityIdWise
	MountCapacityIdHardy
	MountCapacityIdInLove
	MountCapacityIdPrecocious
	MountCapacityIdGeneticallyPredisposed
	MountCapacityIdChameleon
)

type MountCapacityId int

type MountCapacity struct {
	Name        string
	Description string
}

var MountCapacitys = map[MountCapacityId]MountCapacity{
	MountCapacityIdTireless: {
		Name:        "Tireless",
		Description: "A tireless mount has a lot more energy than a normal mount and also recovers energy a lot quicker.",
	},
	MountCapacityIdLoadBearer: {
		Name:        "Load Bearer",
		Description: "A load bearer mount can carry a lot more items than a normal mount.",
	},
	MountCapacityIdReproductive: {
		Name:        "Reproductive",
		Description: "A reproductive mount will give birth to more mounts than a normal mount.",
	},
	MountCapacityIdWise: {
		Name:        "Wise",
		Description: "A wise mount evolves twice as fast as a normal mount.",
	},
	MountCapacityIdHardy: {
		Name:        "Hardy",
		Description: "A hardy mount will have a higher chance of becoming hardy compared to a normal mount.",
	},
	MountCapacityIdInLove: {
		Name:        "In Love",
		Description: "A mount in love will have a higher chance of falling in love compared to a normal mount.",
	},
	MountCapacityIdPrecocious: {
		Name:        "Precocious",
		Description: "A precocious mount will become mature quicker than a normal mount.",
	},
	MountCapacityIdGeneticallyPredisposed: {
		Name:        "GeneticallyPredisposed",
		Description: "A genetically predisposed mount will have a higher chance of passing on its genetic characteristics than a normal mount.",
	},
	MountCapacityIdChameleon: {
		Name:        "Chameleon",
		Description: "Chameleon mounts can change their colours to suit the appearance of the rider.",
	},
}

func (g MountCapacityId) Validate() error {
	_, ok := MountCapacitys[g]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (g MountCapacityId) String() string {
	v, ok := MountCapacitys[g]
	if ok {
		return v.Name
	}

	return unknownStr
}

func (g MountCapacityId) Description() string {
	v, ok := MountCapacitys[g]
	if ok {
		return v.Description
	}

	return unknownStr
}
