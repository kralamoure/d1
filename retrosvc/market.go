package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) Markets(ctx context.Context) (map[string]retro.Market, error) {
	return svc.storer.Markets(ctx, svc.gameServerId)
}
