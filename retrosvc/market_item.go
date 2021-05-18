package retrosvc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/retro"
)

func (svc Service) CreateMarketItem(ctx context.Context, item retro.MarketItem) (int, error) {
	err := validation.Validate(item)
	if err != nil {
		return 0, err
	}

	return svc.repo.CreateMarketItem(ctx, item)
}

func (svc Service) DeleteMarketItem(ctx context.Context, id int) error {
	return svc.repo.DeleteMarketItem(ctx, id)
}

func (svc Service) MarketItems(ctx context.Context) (map[int]retro.MarketItem, error) {
	return svc.repo.MarketItems(ctx, svc.gameServerId)
}

func (svc Service) MarketItemsByMarketId(ctx context.Context, marketId string) (map[int]retro.MarketItem, error) {
	return svc.repo.MarketItemsByMarketId(ctx, marketId)
}

func (svc Service) MarketItem(ctx context.Context, id int) (retro.MarketItem, error) {
	return svc.repo.MarketItem(ctx, id)
}
