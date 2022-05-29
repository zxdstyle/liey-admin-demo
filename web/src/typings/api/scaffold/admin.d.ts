/** 后端返回的用户权益相关类型 */
declare namespace Api {
  interface Admin {
    id: number;
    username?: string;
    email?: string;
    avatar?: string;
    is_active: boolean;
    register_at?: string;
    created_at?: string;
    updated_at?: string;
  }
}
