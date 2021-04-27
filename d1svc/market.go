package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) Markets(ctx context.Context) (map[string]d1.Market, error) {
	return svc.repo.Markets(ctx, svc.gameServerId)
}
