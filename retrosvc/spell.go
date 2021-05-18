package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) Spells(ctx context.Context) (map[int]retro.Spell, error) {
	return svc.storer.Spells(ctx)
}
