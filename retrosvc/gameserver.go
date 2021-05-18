package retrosvc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/retro"
	"github.com/kralamoure/retro/retrotyp"
)

func (svc Service) CreateGameServer(ctx context.Context, gameServer retro.GameServer) error {
	err := validation.Validate(gameServer)
	if err != nil {
		return err
	}

	return svc.repo.CreateGameServer(ctx, gameServer)
}

func (svc Service) GameServers(ctx context.Context) (map[int]retro.GameServer, error) {
	return svc.repo.GameServers(ctx)
}

func (svc Service) GameServer(ctx context.Context, id int) (retro.GameServer, error) {
	return svc.repo.GameServer(ctx, id)
}

func (svc Service) SetGameServerState(ctx context.Context, state retrotyp.GameServerState) error {
	return svc.repo.SetGameServerState(ctx, svc.gameServerId, state)
}
