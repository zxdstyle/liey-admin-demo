package permission

import (
	"github.com/zxdstyle/liey-admin-scaffold/http/repository"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type Logic struct {
	*bases.BaseLogic
}

func NewLogic() *Logic {
	return &Logic{
		BaseLogic: bases.NewBaseLogic(repository.Permission),
	}
}
