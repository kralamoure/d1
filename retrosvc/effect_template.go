package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) EffectTemplates(ctx context.Context) (map[int]retro.EffectTemplate, error) {
	return svc.repo.EffectTemplates(ctx)
}
