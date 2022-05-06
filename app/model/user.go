package model

import "github.com/zxdstyle/liey-admin/framework/http/bases"

type User struct {
	bases.Model
	Username string
	Password string
}
