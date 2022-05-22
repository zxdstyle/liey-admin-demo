package routes

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin-demo/app/handler"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/middleware"
	"github.com/zxdstyle/liey-admin/framework/http/server"
)

func init() {
	s := adm.Server()
	s.Group("api/v1", func(v1 *server.RouterGroup) {
		v1.Middleware(ghttp.MiddlewareCORS)
		v1.POST("login", handler.Auth.Login)
		v1.Resource("admins", handler.Admin{})
		v1.Middleware(middleware.Authenticate("api"))
		v1.GET("userinfo", handler.Auth.Userinfo)
		v1.Resource("menus", handler.ApiMenu)
		v1.GET("tree-menus", handler.ApiMenu.TreeData)
	})
}
