package app

import (
	"github.com/zxdstyle/liey-admin-demo/app/commands"
	"github.com/zxdstyle/liey-admin-demo/app/events"
	"github.com/zxdstyle/liey-admin-demo/app/jobs"
	"github.com/zxdstyle/liey-admin-demo/app/subscribers"
	_ "github.com/zxdstyle/liey-admin-demo/routes"
	"github.com/zxdstyle/liey-admin/console"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/plugins"
	"github.com/zxdstyle/liey-admin/framework/queue/job"
)

type Kernel struct {
}

// Boot 启动前进行初始化设置
func (Kernel) Boot() {
	//g.Log().SetHandlers(logger.LoggingJsonHandler)
	console.RegisterCmd(commands.InstallCommand)

	adm.ListenEvent(events.UserLogin{}, subscribers.SendNotification{}, subscribers.SendEmail{})
}

// Plugins 注册插件，确保插件唯一性，不允许重复注册插件
func (Kernel) Plugins() []plugins.Plugin {
	return []plugins.Plugin{}
}

func (Kernel) Events() {

}

// Queues 注册队列
func (Kernel) Queues() []job.Job {
	return []job.Job{
		jobs.SendEmail{},
	}
}
