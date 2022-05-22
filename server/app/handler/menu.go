package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/logic"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

var ApiMenu = &Menu{}

type Menu struct {
}

func (r Menu) Index(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mos := &model.Menus{}
	if req.NeedPaginate() {
		paginator := req.Paginator(mos)
		if err := logic.Menu.Paginate(ctx, req, paginator); err != nil {
			return nil, err
		}
		return responses.Success(paginator), nil
	}

	if err := logic.Menu.All(ctx, req, mos); err != nil {
		return nil, err
	}
	return responses.Success(mos), nil
}

func (r Menu) TreeData(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mos := &model.Menus{}
	if err := logic.Menu.TreeData(ctx, req, mos); err != nil {
		return nil, err
	}
	return responses.Success(mos), nil
}

func (r Menu) Show(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Menu{}
	mo.SetKey(req.ResourceID("menu"))
	if err := logic.Menu.Show(ctx, req, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Menu) Update(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Menu{}
	if err := req.Parse(mo); err != nil {
		return nil, err
	}

	mo.SetKey(req.ResourceID("menu"))
	if err := logic.Menu.Update(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Menu) Create(ctx context.Context, req requests.Request) (*responses.Response, error) {
	mo := &model.Menu{}
	if err := req.Validate(mo); err != nil {
		return nil, err
	}

	if err := logic.Menu.Create(ctx, mo); err != nil {
		return nil, err
	}
	return responses.Success(mo), nil
}

func (r Menu) Destroy(ctx context.Context, req requests.Request) (*responses.Response, error) {
	if err := logic.Menu.DestroyById(ctx, req.ResourceID("menu")); err != nil {
		return nil, err
	}
	return responses.NoContent(), nil
}
