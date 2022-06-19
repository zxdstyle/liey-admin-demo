package model

import (
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type (
	Menu struct {
		bases.Model
		ParentId     *uint   `gorm:"not null;column:parent_id;comment:父菜单ID" json:"parent_id" v:"required"`       //父菜单ID
		Path         *string `gorm:"not null;column:path;comment:路由path" json:"path" v:"required"`                //路由path
		Title        *string `gorm:"not null;column:title;comment:标题" json:"title" v:"required"`                  //标题
		Name         *string `gorm:"not null;column:name;comment:路由name" json:"name" v:"required"`                //路由name
		Icon         *string `gorm:"not null;column:icon;comment:图标" json:"icon" v:"required"`                    //图标
		RequiresAuth *bool   `gorm:"not null;default:1;column:requires_auth;comment:是否需要权限" json:"requires_auth"` // 是否需要权限
		Hidden       *bool   `gorm:"not null;default:0;column:hidden;comment:是否在菜单隐藏" json:"hidden"`              //是否在菜单隐藏
		SortNum      *int    `gorm:"not null;default:1;column:sort_num;comment:排序标记" json:"sort_num"`             //排序标记
		Keepalive    *bool   `gorm:"not null;default:1;column:keepalive;comment:缓存" json:"keepalive"`             //缓存
		IsDefault    *bool   `gorm:"not null;default:0;column:is_default;comment:默认菜单" json:"is_default"`         //默认菜单

		Children *Menus `gorm:"-" json:"children,omitempty"`
		Roles    *Roles `gorm:"many2many:role_has_menus" json:"roles,omitempty"`
	}

	Menus []*Menu
)

func (m Menus) GetModel(i int) bases.RepositoryModel {
	if len(m) == 0 {
		return &Menu{}
	}
	return m[i]
}

func (m Menus) Len() int {
	return len(m)
}
