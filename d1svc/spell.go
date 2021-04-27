package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) Spells(ctx context.Context) (map[int]d1.Spell, error) {
	return svc.repo.Spells(ctx)
}
