package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) GameMaps(ctx context.Context) (map[int]retro.GameMap, error) {
	return svc.storer.GameMaps(ctx)
}
