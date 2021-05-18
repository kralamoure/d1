package retrosvc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/retro"
)

func (svc Service) CreateCharacter(ctx context.Context, character retro.Character) (int, error) {
	err := validation.Validate(character)
	if err != nil {
		return 0, err
	}

	return svc.storer.CreateCharacter(ctx, character)
}

func (svc Service) UpdateCharacter(ctx context.Context, character retro.Character) error {
	err := validation.Validate(character)
	if err != nil {
		return err
	}

	return svc.storer.UpdateCharacter(ctx, character)
}

func (svc Service) DeleteCharacter(ctx context.Context, id int) error {
	return svc.storer.DeleteCharacter(ctx, id)
}

func (svc Service) AllCharacters(ctx context.Context) (map[int]retro.Character, error) {
	return svc.storer.AllCharacters(ctx)
}

func (svc Service) AllCharactersByAccountId(ctx context.Context, accountId string) (map[int]retro.Character, error) {
	return svc.storer.AllCharactersByAccountId(ctx, accountId)
}

func (svc Service) Characters(ctx context.Context) (map[int]retro.Character, error) {
	return svc.storer.Characters(ctx, svc.gameServerId)
}

func (svc Service) CharactersByAccountId(ctx context.Context, accountId string) (map[int]retro.Character, error) {
	return svc.storer.CharactersByAccountId(ctx, svc.gameServerId, accountId)
}

func (svc Service) CharactersByGameMapId(ctx context.Context, gameMapId int) (map[int]retro.Character, error) {
	return svc.storer.CharactersByGameMapId(ctx, svc.gameServerId, gameMapId)
}

func (svc Service) Character(ctx context.Context, id int) (retro.Character, error) {
	return svc.storer.Character(ctx, id)
}
