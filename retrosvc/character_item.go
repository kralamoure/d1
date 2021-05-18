package retrosvc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/retro"
)

func (svc Service) CreateCharacterItem(ctx context.Context, item retro.CharacterItem) (int, error) {
	err := validation.Validate(item)
	if err != nil {
		return 0, err
	}

	return svc.storer.CreateCharacterItem(ctx, item)
}

func (svc Service) UpdateCharacterItem(ctx context.Context, item retro.CharacterItem) error {
	err := validation.Validate(item)
	if err != nil {
		return err
	}

	return svc.storer.UpdateCharacterItem(ctx, item)
}

func (svc Service) DeleteCharacterItem(ctx context.Context, id int) error {
	return svc.storer.DeleteCharacterItem(ctx, id)
}

func (svc Service) CharacterItemsByCharacterId(ctx context.Context, characterId int) (map[int]retro.CharacterItem, error) {
	return svc.storer.CharacterItemsByCharacterId(ctx, characterId)
}

func (svc Service) CharacterItem(ctx context.Context, id int) (retro.CharacterItem, error) {
	return svc.storer.CharacterItem(ctx, id)
}
