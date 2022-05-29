import Resource from '@/service/api/base/resource';

class Permission extends Resource<Api.Permission> {
  resource = 'permissions';
}

export default new Permission();
