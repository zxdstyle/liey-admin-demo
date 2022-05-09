package app

import (
	scaffold "github.com/zxdstyle/liey-admin-scaffold"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/auth/guards"
	"github.com/zxdstyle/liey-admin/framework/plugins"
)

type Kernel struct {
}

// Boot 服务启动前进行初始化设置
func (Kernel) Boot() {
	//g.Log().SetHandlers(logger.LoggingJsonHandler)
	guards.SetGuard("api", guards.NewJWTGuard(model.Admin{}))
}

// Plugins 注册插件，确保插件唯一性，不允许重复注册插件
func (Kernel) Plugins() []plugins.Plugin {
	return []plugins.Plugin{
		plugins.WithRename(scaffold.Plugin{}, "scaffolds"),
	}
}
