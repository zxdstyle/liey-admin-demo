package role

import (
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type DbRepository struct {
	*bases.GormRepository
}

var _ Repository = DbRepository{}

func NewDbRepository() *DbRepository {
	return &DbRepository{
		GormRepository: bases.NewGormRepository(adm.DB().Model(model.User{})),
	}
}
