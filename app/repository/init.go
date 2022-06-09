package repository

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/zxdstyle/liey-admin-demo/app/repository/admin"
	"github.com/zxdstyle/liey-admin-demo/app/repository/menu"
	"github.com/zxdstyle/liey-admin-demo/app/repository/permission"
	"github.com/zxdstyle/liey-admin-demo/app/repository/role"
)

var repositories = gmap.NewStrAnyMap(true)

func Admin() admin.Repository {
	return repositories.GetOrSetFuncLock("repository.admin", func() interface{} {
		return admin.NewDbRepository()
	}).(admin.Repository)
}

func Role() role.Repository {
	return repositories.GetOrSetFuncLock("repository.role", func() interface{} {
		return role.NewDbRepository()
	}).(role.Repository)
}

func Menu() menu.Repository {
	return repositories.GetOrSetFuncLock("repository.menu", func() interface{} {
		return menu.NewDbRepository()
	}).(menu.Repository)
}

func Permission() permission.Repository {
	return repositories.GetOrSetFuncLock("repository.permission", func() interface{} {
		return permission.NewDbRepository()
	}).(permission.Repository)
}
