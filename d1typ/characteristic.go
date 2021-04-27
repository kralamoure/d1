package d1typ

type Characteristic struct {
	Id        CharacteristicId
	Base      int
	Equipment int
	Feat      int
	Boost     int
}

func (c Characteristic) Total() int {
	return c.Base + c.Equipment + c.Feat + c.Boost
}
