package permission

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
		BaseLogic: bases.NewBaseLogic(repository.Permission()),
	}
}

func (*Logic) Create(ctx context.Context, mo bases.RepositoryModel) error {
	val := mo.(*model.Permission)
	val.Children = nil
	return repository.Permission().Update(ctx, mo)
}

func (*Logic) Update(ctx context.Context, mo bases.RepositoryModel) error {
	val := mo.(*model.Permission)
	val.Children = nil
	return repository.Permission().Update(ctx, mo)
}

func (*Logic) TreeData(ctx context.Context, req requests.Request, permissions *model.Permissions) error {
	return repository.Permission().TreeData(ctx, permissions)
}

func (l *Logic) Destroy(ctx context.Context, mo bases.RepositoryModel) error {
	var children model.Permissions
	if err := repository.Permission().GetChildren(ctx, mo.GetKey(), &children); err != nil {
		return err
	}

	if err := l.BaseLogic.Destroy(ctx, mo); err != nil {
		return err
	}

	return repository.Permission().BatchDestroy(ctx, children)
}
