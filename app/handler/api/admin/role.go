package admin

import (
	"context"
	"fmt"
	"github.com/zxdstyle/liey-admin-demo/app/logic"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type RoleHandler struct {
	
}

func (r RoleHandler) Index(ctx context.Context, req requests.Request) (*responses.Response, error) {
	return nil, fmt.Errorf("123123")
}

func (r RoleHandler) Show(ctx context.Context, req requests.Request) (*responses.Response, error) {
	role := &model.User{}
	if err := logic.Role.Show(ctx, []string{}, role); err != nil {
		return nil, err
	}
	return responses.Success(role), nil
}

func (r RoleHandler) Update(ctx context.Context, req requests.Request) (*responses.Response, error) {
	panic("implement me")
}

func (r RoleHandler) Create(ctx context.Context, req requests.Request) (*responses.Response, error) {
	panic("implement me")
}

func (r RoleHandler) Destroy(ctx context.Context, req requests.Request) (*responses.Response, error) {
	panic("implement me")
}

