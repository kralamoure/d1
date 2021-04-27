package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) MountTemplates(ctx context.Context) (map[int]d1.MountTemplate, error) {
	return svc.repo.MountTemplates(ctx)
}
