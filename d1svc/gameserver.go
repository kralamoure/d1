package d1svc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/d1"
	"github.com/kralamoure/d1/d1typ"
)

func (svc Service) CreateGameServer(ctx context.Context, gameServer d1.GameServer) error {
	err := validation.Validate(gameServer)
	if err != nil {
		return err
	}

	return svc.repo.CreateGameServer(ctx, gameServer)
}

func (svc Service) GameServers(ctx context.Context) (map[int]d1.GameServer, error) {
	return svc.repo.GameServers(ctx)
}

func (svc Service) GameServer(ctx context.Context, id int) (d1.GameServer, error) {
	return svc.repo.GameServer(ctx, id)
}

func (svc Service) SetGameServerState(ctx context.Context, state d1typ.GameServerState) error {
	return svc.repo.SetGameServerState(ctx, svc.gameServerId, state)
}
