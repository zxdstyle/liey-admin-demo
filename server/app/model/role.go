package model

import "github.com/zxdstyle/liey-admin/framework/http/bases"

type (
	Role struct {
		bases.Model
		Name *string `gorm:"not null;type:varchar(64);comment:角色名称" json:"name,omitempty" v:"required"`
		Slug *string `gorm:"not null;type:varchar(32);unique;comment:角色标识" json:"slug,omitempty" v:"required,unique-db=roles"`

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

func (r Roles) Len() int {
	return len(r)
}
