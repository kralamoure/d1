package d1svc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/d1"
)

func (svc Service) CreateCharacter(ctx context.Context, character d1.Character) (int, error) {
	err := validation.Validate(character)
	if err != nil {
		return 0, err
	}

	return svc.repo.CreateCharacter(ctx, character)
}

func (svc Service) UpdateCharacter(ctx context.Context, character d1.Character) error {
	err := validation.Validate(character)
	if err != nil {
		return err
	}

	return svc.repo.UpdateCharacter(ctx, character)
}

func (svc Service) DeleteCharacter(ctx context.Context, id int) error {
	return svc.repo.DeleteCharacter(ctx, id)
}

func (svc Service) AllCharacters(ctx context.Context) (map[int]d1.Character, error) {
	return svc.repo.AllCharacters(ctx)
}

func (svc Service) AllCharactersByAccountId(ctx context.Context, accountId string) (map[int]d1.Character, error) {
	return svc.repo.AllCharactersByAccountId(ctx, accountId)
}

func (svc Service) Characters(ctx context.Context) (map[int]d1.Character, error) {
	return svc.repo.Characters(ctx, svc.gameServerId)
}

func (svc Service) CharactersByAccountId(ctx context.Context, accountId string) (map[int]d1.Character, error) {
	return svc.repo.CharactersByAccountId(ctx, svc.gameServerId, accountId)
}

func (svc Service) CharactersByGameMapId(ctx context.Context, gameMapId int) (map[int]d1.Character, error) {
	return svc.repo.CharactersByGameMapId(ctx, svc.gameServerId, gameMapId)
}

func (svc Service) Character(ctx context.Context, id int) (d1.Character, error) {
	return svc.repo.Character(ctx, id)
}
