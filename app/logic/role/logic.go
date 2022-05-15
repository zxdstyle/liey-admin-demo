package role

import (
	"context"
	"fmt"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin-scaffold/http/repository"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type Logic struct {
	*bases.BaseLogic
}

func NewLogic() *Logic {
	return &Logic{
		BaseLogic: bases.NewBaseLogic(repository.Role),
	}
}

func (l Logic) Create(ctx context.Context, mo bases.RepositoryModel) error {
	role := mo.(*model.Role)
	if err := l.checkPermissions(ctx, role.Permissions); err != nil {
		return err
	}

	if err := repository.Role.Create(ctx, mo); err != nil {
		return err
	}

	if role.Permissions != nil {
		return repository.Role.AttachPermissions(ctx, role, role.Permissions)
	}
	return nil
}

func (l Logic) Update(ctx context.Context, mo bases.RepositoryModel) error {
	role := mo.(*model.Role)
	if err := l.checkPermissions(ctx, role.Permissions); err != nil {
		return err
	}

	if err := repository.Role.Update(ctx, mo); err != nil {
		return err
	}

	if role.Permissions != nil {
		return repository.Role.SyncPermissions(ctx, role, role.Permissions)
	}
	return nil
}

func (Logic) checkPermissions(ctx context.Context, permissions *model.Permissions) error {
	if permissions == nil {
		return nil
	}

	var keys []uint
	for _, permission := range *permissions {
		keys = append(keys, permission.GetKey())
	}
	var exists bool
	if err := repository.Permission.ExistsByKeys(ctx, keys, &exists); err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("invalid permissions")
	}

	return nil
}
