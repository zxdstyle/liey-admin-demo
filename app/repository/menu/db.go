package menu

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type dbRepository struct {
	*bases.GormRepository
	cache TreeRepository
}

func NewDbRepository() *dbRepository {
	r := &dbRepository{
		GormRepository: bases.NewGormRepository(adm.DB().Model(model.Menu{})),
		cache:          newTreeRepository(),
	}
	r.doInit()
	return r
}

func (db *dbRepository) doInit() {
	ctx := context.Background()
	var menus model.Menus
	if err := db.Orm.WithContext(ctx).Find(&menus).Error; err != nil {
		g.Log().Error(ctx, err)
	}
	if err := db.cache.Set(ctx, menus...); err != nil {
		g.Log().Error(ctx, err)
	}
}

func (db *dbRepository) TreeData(ctx context.Context, menus *model.Menus) error {
	data := db.cache.TreeData()
	*menus = data
	return nil
}

func (db *dbRepository) Create(ctx context.Context, mo bases.RepositoryModel) error {
	if err := db.GormRepository.Create(ctx, mo); err != nil {
		return err
	}
	return db.cache.Set(ctx, mo.(*model.Menu))
}

func (db *dbRepository) Update(ctx context.Context, mo bases.RepositoryModel) error {
	if err := db.GormRepository.Update(ctx, mo); err != nil {
		return err
	}
	return db.cache.Set(ctx, mo.(*model.Menu))
}

func (db *dbRepository) Destroy(ctx context.Context, mo bases.RepositoryModel) error {
	if err := db.GormRepository.Destroy(ctx, mo); err != nil {
		return err
	}
	return db.cache.Del(ctx, mo.GetKey())
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
