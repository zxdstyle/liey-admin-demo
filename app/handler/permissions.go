package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-scaffold/http/logic"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type Permission struct {
}

func (r Permission) Index(ctx context.Context, req requests.Request) (*responses.Response, error) {
	Permissions := &model.Permissions{}
	if req.NeedPaginate() {
		paginator := req.Paginator(Permissions)
		if err := logic.Permission.Paginate(ctx, req, paginator); err != nil {
			return nil, err
		}
		return responses.Success(paginator), nil
	}

	if err := logic.Permission.All(ctx, req, Permissions); err != nil {
		return nil, err
	}
	return responses.Success(Permissions), nil
}

func (r Permission) Show(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Permission{}
	mo.SetKey(req.ResourceID("permission"))
	if err := logic.Permission.Show(ctx, req.GetWithResources(), mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Permission) Update(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Permission{}
	mo.SetKey(req.ResourceID("permission"))
	if err := req.Parse(mo); err != nil {
		return nil, err
	}

	if err := logic.Permission.Update(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Permission) Create(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Permission{}
	if err := req.Validate(mo); err != nil {
		return nil, err
	}

	if err := logic.Permission.Create(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Permission) Destroy(ctx context.Context, req requests.Request) (*responses.Response, error) {
	if err := logic.Permission.DestroyById(ctx, req.ResourceID("Permission")); err != nil {
		return nil, err
	}
	return responses.NoContent(), nil
}
