import { request } from '@/service/request';
import Resource from '@/service/api/base/resource';

class Menu extends Resource<Api.Menu> {
  resource = 'menus';

  TreeData = (params = {}) => {
    return request.get(`/api/${this.version}/tree-menus`, { params });
  };
}

export default new Menu();
