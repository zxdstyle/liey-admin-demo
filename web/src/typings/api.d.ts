// 后端接口返回的数据类型

interface ApiResponse {
  code: number;
  message: string;
  data: any;
}

/** 后端返回的路由相关类型 */
declare namespace ApiRoute {
  /** 后端返回的路由数据类型 */
  interface Route {
    /** 动态路由 */
    routes: AuthRoute.Route[];
    /** 路由首页对应的key */
    home: AuthRoute.RouteKey;
  }
}

declare namespace ApiDemo {
  interface DataWithAdapter {
    dataId: string;
    dataName: string;
  }
}
