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

	return svc.storer.CreateMount(ctx, mount)
}

func (svc Service) UpdateMount(ctx context.Context, mount retro.Mount) error {
	err := validation.Validate(mount)
	if err != nil {
		return err
	}

	return svc.storer.UpdateMount(ctx, mount)
}

func (svc Service) DeleteMount(ctx context.Context, id int) error {
	return svc.storer.DeleteMount(ctx, id)
}

func (svc Service) Mount(ctx context.Context, id int) (retro.Mount, error) {
	return svc.storer.Mount(ctx, id)
}

func (svc Service) Mounts(ctx context.Context) (map[int]retro.Mount, error) {
	return svc.storer.Mounts(ctx)
}

func (svc Service) MountsByCharacterId(ctx context.Context, characterId int) (map[int]retro.Mount, error) {
	return svc.storer.MountsByCharacterId(ctx, characterId)
}
