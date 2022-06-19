package menu

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/clauses"
	"sort"
)

type dbRepository struct {
	*bases.GormRepository
	cache TreeRepository
}

func NewDbRepository() *dbRepository {
	return &dbRepository{
		GormRepository: bases.NewGormRepository(adm.DB().Model(model.Menu{})),
	}
}

func (db *dbRepository) GetMenusByRoles(ctx context.Context, roles []uint, menus *model.Menus) error {
	if err := db.Orm.WithContext(ctx).Joins("INNER JOIN `role_has_menus` ON `menus`.`id` = `role_has_menus`.`menu_id`").Where("`role_id` IN ?", roles).Find(menus).Error; err != nil {
		return err
	}
	return db.MakeTreeData(menus)
}

func (db *dbRepository) GetChildren(ctx context.Context, pid uint, menus *model.Menus) error {
	data := db.cache.Data(ctx)
	db.resolveChildren(pid, data, menus)
	return nil
}

func (db *dbRepository) resolveChildren(pid uint, data map[uint]*model.Menu, menus *model.Menus) {
	for idx, child := range data {
		if child.GetKey() == pid {
			*menus = append(*menus, data[idx])
			db.resolveChildren(child.ID, data, menus)
		}
	}
}

// AttachRoles 添加角色
func (db *dbRepository) AttachRoles(ctx context.Context, menu *model.Menu, roles *model.Roles) error {
	if roles == nil || menu == nil {
		return nil
	}

	if err := db.Orm.WithContext(ctx).Model(menu).Omit("Roles.*").Association("Roles").Append(roles); err != nil {
		return err
	}

	return db.GormRepository.Show(ctx, []clauses.Clause{clauses.NewPreloadClause("with.roles", "*")}, menu)
}

// SyncRoles 同步角色
func (db *dbRepository) SyncRoles(ctx context.Context, menu *model.Menu, roles *model.Roles) error {
	if roles == nil || menu == nil {
		return nil
	}

	tx := db.Orm.WithContext(ctx).Begin()

	if err := tx.Model(menu).Association("Roles").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(menu).Omit("Roles.*").Association("Roles").Append(roles); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return db.GormRepository.Show(ctx, []clauses.Clause{clauses.NewPreloadClause("with.roles", "*")}, menu)
}

func (db *dbRepository) MakeTreeData(menus *model.Menus) error {
	if menus == nil {
		return nil
	}
	refer := make(map[uint]*model.Menu, 0)
	tree := make(model.Menus, 0)
	for idx, menu := range *menus {
		refer[menu.ID] = (*menus)[idx]
	}

	for _, menu := range *menus {
		pid := *menu.ParentId
		if pid == 0 {
			tree = append(tree, refer[menu.GetKey()])
		} else {
			if _, ok := refer[pid]; ok {
				if refer[pid].Children == nil {
					refer[pid].Children = &model.Menus{}
				}
				*refer[pid].Children = append(*refer[pid].Children, refer[menu.GetKey()])
			}
		}
	}
	db.SortTreeData(&tree)
	*menus = tree
	return nil
}

func (db *dbRepository) SortTreeData(tree *model.Menus) {
	sort.SliceStable(*tree, func(i, j int) bool {
		return *((*tree)[i].SortNum) > *((*tree)[j].SortNum)
	})

	for _, menu := range *tree {
		if menu.Children != nil && len(*menu.Children) > 0 {
			db.SortTreeData(menu.Children)
		}
	}
}
