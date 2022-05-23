package auth

import (
	"context"
	"fmt"
	"github.com/zxdstyle/liey-admin-demo/app/enums"
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

func (l Logic) UserRoutes(ctx context.Context, req requests.Request, resp *UserRouteResp) error {
	var menus model.Menus
	if err := repository.Menu.All(ctx, req, &menus); err != nil {
		return err
	}

	var userRoutes []*UserRoute
	for _, menu := range menus {
		userRoutes = append(userRoutes, l.transformToRoute(*menu))
		if menu.IsDefault != nil && *menu.IsDefault {
			resp.Home = menu.Name
		}
	}

	refer := make(map[uint]*UserRoute, 0)
	tree := make([]*UserRoute, 0)
	for idx, route := range userRoutes {
		refer[route.ID] = userRoutes[idx]
	}

	for idx, route := range userRoutes {
		pid := *route.ParentId
		if pid == 0 {
			tree = append(tree, userRoutes[idx])
		} else {
			if _, ok := refer[pid]; ok {
				if refer[pid].Children == nil {
					ur := make([]*UserRoute, 0)
					refer[pid].Children = &ur
				}

				*refer[pid].Children = append(*refer[pid].Children, userRoutes[idx])
			}
		}
	}
	l.resolveComponent(&tree)
	resp.Routes = &tree
	return nil
}

func (Logic) transformToRoute(menu model.Menu) *UserRoute {
	return &UserRoute{
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
		},
		Children: nil,
	}
}

func (l Logic) resolveComponent(routes *[]*UserRoute) {
	if routes == nil {
		return
	}

	for idx, route := range *routes {
		if route.Children == nil {
			(*routes)[idx].Component = &enums.RouteComponentSelf
		} else if route.ParentId != nil && *route.ParentId != 0 {
			(*routes)[idx].Component = &enums.RouteComponentMulti
		}
		l.resolveComponent((*routes)[idx].Children)
	}
}
