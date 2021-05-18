package retrosvc

import (
	"errors"

	"github.com/happybydefault/logger"

	"github.com/kralamoure/retro"
)

type Config struct {
	GameServerId int
	Storer       retro.Storer
	Logger       logger.Logger
}

type Service struct {
	gameServerId int
	storer       retro.Storer
	logger       logger.Logger
}

func NewService(c Config) (*Service, error) {
	if c.Storer == nil {
		return nil, errors.New("nil storer")
	}
	if c.Logger == nil {
		c.Logger = logger.Noop{}
	}
	svc := &Service{
		gameServerId: c.GameServerId,
		storer:       c.Storer,
		logger:       c.Logger,
	}
	return svc, nil
}
