package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin/framework/adm"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
)

func Routes(ctx context.Context, req requests.Request) (*responses.Response, error) {
	routes := adm.Server().Routes()
	return responses.Success(routes), nil
}
