package retrosvc

import (
	"errors"

	"github.com/happybydefault/logger"

	"github.com/kralamoure/retro"
)

type Config struct {
	GameServerId int
	Repo         retro.Repo
	Logger       logger.Logger
}

type Service struct {
	gameServerId int
	repo         retro.Repo
	logger       logger.Logger
}

func NewService(c Config) (*Service, error) {
	if c.Repo == nil {
		return nil, errors.New("nil repository")
	}
	if c.Logger == nil {
		c.Logger = logger.Noop{}
	}
	svc := &Service{
		gameServerId: c.GameServerId,
		repo:         c.Repo,
		logger:       c.Logger,
	}
	return svc, nil
}
