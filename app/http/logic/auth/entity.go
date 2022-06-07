package auth

import "github.com/zxdstyle/liey-admin-demo/app/enums"

type (
	LoginByPwd struct {
		Email    *string `json:"email" v:"required,email"`
		Password *string `json:"password" v:"required"`
	}

	LoginResp struct {
		Email  *string `json:"email"`
		Token  *string `json:"token"`
		Avatar *string `json:"avatar"`
	}

	UserRouteResp struct {
		Home   *string       `json:"home"`
		Routes *[]*UserRoute `json:"routes"`
	}

	RouteMeta struct {
		Title        *string `json:"title"`
		RequiresAuth *bool   `json:"requiresAuth"`
		KeepAlive    *bool   `json:"keepAlive"`
		Icon         *string `json:"icon"`
		Hide         *bool   `json:"hide"`
		Href         *string `json:"href,omitempty"`
		Order        *int    `json:"order,omitempty"`
	}

	UserRoute struct {
		ID        uint                  `json:"-"`
		ParentId  *uint                 `json:"-"`
		Name      *string               `json:"name"`
		Path      *string               `json:"path"`
		Component *enums.RouteComponent `json:"component,omitempty"`
		Meta      *RouteMeta            `json:"meta"`
		Children  *[]*UserRoute         `json:"children,omitempty"`
	}
)
