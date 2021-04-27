package d1svc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/d1"
)

func (svc Service) CreateCharacterItem(ctx context.Context, item d1.CharacterItem) (int, error) {
	err := validation.Validate(item)
	if err != nil {
		return 0, err
	}

	return svc.repo.CreateCharacterItem(ctx, item)
}

func (svc Service) UpdateCharacterItem(ctx context.Context, item d1.CharacterItem) error {
	err := validation.Validate(item)
	if err != nil {
		return err
	}

	return svc.repo.UpdateCharacterItem(ctx, item)
}

func (svc Service) DeleteCharacterItem(ctx context.Context, id int) error {
	return svc.repo.DeleteCharacterItem(ctx, id)
}

func (svc Service) CharacterItemsByCharacterId(ctx context.Context, characterId int) (map[int]d1.CharacterItem, error) {
	return svc.repo.CharacterItemsByCharacterId(ctx, characterId)
}

func (svc Service) CharacterItem(ctx context.Context, id int) (d1.CharacterItem, error) {
	return svc.repo.CharacterItem(ctx, id)
}
