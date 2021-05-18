package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) NPCTemplates(ctx context.Context) (map[int]retro.NPCTemplate, error) {
	return svc.repo.NPCTemplates(ctx)
}
