declare namespace Api {
  interface Role {
    id: number;
    name: string;
    slug: string;
    created_at: string;
    updated_at: string;

    permissions?: Permission[] | number[];
  }
}
