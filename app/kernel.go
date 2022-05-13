package app

import (
	"github.com/zxdstyle/liey-admin-demo/app/jobs"
	scaffold "github.com/zxdstyle/liey-admin-scaffold"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/auth/guards"
	"github.com/zxdstyle/liey-admin/framework/plugins"
	"github.com/zxdstyle/liey-admin/framework/queue/job"
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

// Queues 注册队列
func (Kernel) Queues() []job.Job {
	return []job.Job{
		jobs.SendEmail{},
	}
}
