package retrosvc

import (
	"errors"

	"github.com/happybydefault/logging"

	"github.com/kralamoure/retro"
)

type Config struct {
	GameServerId int
	Storer       retro.Storer
	Logger       logging.Logger
}

type Service struct {
	gameServerId int
	storer       retro.Storer
	logging      logging.Logger
}

func NewService(c Config) (*Service, error) {
	if c.Storer == nil {
		return nil, errors.New("nil storer")
	}
	if c.Logger == nil {
		c.Logger = logging.Noop{}
	}
	svc := &Service{
		gameServerId: c.GameServerId,
		storer:       c.Storer,
		logging:      c.Logger,
	}
	return svc, nil
}
