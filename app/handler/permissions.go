package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/logic"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

var ApiPermission = &Permission{}

type Permission struct {
}

func (r Permission) Index(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mos := &model.Permissions{}
	if req.NeedPaginate() {
		paginator := req.Paginator(mos)
		if err := logic.Permission().Paginate(ctx, req, paginator); err != nil {
			return nil, err
		}
		return responses.Success(paginator.Data).WithMeta(paginator.Meta), nil
	}

	if err := logic.Permission().All(ctx, req, mos); err != nil {
		return nil, err
	}
	return responses.Success(mos), nil
}

func (r Permission) TreeData(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mos := &model.Permissions{}
	if err := logic.Permission().TreeData(ctx, req, mos); err != nil {
		return nil, err
	}
	return responses.Success(mos), nil
}

func (r Permission) Show(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Permission{}
	mo.SetKey(req.ResourceID("permission"))
	if err := logic.Permission().Show(ctx, req, mo); err != nil {
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

	if err := logic.Permission().Update(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Permission) Create(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Permission{}
	if err := req.Validate(mo); err != nil {
		return nil, err
	}

	if err := logic.Permission().Create(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Permission) Destroy(ctx context.Context, req requests.Request) (*responses.Response, error) {
	if err := logic.Permission().DestroyById(ctx, req.ResourceID("permission")); err != nil {
		return nil, err
	}
	return responses.NoContent(), nil
}
