package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) EffectTemplates(ctx context.Context) (map[int]d1.EffectTemplate, error) {
	return svc.repo.EffectTemplates(ctx)
}
