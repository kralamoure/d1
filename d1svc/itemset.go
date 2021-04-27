package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) ItemSets(ctx context.Context) (map[int]d1.ItemSet, error) {
	return svc.repo.ItemSets(ctx)
}
