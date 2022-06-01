package permission

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Repository interface {
		bases.Repository
		TreeData(ctx context.Context, permissions *model.Permissions) error
		GetChildren(ctx context.Context, pid uint, permissions *model.Permissions) error
	}

	TreeRepository interface {
		Set(ctx context.Context, permissions ...*model.Permission) error
		Get(ctx context.Context, key uint) (*model.Permission, error)
		Del(ctx context.Context, keys ...uint) error
		Data(ctx context.Context) map[uint]*model.Permission
		TreeData() model.Permissions
	}
)
