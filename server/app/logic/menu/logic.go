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
	if err := repository.Menu.All(ctx, req, menus); err != nil {
		return err
	}

	refer := make(map[uint]*model.Menu, 0)
	tree := make([]*model.Menu, 0)
	for idx, menu := range *menus {
		refer[menu.ID] = (*menus)[idx]
	}

	for idx, menu := range *menus {
		pid := *menu.ParentId
		if pid == 0 {
			tree = append(tree, (*menus)[idx])
		} else {
			if _, ok := refer[pid]; ok {
				if refer[pid].Children == nil {
					refer[pid].Children = &model.Menus{}
				}
				*refer[pid].Children = append(*refer[pid].Children, (*menus)[idx])
			}
		}
	}
	*menus = tree
	return nil
}
