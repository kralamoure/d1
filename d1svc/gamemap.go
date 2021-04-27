package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) GameMaps(ctx context.Context) (map[int]d1.GameMap, error) {
	return svc.repo.GameMaps(ctx)
}
