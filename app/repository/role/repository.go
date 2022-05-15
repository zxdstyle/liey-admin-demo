package role

import (
	"context"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type Repository interface {
	bases.Repository
	AttachPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error
	SyncPermissions(ctx context.Context, role *model.Role, permissions *model.Permissions) error
}
