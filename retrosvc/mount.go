package retrosvc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/kralamoure/retro"
)

func (svc Service) CreateMount(ctx context.Context, mount retro.Mount) (int, error) {
	err := validation.Validate(mount)
	if err != nil {
		return 0, err
	}

	return svc.repo.CreateMount(ctx, mount)
}

func (svc Service) UpdateMount(ctx context.Context, mount retro.Mount) error {
	err := validation.Validate(mount)
	if err != nil {
		return err
	}

	return svc.repo.UpdateMount(ctx, mount)
}

func (svc Service) DeleteMount(ctx context.Context, id int) error {
	return svc.repo.DeleteMount(ctx, id)
}

func (svc Service) Mount(ctx context.Context, id int) (retro.Mount, error) {
	return svc.repo.Mount(ctx, id)
}

func (svc Service) Mounts(ctx context.Context) (map[int]retro.Mount, error) {
	return svc.repo.Mounts(ctx)
}

func (svc Service) MountsByCharacterId(ctx context.Context, characterId int) (map[int]retro.Mount, error) {
	return svc.repo.MountsByCharacterId(ctx, characterId)
}
