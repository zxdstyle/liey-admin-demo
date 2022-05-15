package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-scaffold/http/logic"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type Role struct {
}

func (r Role) Index(ctx context.Context, req requests.Request) (*responses.Response, error) {
	roles := &model.Roles{}
	if req.NeedPaginate() {
		paginator := req.Paginator(roles)
		if err := logic.Role.Paginate(ctx, req, paginator); err != nil {
			return nil, err
		}
		return responses.Success(paginator), nil
	}

	if err := logic.Role.All(ctx, req, roles); err != nil {
		return nil, err
	}
	return responses.Success(roles), nil
}

func (r Role) Show(ctx context.Context, req requests.Request) (*responses.Response, error) {
	role := &model.Role{}
	role.SetKey(req.ResourceID("role"))
	if err := logic.Role.Show(ctx, req.GetWithResources(), role); err != nil {
		return nil, err
	}
	return responses.Success(role), nil
}

func (r Role) Update(ctx context.Context, req requests.Request) (*responses.Response, error) {
	role := &model.Role{}
	if err := req.Parse(role); err != nil {
		return nil, err
	}

	role.SetKey(req.ResourceID("role"))
	if err := logic.Role.Update(ctx, role); err != nil {
		return nil, err
	}
	return responses.Success(role), nil
}

func (r Role) Create(ctx context.Context, req requests.Request) (*responses.Response, error) {
	role := &model.Role{}
	if err := req.Validate(role); err != nil {
		return nil, err
	}

	if err := logic.Role.Create(ctx, role); err != nil {
		return nil, err
	}
	return responses.Success(role), nil
}

func (r Role) Destroy(ctx context.Context, req requests.Request) (*responses.Response, error) {
	if err := logic.Role.DestroyById(ctx, req.ResourceID("role")); err != nil {
		return nil, err
	}
	return responses.NoContent(), nil
}
