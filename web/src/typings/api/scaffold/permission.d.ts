/** 后端返回的用户权益相关类型 */
declare namespace Api {
  interface Permission {
    id: number;
    name?: string;
    slug?: string;
    rules?: PermissionRule[];
  }

  interface PermissionRule {
    http_method: HttpMethod[];
    http_path: string;
  }

  type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE';
}
