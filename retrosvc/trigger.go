package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) Triggers(ctx context.Context) (map[string]retro.Trigger, error) {
	return svc.storer.Triggers(ctx)
}

func (svc Service) TriggerByGameMapIdAndCellId(ctx context.Context, gameMapId, cellId int) (retro.Trigger, error) {
	return svc.storer.TriggerByGameMapIdAndCellId(ctx, gameMapId, cellId)
}
