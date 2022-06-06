package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/zxdstyle/liey-admin/framework/http/bases"
	"github.com/zxdstyle/liey-admin/framework/support/crypto"
	"gorm.io/gorm"
)

type (
	Admin struct {
		bases.Model
		Username   *string     `json:"username,omitempty" gorm:"not null;type:varchar(64);comment:用户名"`
		Email      *string     `json:"email,omitempty" gorm:"not null;type:varchar(64);unique:uniq_email;comment:邮箱" v:"required,unique-db=admins"`
		Password   *string     `json:"-" gorm:"not null;comment:密码"`
		Avatar     *string     `json:"avatar,omitempty" gorm:"not null;default:'';comment:头像"`
		IsActive   *uint8      `json:"is_active,omitempty" gorm:"not null;type:tinyint;default:1;comment:状态 0：禁用 1：启用"`
		RegisterAt *gtime.Time `json:"register_at,omitempty" gorm:"comment:用户注册时间"`
	}

	Admins []*Admin
)

func (a *Admin) BeforeSave(tx *gorm.DB) error {
	if a.Password != nil {
		hash, err := crypto.PasswordHash(*a.Password)
		if err != nil {
			return err
		}
		a.Password = &hash
	}
	return nil
}

func (r Admins) GetModel(i int) bases.RepositoryModel {
	if len(r) == 0 {
		return &Admin{}
	}
	return r[i]
}

func (r Admins) Len() int {
	return len(r)
}

func (Admin) GuardName() string {
	return "api"
}
