import Resource from '@/service/api/base/resource';

class Role extends Resource<Api.Role> {
  resource = 'roles';
}

export default new Role();
