package permission

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Repository interface {
		bases.Repository
		GetRoutes(ctx context.Context, permissions *model.Permissions) error
		TreeData(ctx context.Context, permissions *model.Permissions) error
		GetChildren(ctx context.Context, pid uint, permissions *model.Permissions) error
	}

	CacheRepository interface {
		Sets(ctx context.Context, mos model.Permissions) error
		Set(ctx context.Context, mo model.Permission) error
		Get(ctx context.Context, key uint) (mo model.Permission, err error)
		Gets(ctx context.Context, keys ...uint) (mos model.Permissions, err error)
		Del(ctx context.Context, keys ...uint) error
		Data(ctx context.Context) map[uint]model.Permission
		Iterator(f func(key uint, permission model.Permission) bool)
	}

	TreeRepository interface {
		Set(ctx context.Context, permissions ...*model.Permission) error
		Get(ctx context.Context, key uint) (*model.Permission, error)
		Del(ctx context.Context, keys ...uint) error
		Data(ctx context.Context) map[uint]*model.Permission
		TreeData() model.Permissions
	}
)
