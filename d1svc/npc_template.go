package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) NPCTemplates(ctx context.Context) (map[int]d1.NPCTemplate, error) {
	return svc.repo.NPCTemplates(ctx)
}
