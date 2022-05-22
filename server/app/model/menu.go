package model

import (
	"github.com/zxdstyle/liey-admin/framework/http/bases"
	"gorm.io/gorm"
)

type (
	Menu struct {
		bases.Model
		ParentId  *uint          `gorm:"column:parent_id;comment:父菜单ID" json:"parent_id" v:"required"` //父菜单ID
		Path      *string        `gorm:"column:path;comment:路由path" json:"path" v:"required"`          //路由path
		Title     *string        `gorm:"column:title;comment:标题" json:"title" v:"required"`            //标题
		Name      *string        `gorm:"column:name;comment:路由name" json:"name" v:"required"`          //路由name
		Icon      *string        `gorm:"column:icon;comment:图标" json:"icon"`                           //图标
		Hidden    *bool          `gorm:"default:0;column:hidden;comment:是否在菜单隐藏" json:"hidden"`        //是否在菜单隐藏
		SortNum   *int           `gorm:"default:1;column:sort_num;comment:排序标记" json:"sort_num"`       //排序标记
		Keepalive *bool          `gorm:"default:1;column:keepalive;comment:缓存" json:"keepalive"`       //缓存
		IsDefault *bool          `gorm:"default:0;column:is_default;comment:默认菜单" json:"is_default"`   //默认菜单
		DeletedAt gorm.DeletedAt `json:"-"`

		Children *Menus `gorm:"-" json:"children,omitempty"`
	}

	Menus []*Menu
)

func (m Menus) GetModel(i int) bases.RepositoryModel {
	if len(m) == 0 {
		return &Menu{}
	}
	return m[i]
}