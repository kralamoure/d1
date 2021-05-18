package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) ItemTemplates(ctx context.Context) (map[int]retro.ItemTemplate, error) {
	return svc.storer.ItemTemplates(ctx)
}
