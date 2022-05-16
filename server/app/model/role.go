package model

import "github.com/zxdstyle/liey-admin/framework/http/bases"

type (
	Role struct {
		bases.Model
		Name    *string `gorm:"not null;type:varchar(64);unique;comment:角色名称" json:"name,omitempty" v:"required,unique-db=roles"`
		Status  *int    `gorm:"not null;type:tinyint;default:1;comment:状态 0：禁用 1：启用" json:"status,omitempty"`
		SortNum *int    `gorm:"not null;type:int;default:1;comment:排序值" json:"sort_num,omitempty"`

		Permissions *Permissions `gorm:"many2many:role_has_permissions;" json:"permissions,omitempty"`
	}

	Roles []*Role
)

func (r Roles) GetModel(i int) bases.RepositoryModel {
	if len(r) == 0 {
		return &Role{}
	}
	return r[i]
}
