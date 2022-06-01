package menu

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin-demo/app/repository"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
)

type Logic struct {
	*bases.BaseLogic
}

func NewLogic() *Logic {
	return &Logic{
		BaseLogic: bases.NewBaseLogic(repository.Menu),
	}
}

func (*Logic) TreeData(ctx context.Context, req requests.Request, menus *model.Menus) error {
	return repository.Menu.TreeData(ctx, menus)
}

func (*Logic) Create(ctx context.Context, mo bases.RepositoryModel) error {
	val := mo.(*model.Menu)
	val.Children = nil
	return repository.Menu.Create(ctx, mo)
}

func (*Logic) Update(ctx context.Context, mo bases.RepositoryModel) error {
	val := mo.(*model.Menu)
	val.Children = nil
	return repository.Menu.Update(ctx, mo)
}

func (l *Logic) Destroy(ctx context.Context, mo bases.RepositoryModel) error {
	var children model.Menus
	if err := repository.Menu.GetChildren(ctx, mo.GetKey(), &children); err != nil {
		return err
	}

	if err := l.BaseLogic.Destroy(ctx, mo); err != nil {
		return err
	}

	return repository.Menu.BatchDestroy(ctx, children)
}
