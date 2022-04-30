package model

import (
	"github.com/zxdstyle/liey-admin/framework/http/bases"
)

type User struct {
	bases.Model
	Name *string `gorm:"column:name" json:"name,omitempty"`
}
