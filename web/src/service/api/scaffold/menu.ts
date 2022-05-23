import { request } from '@/service/request';
import Resource from '@/service/api/base/resource';

class Menu extends Resource {
  resource = 'menus';

  TreeData = (params = {}) => {
    return request.get(`/api/${this.version}/tree-menus`, { params });
  };
}

export default new Menu();
