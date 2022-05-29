import Resource from '@/service/api/base/resource';

class Admin extends Resource<Api.Admin> {
  resource = 'admins';
}

export default new Admin();
