import { request } from '@/service/request';

export default class Resource {
  resource = '';

  version = 'v1';

  Index = () => {
    return request.get(`/api/${this.version}/${this.resource}`);
  };

  Show = (id: number) => {
    return request.get(`/api/${this.version}/${this.resource}/${id}`);
  };

  Update = (id: number, data = {}) => {
    return request.put(`/api/${this.version}/${this.resource}/${id}`, data);
  };

  Create = (data: any) => {
    return request.post(`/api/${this.version}/${this.resource}`, data);
  };

  Destroy = (id: number) => {
    return request.delete(`/api/${this.version}/${this.resource}/${id}`);
  };
}
