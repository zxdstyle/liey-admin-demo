package jobs

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/queue/serializer"
)

type SendEmail struct {
	serializer.UintSerializer
}

func (job SendEmail) Concurrency() int {
	return 100
}

func (job SendEmail) Handle(ctx context.Context, payload []byte) error {
	var p uint
	if err := job.UnSerialize(payload, &p); err != nil {
		return err
	}

	g.Log().Notice(ctx, p)
	return nil
}
