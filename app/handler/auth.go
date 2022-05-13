package handler

import (
	"context"
	"github.com/zxdstyle/liey-admin-demo/app/jobs"
	"github.com/zxdstyle/liey-admin-scaffold/http/model"
	"github.com/zxdstyle/liey-admin/framework/http/requests"
	"github.com/zxdstyle/liey-admin/framework/http/responses"
	"github.com/zxdstyle/liey-admin/framework/queue"
	"github.com/zxdstyle/liey-admin/framework/support"
)

func LoginByPassword(ctx context.Context, req requests.Request) (*responses.Response, error) {
	adm := model.Admin{}
	adm.SetKey(21)
	token, err := support.JWT().CreateToken(adm)
	if err != nil {
		return nil, err
	}
	return responses.Success(token), nil
}

func UserInfo(ctx context.Context, req requests.Request) (*responses.Response, error) {
	auth := req.ID()
	if err := queue.Dispatch(jobs.SendEmail{}, 12); err != nil {
		return nil, err
	}
	return responses.Success(auth), nil
}
