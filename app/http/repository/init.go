package repository

import (
	"github.com/gogf/gf/v2/container/gmap"
	admin2 "github.com/zxdstyle/liey-admin-demo/app/http/repository/admin"
	menu2 "github.com/zxdstyle/liey-admin-demo/app/http/repository/menu"
	permission2 "github.com/zxdstyle/liey-admin-demo/app/http/repository/permission"
	role2 "github.com/zxdstyle/liey-admin-demo/app/http/repository/role"
)

var repositories = gmap.NewStrAnyMap(true)

func Admin() admin2.Repository {
	return repositories.GetOrSetFuncLock("repository.admin", func() interface{} {
		return admin2.NewDbRepository()
	}).(admin2.Repository)
}

func Role() role2.Repository {
	return repositories.GetOrSetFuncLock("repository.role", func() interface{} {
		return role2.NewDbRepository()
	}).(role2.Repository)
}

func Menu() menu2.Repository {
	return repositories.GetOrSetFuncLock("repository.menu", func() interface{} {
		return menu2.NewDbRepository()
	}).(menu2.Repository)
}

func Permission() permission2.Repository {
	return repositories.GetOrSetFuncLock("repository.permission", func() interface{} {
		return permission2.NewDbRepository()
	}).(permission2.Repository)
}
