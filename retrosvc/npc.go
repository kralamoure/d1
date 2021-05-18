package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) NPCs(ctx context.Context) (map[string]retro.NPC, error) {
	return svc.storer.NPCs(ctx, svc.gameServerId)
}
