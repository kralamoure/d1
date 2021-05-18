package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) MountTemplates(ctx context.Context) (map[int]retro.MountTemplate, error) {
	return svc.storer.MountTemplates(ctx)
}
