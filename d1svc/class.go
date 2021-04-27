package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
	"github.com/kralamoure/d1/d1typ"
)

func (svc Service) Classes(ctx context.Context) (map[d1typ.ClassId]d1.Class, error) {
	return svc.repo.Classes(ctx)
}
