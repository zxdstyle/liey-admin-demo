import { request } from '@/service/request';

export default class Resource<T> {
  resource = '';

  version = 'v1';

  Index = (params = {}) => {
    return request.get<T[]>(`/api/${this.version}/${this.resource}`, { params });
  };

  Show = (id: number) => {
    return request.get<T>(`/api/${this.version}/${this.resource}/${id}`);
  };

  Update = (id: number, data = {}) => {
    return request.put<T>(`/api/${this.version}/${this.resource}/${id}`, data);
  };

  Create = (data: any) => {
    return request.post<T>(`/api/${this.version}/${this.resource}`, data);
  };

  Destroy = (id: number) => {
    return request.delete(`/api/${this.version}/${this.resource}/${id}`);
  };
}
