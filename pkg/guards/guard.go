package guards

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/zxdstyle/liey-admin/framework/adm"
)

type Guard interface {
	Can(authID uint, path, action string) error
}

var (
	instance Guard
)

func init() {
	instance = initGuard()
}

func initGuard() Guard {
	val, err := NewCasbinGate(adm.DB())
	if err != nil || val == nil {
		g.Log().Fatal(context.Background(), err)
	}
	return val
}

func GetGuard() Guard {
	return instance
}
