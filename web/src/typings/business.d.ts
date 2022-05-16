/** 用户相关模块 */
declare namespace Auth {
  /**
   * 用户角色类型(前端静态路由用角色类型进行路由权限的控制)
   * - super: 超级管理员(该权限具有所有路由数据)
   * - admin: 管理员
   * - user: 用户
   * - custom: 自定义角色
   */
  type RoleType = keyof typeof import('@/enum').EnumUserRole;

  /** 用户信息 */
  interface UserInfo {
    /** 用户id */
    id: number;
    /** 邮箱 */
    email: string;
    /** 用户名 */
    username: string;
    /** 头像 */
    avatar: string;
    /** 注册时间 */
    created_at: string;
  }
}

declare namespace Demo {
  interface DataWithAdapter {
    id: string;
    name: string;
  }
}
