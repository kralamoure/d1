package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) ItemSets(ctx context.Context) (map[int]retro.ItemSet, error) {
	return svc.repo.ItemSets(ctx)
}
