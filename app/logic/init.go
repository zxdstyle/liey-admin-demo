package logic

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/zxdstyle/liey-admin-demo/app/logic/admin"
	"github.com/zxdstyle/liey-admin-demo/app/logic/auth"
	"github.com/zxdstyle/liey-admin-demo/app/logic/menu"
	"github.com/zxdstyle/liey-admin-demo/app/logic/permission"
	"github.com/zxdstyle/liey-admin-demo/app/logic/role"
)

var (
	logics = gmap.NewStrAnyMap(true)
)

func Auth() *auth.Logic {
	return logics.GetOrSetFuncLock("logics.auth", func() interface{} {
		return auth.NewLogic()
	}).(*auth.Logic)
}

func Admin() *admin.Logic {
	return logics.GetOrSetFuncLock("logics.admin", func() interface{} {
		return admin.NewLogic()
	}).(*admin.Logic)
}

func Menu() *menu.Logic {
	return logics.GetOrSetFuncLock("logics.menu", func() interface{} {
		return menu.NewLogic()
	}).(*menu.Logic)
}

func Role() *role.Logic {
	return logics.GetOrSetFuncLock("logics.role", func() interface{} {
		return role.NewLogic()
	}).(*role.Logic)
}

func Permission() *permission.Logic {
	return logics.GetOrSetFuncLock("logics.permission", func() interface{} {
		return permission.NewLogic()
	}).(*permission.Logic)
}
