package model

import (
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Permission struct {
		bases.Model
		Path        *string `json:"path"`
		Method      *string `json:"method"`
		Description *string `json:"description"`
		ParentId    *uint   `json:"parent_id" v:"required"`

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
