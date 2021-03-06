package admin

import (
	"github.com/zxdstyle/liey-admin-demo/app/repository"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type Logic struct {
	*bases.BaseLogic
}

func NewLogic() *Logic {
	return &Logic{
		BaseLogic: bases.NewBaseLogic(repository.Admin()),
	}
}
