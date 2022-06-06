package model

import "github.com/zxdstyle/liey-admin/framework/http/bases"

type RoleHasPermission struct {
	bases.OnlyKeyModel
	RoleId       uint `json:"role_id" gorm:"not null;index:idx_role_permission,unique;comment:角色ID"`
	PermissionId uint `json:"permission_id" gorm:"not null;index:idx_role_permission;comment:权限ID"`
}
