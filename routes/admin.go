package routes

import (
	"github.com/zxdstyle/liey-admin-demo/app/handler"
	"github.com/zxdstyle/liey-admin-scaffold"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/server"
)

func init() {
	s := adm.Server()
	s.Group("api/v1", func(group *server.RouterGroup) {

		// 脚手架路由注册
		scaffold.RegisterRoutes(group)

		group.POST("login", handler.LoginByPassword)

	})
}
