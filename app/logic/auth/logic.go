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
	var menus model.Menus
	if err := repository.Menu().TreeData(ctx, &menus); err != nil {
		return err
	}
	if menus == nil {
		return nil
	}

	home := "/dashboard/workbench"
	resp.Home = &home
	resp.Routes = l.resolveComponent(&menus)
	return nil
}

// 菜单转换为路由格式
func (Logic) transformToRoute(menu model.Menu) UserRoute {
	return UserRoute{
		ID:        menu.ID,
		Name:      menu.Name,
		Path:      menu.Path,
		Component: &enums.RouteComponentBasic,
		ParentId:  menu.ParentId,
		Meta: &RouteMeta{
			Title:        menu.Title,
			RequiresAuth: menu.RequiresAuth,
			KeepAlive:    menu.Keepalive,
			Icon:         menu.Icon,
			Hide:         menu.Hidden,
			//Href:         menu.Path,
			Order: menu.SortNum,
		},
		Children: nil,
	}
}

func (l Logic) resolveComponent(menus *model.Menus) *[]*UserRoute {
	routes := make([]*UserRoute, 0)
	for idx, menu := range *menus {
		route := l.transformToRoute(*(*menus)[idx])
		if menu.Children == nil {
			route.Component = &enums.RouteComponentSelf
		} else if menu.ParentId != nil && *menu.ParentId != 0 {
			route.Component = &enums.RouteComponentMulti
			route.Children = l.resolveComponent((*menus)[idx].Children)
		} else {
			route.Children = l.resolveComponent((*menus)[idx].Children)
		}
		routes = append(routes, &route)
	}
	return &routes
}
