package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) ItemTemplates(ctx context.Context) (map[int]d1.ItemTemplate, error) {
	return svc.repo.ItemTemplates(ctx)
}
