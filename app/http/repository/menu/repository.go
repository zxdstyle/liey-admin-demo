package menu

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Repository interface {
		bases.Repository
		TreeData(ctx context.Context, menus *model.Menus) error
		GetChildren(ctx context.Context, pid uint, menus *model.Menus) error
	}

	TreeRepository interface {
		Set(ctx context.Context, menus ...*model.Menu) error
		Get(ctx context.Context, key uint) (*model.Menu, error)
		Del(ctx context.Context, keys ...uint) error
		Data(ctx context.Context) map[uint]*model.Menu
		TreeData() model.Menus
	}
)
