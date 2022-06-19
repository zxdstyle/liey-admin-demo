package auth

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/zxdstyle/liey-admin-demo/app/enums"
	events2 "github.com/zxdstyle/liey-admin-demo/app/events"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin-demo/app/repository"
	"github.com/zxdstyle/liey-admin/framework/events"
	"github.com/zxdstyle/liey-admin/framework/http/queryBuilder/clauses"
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
	req.AddClauses(clauses.NewPreloadClause("with.roles", "*"))
	if err := repository.Menu().All(ctx, req, &menus); err != nil {
		return err
	}

	if er := repository.Menu().MakeTreeData(&menus); er != nil {
		return er
	}

	authID := req.ID()
	var roles model.Roles
	if err := repository.Role().GetRoles(ctx, authID, &roles); err != nil {
		return err
	}

	roleID := gset.NewIntSet()
	for _, role := range roles {
		roleID.Add(int(role.GetKey()))
	}

	menus = l.cleanUnauthorized(menus, roleID)

	home := "dashboard_workbench"
	resp.Home = &home
	resp.Routes = l.resolveComponent(&menus)
	return nil
}

func (l *Logic) cleanUnauthorized(menus model.Menus, roles *gset.IntSet) model.Menus {
	routes := model.Menus{}
	for i, menu := range menus {
		if menu.Roles == nil {
			continue
		}

		for _, role := range *menu.Roles {
			if roles.Contains(int(role.GetKey())) {
				routes = append(routes, menus[i])
				goto next
			}
		}

	next:

		if menu.Children != nil {
			children := l.cleanUnauthorized(*menu.Children, roles)
			menu.Children = &children
		}
	}
	return routes
}

// 菜单转换为路由格式
func (Logic) transformToRoute(p model.Menu) UserRoute {
	return UserRoute{
		ID:        p.ID,
		Name:      p.Name,
		Path:      p.Path,
		Component: &enums.RouteComponentBasic,
		ParentId:  p.ParentId,
		Meta: &RouteMeta{
			Title:        p.Title,
			RequiresAuth: p.RequiresAuth,
			KeepAlive:    p.Keepalive,
			Icon:         p.Icon,
			Order:        p.SortNum,
			Hide:         p.Hidden,
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
		}

		if menu.Children != nil {
			route.Children = l.resolveComponent((*menus)[idx].Children)
		}

		routes = append(routes, &route)
	}
	return &routes
}
