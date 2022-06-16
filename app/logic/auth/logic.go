package auth

import (
	"context"
	"fmt"
	"github.com/zxdstyle/liey-admin-demo/app/enums"
	events2 "github.com/zxdstyle/liey-admin-demo/app/events"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin-demo/app/repository"
	"github.com/zxdstyle/liey-admin/framework/events"
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
	if err := repository.Admin().First(ctx, &admin); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("账号或密码错误")
		}
		return nil, err
	}

	if !crypto.PasswordVerify(*req.Password, *admin.Password) {
		return nil, fmt.Errorf("账号或密码错误")
	}

	events.Dispatch(ctx, events2.NewUserLogin(admin))

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
	if err := repository.Admin().First(ctx, mo); err != nil {
		return nil, err
	}
	return mo, nil
}

func (l Logic) UserRoutes(ctx context.Context, req requests.Request, resp *UserRouteResp) error {
	var permissions model.Permissions
	if err := repository.Permission().GetRoutes(ctx, &permissions); err != nil {
		return err
	}
	if permissions == nil {
		return nil
	}

	if err := repository.Permission().TreeData(ctx, &permissions); err != nil {
		return err
	}

	home := "/dashboard/workbench"
	resp.Home = &home
	resp.Routes = l.resolveComponent(&permissions)
	return nil
}

// 菜单转换为路由格式
func (Logic) transformToRoute(p model.Permission) UserRoute {
	r := UserRoute{
		ID:        p.ID,
		Name:      p.Slug,
		Path:      p.Path,
		Component: &enums.RouteComponentBasic,
		ParentId:  p.ParentId,
		Meta: &RouteMeta{
			Title:        p.Title,
			RequiresAuth: p.RequireAuth,
			KeepAlive:    p.Keepalive,
			Icon:         p.Icon,
			Order:        p.SortNum,
			Hide:         new(bool),
		},
		Children: nil,
	}

	if *p.Type != enums.PermissionTypeMenu {
		r.Meta.Hide = &enums.True
	}

	return r
}

func (l Logic) resolveComponent(menus *model.Permissions) *[]*UserRoute {
	routes := make([]*UserRoute, 0)
	for idx, menu := range *menus {
		route := l.transformToRoute(*(*menus)[idx])
		if *menu.Type == enums.PermissionTypePage || menu.Children == nil || !l.containMenu(*menu.Children) {
			route.Component = &enums.RouteComponentSelf
		} else if menu.ParentId != nil && *menu.ParentId != 0 {
			route.Component = &enums.RouteComponentMulti
		}

		if menu.Children != nil {
			route.Children = l.resolveComponent((*menus)[idx].Children)
		}

		routes = append(routes, &route)
	}
	return &routes
}

func (l Logic) containMenu(permissions model.Permissions) bool {
	for _, permission := range permissions {
		if *permission.Type == enums.PermissionTypeMenu {
			return true
		}
	}
	return false
}
