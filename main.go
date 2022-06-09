package main

import (
	"github.com/zxdstyle/liey-admin-demo/app"
	"github.com/zxdstyle/liey-admin/framework/adm"
)

//go:generate go build -o adm main.go
func main() {
	adm.Start(app.Kernel{})
}
