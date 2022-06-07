package main

import (
	"github.com/zxdstyle/liey-admin-demo/app/http"
	_ "github.com/zxdstyle/liey-admin-demo/routes"
	"github.com/zxdstyle/liey-admin/framework/adm"
)

//go:generate go build -o adm main.go
//go:generate go build -o cli cli.go
func main() {
	adm.Start(http.Kernel{})
}
