package repository

import (
	"github.com/zxdstyle/liey-admin-demo/app/repository/admin"
	"github.com/zxdstyle/liey-admin-demo/app/repository/menu"
	"github.com/zxdstyle/liey-admin-demo/app/repository/permission"
	"github.com/zxdstyle/liey-admin-demo/app/repository/role"
)

var (
	Admin      admin.Repository      = admin.NewDbRepository()
	Role       role.Repository       = role.NewDbRepository()
	Menu       menu.Repository       = menu.NewDbRepository()
	Permission permission.Repository = permission.NewDbRepository()
)
