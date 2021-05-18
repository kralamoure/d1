package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) Triggers(ctx context.Context) (map[string]retro.Trigger, error) {
	return svc.repo.Triggers(ctx)
}

func (svc Service) TriggerByGameMapIdAndCellId(ctx context.Context, gameMapId, cellId int) (retro.Trigger, error) {
	return svc.repo.TriggerByGameMapIdAndCellId(ctx, gameMapId, cellId)
}
