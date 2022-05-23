package enums

type RouteComponent string

var (
	RouteComponentBasic = RouteComponent("basic") //  基础布局，具有公共部分的布局
	RouteComponentBlank = RouteComponent("blank") // 空白布局
	RouteComponentMulti = RouteComponent("multi") // 多级路由布局(三级路由或三级以上时，除第一级路由和最后一级路由，其余的采用该布局)
	RouteComponentSelf  = RouteComponent("self")  // 作为子路由，使用自身的布局(作为最后一级路由，没有子路由)
)

func (c RouteComponent) Value() string {
	return string(c)
}

func (c RouteComponent) Label() string {
	switch c {
	case RouteComponentBasic:
		return "基础布局"
	case RouteComponentBlank:
		return "空白布局"
	case RouteComponentMulti:
		return "多级路由布局"
	case RouteComponentSelf:
		return "普通布局"
	default:
		return "Unknown"
	}
}
