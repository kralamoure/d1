package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) NPCs(ctx context.Context) (map[string]retro.NPC, error) {
	return svc.repo.NPCs(ctx, svc.gameServerId)
}
