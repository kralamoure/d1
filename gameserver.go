package retro

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"github.com/kralamoure/retro/retrotyp"
)

type GameServer struct {
	Id         int
	Host       string
	Port       string
	State      retrotyp.GameServerState
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
