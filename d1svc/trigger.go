package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) Triggers(ctx context.Context) (map[string]d1.Trigger, error) {
	return svc.repo.Triggers(ctx)
}

func (svc Service) TriggerByGameMapIdAndCellId(ctx context.Context, gameMapId, cellId int) (d1.Trigger, error) {
	return svc.repo.TriggerByGameMapIdAndCellId(ctx, gameMapId, cellId)
}
