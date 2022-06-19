package permission

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin-demo/app/enums"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
	"sort"
)

type dbRepository struct {
	*bases.GormRepository
	memory CacheRepository
}

func NewDbRepository() *dbRepository {
	r := &dbRepository{
		GormRepository: bases.NewGormRepository(adm.DB().Model(model.Permission{})),
		memory:         newMemoryRepository(),
	}
	r.doInit()
	return r
}

func (db *dbRepository) doInit() {
	ctx := context.Background()
	var permissions model.Permissions
	if err := db.Orm.WithContext(ctx).Find(&permissions).Error; err != nil {
		g.Log().Error(ctx, err)
	}
	if err := db.memory.Sets(ctx, permissions); err != nil {
		g.Log().Error(ctx, err)
	}
}

// GetRoutes 获取指定类型的权限
func (db *dbRepository) GetRoutes(ctx context.Context, permissions *model.Permissions) error {
	return db.Orm.WithContext(ctx).Where("`type` = ? OR `type` = ?", enums.PermissionTypeMenu, enums.PermissionTypePage).Find(permissions).Error
}

func (db *dbRepository) TreeData(ctx context.Context, permissions *model.Permissions) error {
	data := db.memory.Data(ctx)
	*permissions = db.makeTreeData(ctx, data)
	return nil
}

func (db *dbRepository) Create(ctx context.Context, mo bases.RepositoryModel) error {
	if err := db.GormRepository.Create(ctx, mo); err != nil {
		return err
	}
	return db.memory.Set(ctx, *mo.(*model.Permission))
}

func (db *dbRepository) Update(ctx context.Context, mo bases.RepositoryModel) error {
	if err := db.GormRepository.Update(ctx, mo); err != nil {
		return err
	}
	return db.memory.Set(ctx, *mo.(*model.Permission))
}

func (db *dbRepository) Destroy(ctx context.Context, mo bases.RepositoryModel) error {
	if err := db.GormRepository.Destroy(ctx, mo); err != nil {
		return err
	}
	return db.memory.Del(ctx, mo.GetKey())
}

func (db *dbRepository) GetChildren(ctx context.Context, pid uint, permissions *model.Permissions) error {
	data := db.memory.Data(ctx)
	db.resolveChildren(pid, data, permissions)
	return nil
}

func (db *dbRepository) resolveChildren(pid uint, data map[uint]model.Permission, permissions *model.Permissions) {
	for idx, child := range data {
		if child.GetKey() == pid {
			mo := data[idx]
			*permissions = append(*permissions, &mo)
			db.resolveChildren(child.ID, data, permissions)
		}
	}
}

func (db *dbRepository) makeTreeData(ctx context.Context, permissions map[uint]model.Permission) model.Permissions {
	refer := make(map[uint]*model.Permission, 0)
	tree := make(model.Permissions, 0)
	for idx, Permissions := range permissions {
		val := (permissions)[idx]
		refer[Permissions.ID] = &val
	}

	for idx, permission := range refer {
		pid := *permission.ParentId
		if pid == 0 {
			tree = append(tree, (refer)[idx])
		} else {
			if _, ok := refer[pid]; ok {
				if refer[pid].Children == nil {
					refer[pid].Children = &model.Permissions{}
				}
				*refer[pid].Children = append(*refer[pid].Children, (refer)[idx])
			}
		}
	}
	db.sortTreeData(&tree)
	return tree
}

func (db *dbRepository) sortTreeData(tree *model.Permissions) {
	sort.SliceStable(*tree, func(i, j int) bool {
		return *((*tree)[i].SortNum) > *((*tree)[j].SortNum) || (*tree)[i].ID > (*tree)[j].ID
	})

	for _, Permissions := range *tree {
		if Permissions.Children != nil && len(*Permissions.Children) > 0 {
			db.sortTreeData(Permissions.Children)
		}
	}
}
