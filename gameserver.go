package d1

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/kralamoure/d1/d1typ"
)

type GameServer struct {
	Id         int
	Host       string
	Port       string
	State      d1typ.GameServerState
	Completion int
}

func (a GameServer) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Host,
			validation.Required,
			is.Host,
		),
		validation.Field(&a.Port,
			validation.Required,
			is.Port,
		),
		validation.Field(&a.State),
		validation.Field(&a.Completion, validation.Min(0)),
	)
}
