import { request } from '@/service/request';
import Resource from '@/service/api/base/resource';

class Menu extends Resource {
  resource = 'menus';

  TreeData = () => {
    return request.get(`/api/${this.version}/tree-menus`);
  };
}

export default new Menu();
