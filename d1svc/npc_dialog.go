package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) NPCDialogs(ctx context.Context) (map[int]d1.NPCDialog, error) {
	return svc.repo.NPCDialogs(ctx)
}
