package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/auth"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

func LoginByPassword(ctx context.Context, req requests.Request) (*responses.Response, error) {
	user := &model.User{}
	guard, er := auth.Guard("web")
	if er != nil {
		return nil, er
	}

	if err := guard.Attempt(user); err != nil {
		return nil, err
	}

	return nil, nil
}
