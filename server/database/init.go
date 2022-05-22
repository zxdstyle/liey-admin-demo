package database

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin-demo/database/migrations"
	migrator "github.com/zxdstyle/liey-admin/framework/database/migrations"
)

func init() {
	ctx := context.Background()
	if err := migrator.RegisterMigration("scaffold", migrations.ScaffoldMigration{}); err != nil {
		g.Log().Error(ctx, err)
	}
}
