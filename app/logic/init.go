package logic

import (
	"github.com/zxdstyle/liey-admin-demo/app/logic/admin"
	"github.com/zxdstyle/liey-admin-demo/app/logic/auth"
	"github.com/zxdstyle/liey-admin-demo/app/logic/menu"
	"github.com/zxdstyle/liey-admin-demo/app/logic/permission"
	"github.com/zxdstyle/liey-admin-demo/app/logic/role"
)

var (
	Auth       = auth.NewLogic()
	Admin      = admin.NewLogic()
	Menu       = menu.NewLogic()
	Role       = role.NewLogic()
	Permission = permission.NewLogic()
)
