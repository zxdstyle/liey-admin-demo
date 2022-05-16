package auth

import (
	"context"
	"fmt"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin-demo/app/repository"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/support"
	"github.com/zxdstyle/liey-admin/framework/support/crypto"
	"gorm.io/gorm"
)

type Logic struct {
}

func NewLogic() *Logic {
	return &Logic{}
}

func (Logic) Login(ctx context.Context, req LoginByPwd) (*LoginResp, error) {
	admin := model.Admin{Email: req.Email}
	if err := repository.Admin.First(ctx, &admin); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("账号或密码错误")
		}
		return nil, err
	}

	if !crypto.PasswordVerify(*req.Password, *admin.Password) {
		return nil, fmt.Errorf("账号或密码错误")
	}

	token, err := support.JWT().CreateToken(admin)
	if err != nil {
		return nil, err
	}

	return &LoginResp{
		Email:  admin.Email,
		Avatar: admin.Avatar,
		Token:  &token,
	}, nil
}

func (Logic) Userinfo(ctx context.Context, req requests.Request) (*model.Admin, error) {
	mo := &model.Admin{}
	mo.SetKey(req.ID())
	if err := repository.Admin.First(ctx, mo); err != nil {
		return nil, err
	}
	return mo, nil
}
