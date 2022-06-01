import Resource from '@/service/api/base/resource';
import { request } from '@/service/request';

class Permission extends Resource<Api.Permission> {
  resource = 'permissions';

  TreeData = (params = {}) => {
    return request.get(`/api/${this.version}/tree-permissions`, { params });
  };
}

export default new Permission();
