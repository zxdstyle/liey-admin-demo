package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/logic"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type Role struct {
}

func (r Role) Index(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mos := &model.Roles{}
	if req.NeedPaginate() {
		paginator := req.Paginator(mos)
		if err := logic.Role().Paginate(ctx, req, paginator); err != nil {
			return nil, err
		}
		return responses.Success(paginator.Data).WithMeta(paginator.Meta), nil
	}

	if err := logic.Role().All(ctx, req, mos); err != nil {
		return nil, err
	}
	return responses.Success(mos), nil
}

func (r Role) Show(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Role{}
	mo.SetKey(req.ResourceID("role"))
	if err := logic.Role().Show(ctx, req, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Role) Update(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Role{}
	if err := req.Parse(mo); err != nil {
		return nil, err
	}

	mo.SetKey(req.ResourceID("role"))
	if err := logic.Role().Update(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Role) Create(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Role{}
	if err := req.Validate(mo); err != nil {
		return nil, err
	}

	if err := logic.Role().Create(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Role) Destroy(ctx context.Context, req requests.Request) (*responses.Response, error) {
	if err := logic.Role().DestroyById(ctx, req.ResourceID("role")); err != nil {
		return nil, err
	}
	return responses.NoContent(), nil
}
