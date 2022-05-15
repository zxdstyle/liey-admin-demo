package repository

import (
	"github.com/zxdstyle/liey-admin-demo/app/repository/admin"
	"github.com/zxdstyle/liey-admin-demo/app/repository/permission"
	"github.com/zxdstyle/liey-admin-demo/app/repository/role"
)

var (
	Admin      = admin.NewDbRepository()
	Role       = role.NewDbRepository()
	Permission = permission.NewDbRepository()
)
