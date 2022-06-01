/** 后端返回的用户权益相关类型 */
declare namespace Api {
  interface Permission {
    id: number;
    name: string;
    slug: string;
    parent_id: number;
    sort_num: number;
    rules: PermissionRule[];

    children?: Permission[];
  }

  interface PermissionRule {
    http_methods: string[];
    http_path: string;
  }
}
