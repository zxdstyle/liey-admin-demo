package model

import (
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Permission struct {
		bases.Model
		Name     *string           `gorm:"not null;type:varchar(64);comment:名称" json:"name" v:"required"`
		Slug     *string           `gorm:"not null;type:varchar(64);unique;comment:标识" json:"slug" v:"required,unique-db=permissions"`
		Rules    *[]PermissionRule `gorm:"not null;serializer:json;comment:权限规则" json:"rules" v:"required,dive"`
		ParentId *uint             `gorm:"default:0;not null;comment:父级权限" json:"parent_id" v:"required"`
		SortNum  *int              `gorm:"default:0;not null;comment:排序值" json:"sort_num"`

		Children *Permissions `gorm:"-" json:"children,omitempty"`
		Roles    *Roles       `gorm:"many2many:role_has_permissions;" json:"roles,omitempty"`
	}

	Permissions []*Permission

	PermissionRule struct {
		HttpMethods *[]string `json:"http_methods" v:"required"`
		HttpPath    *string   `json:"http_path" v:"required"`
	}

	PermissionRules []PermissionRule
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
