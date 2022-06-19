package menu

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Repository interface {
		bases.Repository
		GetMenusByRoles(ctx context.Context, roles []uint, menus *model.Menus) error
		GetChildren(ctx context.Context, pid uint, menus *model.Menus) error
		AttachRoles(ctx context.Context, menu *model.Menu, roles *model.Roles) error
		SyncRoles(ctx context.Context, menu *model.Menu, roles *model.Roles) error
		MakeTreeData(menus *model.Menus) error
		SortTreeData(tree *model.Menus)
	}

	TreeRepository interface {
		Set(ctx context.Context, menus ...*model.Menu) error
		Get(ctx context.Context, key uint) (*model.Menu, error)
		Del(ctx context.Context, keys ...uint) error
		Data(ctx context.Context) map[uint]*model.Menu
		TreeData() model.Menus
	}
)
