package menu

import (
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type dbRepository struct {
	*bases.GormRepository
}

func NewDbRepository() *dbRepository {
	return &dbRepository{
		GormRepository: bases.NewGormRepository(adm.DB().Model(model.Menu{})),
	}
}
