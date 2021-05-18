package retrotyp

const (
	ExchangeNPCSell Exchange = 10
	ExchangeNPCBuy  Exchange = 11
	ExchangePaddock Exchange = 16
)

type Exchange int

var Exchanges = map[Exchange]string{
	ExchangeNPCSell: "NPC Sell",
	ExchangeNPCBuy:  "NPC Buy",
	ExchangePaddock: "Paddock",
}

func (id Exchange) Validate() error {
	_, ok := Exchanges[id]
	if !ok {
		return ErrInvalidValue
	}

	return nil
}

func (id Exchange) String() string {
	name, ok := Exchanges[id]
	if ok {
		return name
	}

	return unknownStr
}
