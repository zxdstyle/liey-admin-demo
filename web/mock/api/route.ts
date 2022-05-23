import type { MockMethod } from 'vite-plugin-mock';
import { routeModel } from '../model';

const apis: MockMethod[] = [
  {
    url: '/mock/menus',
    method: 'post',
    response: (options: Service.MockOption): Service.MockServiceResult => {
      const routeHomeName: AuthRoute.RouteKey = 'dashboard_workbench';
      return {
        code: 200,
        message: 'ok',
        data: {
          routes: routeModel,
          home: routeHomeName
        }
      };
    }
  }
];

export default apis;
