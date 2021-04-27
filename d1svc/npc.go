package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) NPCs(ctx context.Context) (map[string]d1.NPC, error) {
	return svc.repo.NPCs(ctx, svc.gameServerId)
}
