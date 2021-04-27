package d1svc

import (
	"context"

	"github.com/kralamoure/d1"
)

func (svc Service) NPCResponses(ctx context.Context) (map[int]d1.NPCResponse, error) {
	return svc.repo.NPCResponses(ctx)
}
