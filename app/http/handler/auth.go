package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/http/logic"
	"github.com/zxdstyle/liey-admin-demo/app/http/logic/auth"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

var Auth = &apiAuth{}

type apiAuth struct {
}

func (apiAuth) Login(ctx context.Context, req requests.Request) (*responses.Response, error) {
	var entity auth.LoginByPwd
	if err := req.Validate(&entity); err != nil {
		return nil, err
	}

	resp, err := logic.Auth.Login(ctx, entity)
	if err != nil {
		return nil, err
	}

	return responses.Success(resp), nil
}

func (apiAuth) Userinfo(ctx context.Context, req requests.Request) (*responses.Response, error) {
	resp, err := logic.Auth.Userinfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return responses.Success(resp), nil
}

func (apiAuth) UserRoutes(ctx context.Context, req requests.Request) (*responses.Response, error) {
	var resp auth.UserRouteResp
	if err := logic.Auth.UserRoutes(ctx, req, &resp); err != nil {
		return nil, err
	}
	return responses.Success(resp), nil
}
