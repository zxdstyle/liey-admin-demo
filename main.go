package main

import (
	"github.com/zxdstyle/liey-admin-demo/app"
	_ "github.com/zxdstyle/liey-admin-demo/app"
	_ "github.com/zxdstyle/liey-admin-demo/routes"
	"github.com/zxdstyle/liey-admin/framework/adm"
)

func main() {
	_ = adm.RegisterKernel(&app.Kernel{})

	adm.Start()
}
