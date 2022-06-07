package logic

import (
	"github.com/zxdstyle/liey-admin-demo/app/http/logic/admin"
	"github.com/zxdstyle/liey-admin-demo/app/http/logic/auth"
	"github.com/zxdstyle/liey-admin-demo/app/http/logic/menu"
	"github.com/zxdstyle/liey-admin-demo/app/http/logic/permission"
	"github.com/zxdstyle/liey-admin-demo/app/http/logic/role"
)

var (
	Auth       = auth.NewLogic()
	Admin      = admin.NewLogic()
	Menu       = menu.NewLogic()
	Role       = role.NewLogic()
	Permission = permission.NewLogic()
)
