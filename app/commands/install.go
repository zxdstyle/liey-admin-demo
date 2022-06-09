package commands

import (
	"github.com/spf13/cobra"
	"github.com/zxdstyle/liey-admin-demo/app/model"
	"github.com/zxdstyle/liey-admin/framework/adm"
)

var InstallCommand = &cobra.Command{
	Use:   "install",
	Short: "install the framework",
	Run: func(cmd *cobra.Command, args []string) {
		adm.DB().AutoMigrate(
			&model.Admin{},
			&model.Menu{},
			&model.Permission{},
			&model.Role{},
			&model.RoleHasPermission{},
			&model.User{},
		)

	},
}
