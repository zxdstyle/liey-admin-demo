package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/logic"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

type Admin struct {
}

func (r Admin) Index(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mos := &model.Admins{}
	if req.NeedPaginate() {
		paginator := req.Paginator(mos)
		if err := logic.Admin.Paginate(ctx, req, paginator); err != nil {
			return nil, err
		}
		return responses.Success(paginator.Data).WithMeta(paginator.Meta), nil
	}

	if err := logic.Admin.All(ctx, req, mos); err != nil {
		return nil, err
	}
	return responses.Success(mos), nil
}

func (r Admin) Show(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Admin{}
	mo.SetKey(req.ResourceID("admin"))
	if err := logic.Admin.Show(ctx, req, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Admin) Update(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Admin{}
	if err := req.Parse(mo); err != nil {
		return nil, err
	}

	mo.SetKey(req.ResourceID("admin"))
	if err := logic.Admin.Update(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Admin) Create(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Admin{}
	if err := req.Validate(mo); err != nil {
		return nil, err
	}

	if err := logic.Admin.Create(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Admin) Destroy(ctx context.Context, req requests.Request) (*responses.Response, error) {
	if err := logic.Admin.DestroyById(ctx, req.ResourceID("admin")); err != nil {
		return nil, err
	}
	return responses.NoContent(), nil
}
