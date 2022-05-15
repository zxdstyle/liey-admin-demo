package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

var Auth = &apiAuth{}

type apiAuth struct {
}

func (apiAuth) Login(ctx context.Context, req requests.Request) (*responses.Response, error) {
	return nil, nil
}
