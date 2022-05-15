package admin

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type dbRepository struct {
	*bases.GormRepository
}

func NewDbRepository() Repository {
	return &dbRepository{
		GormRepository: bases.NewGormRepository(adm.DB().Model(model.Admin{})),
	}
}
