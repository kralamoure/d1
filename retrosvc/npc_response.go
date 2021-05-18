package retrosvc

import (
	"context"

	"github.com/kralamoure/retro"
)

func (svc Service) NPCResponses(ctx context.Context) (map[int]retro.NPCResponse, error) {
	return svc.repo.NPCResponses(ctx)
}
