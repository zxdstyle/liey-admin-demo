package enums

type PermissionType string

var (
	PermissionTypeMenu   = PermissionType("menu")
	PermissionTypePage   = PermissionType("page")
	PermissionTypeAction = PermissionType("action")
)

func (c PermissionType) Value() string {
	return string(c)
}

func (c PermissionType) Label() string {
	switch c {
	case PermissionTypeMenu:
		return "菜单"
	case PermissionTypePage:
		return "页面"
	case PermissionTypeAction:
		return "操作"
	default:
		return "Unknown"
	}
}
