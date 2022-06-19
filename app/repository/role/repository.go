package role

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type Repository interface {
	bases.Repository
	GetRoles(ctx context.Context, uid uint, roles *model.Roles) error
	AttachPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error
	SyncPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error
}
