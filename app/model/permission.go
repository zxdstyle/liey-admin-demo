package model

import (
	"github.com/zxdstyle/liey-admin-demo/app/enums"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Permission struct {
		bases.Model
		ParentId    *uint                 `gorm:"column:parent_id" json:"parent_id" v:"required,min=0"`                 //父级权限
		Type        *enums.PermissionType `gorm:"column:type" json:"type" v:"required,oneof=menu page action"`          //权限类型
		Slug        *string               `gorm:"column:slug" json:"slug" v:"required,unique-db=permissions"`           //唯一标识
		Path        *string               `gorm:"column:path" json:"path" v:"required"`                                 //path
		Title       *string               `gorm:"column:title;default:''" json:"title" v:"required_unless=type action"` //标题
		Icon        *string               `gorm:"column:icon;default:''" json:"icon" v:"required_if=type menu"`         //图标
		RequireAuth *bool                 `gorm:"column:require_auth;default:1" json:"require_auth"`                    //是否需要权限
		SortNum     *int                  `gorm:"column:sort_num;default:0" json:"sort_num"`                            //排序值

		Children *Permissions `gorm:"-" json:"children,omitempty"`
		Roles    *Roles       `gorm:"many2many:role_has_permissions;" json:"roles,omitempty"`
	}

	Permissions []*Permission
)

func (r Permissions) GetModel(i int) bases.RepositoryModel {
	if len(r) == 0 {
		return &Permission{}
	}
	return r[i]
}

func (r Permissions) Len() int {
	return len(r)
}
