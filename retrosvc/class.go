package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
	"github.com/kralamoure/retro/retrotyp"
)

func (svc Service) Classes(ctx context.Context) (map[retrotyp.ClassId]retro.Class, error) {
	return svc.repo.Classes(ctx)
}
