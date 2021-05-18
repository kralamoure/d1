package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) NPCDialogs(ctx context.Context) (map[int]retro.NPCDialog, error) {
	return svc.repo.NPCDialogs(ctx)
}
