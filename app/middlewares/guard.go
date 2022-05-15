package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zxdstyle/liey-admin-demo/pkg/guards"
)

func Guard(r *ghttp.Request) {
	authID := r.GetCtxVar("AuthID").Uint()
	// 获取请求的PATH
	path := r.Request.URL.Path
	// 获取请求方法
	act := r.Request.Method

	if err := guards.GetGuard().Can(authID, path, act); err != nil {
		r.SetError(err)
		return
	}
	r.Middleware.Next()
}
