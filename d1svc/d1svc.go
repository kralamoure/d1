package d1svc

import (
	"errors"

	"github.com/happybydefault/logger"

	"github.com/kralamoure/d1"
)

type Config struct {
	GameServerId int
	Repo         d1.Repo
	Logger       logger.Logger
}

type Service struct {
	gameServerId int
	repo         d1.Repo
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
