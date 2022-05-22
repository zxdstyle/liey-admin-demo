package migrations

import (
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type ScaffoldMigration struct {
}

func (s ScaffoldMigration) Models() []bases.RepositoryModel {
	return []bases.RepositoryModel{
		&model.Admin{},
		&model.Menu{},
		&model.Permission{},
		&model.Role{},
		&model.RoleHasPermission{},
		&model.User{},
	}
}
