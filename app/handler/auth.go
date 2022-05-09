package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

func LoginByPassword(ctx context.Context, req requests.Request) (*responses.Response, error) {
	return responses.Success("t"), nil
}

func UserInfo(ctx context.Context, req requests.Request) (*responses.Response, error) {
	auth := req.ID()

	return responses.Success(auth), nil
}
