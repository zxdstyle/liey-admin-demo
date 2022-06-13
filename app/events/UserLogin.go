package events

import "github.com/zxdstyle/liey-admin-demo/app/model"

type UserLogin struct {
	user model.Admin
}

func NewUserLogin(user model.Admin) UserLogin {
	return UserLogin{user}
}

func (u UserLogin) Payload() interface{} {
	return u.user
}
