/** 后端返回的用户权益相关类型 */
declare namespace Api {
  interface Menu {
    id: number;
    name?: string;
    path?: string;
    parent_id?: number;
    title: string;
    hidden?: boolean;
    icon?: string;
    is_default?: boolean;
    sort_num?: number;
    created_at?: string;
    updated_at?: string;
    children?: Menu[];
  }
}
