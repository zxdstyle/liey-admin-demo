import { App } from 'vue';
import XEUtils from 'xe-utils';
import zhCN from 'vxe-table/es/locale/lang/zh-CN';
import { VXETable, Table, Column } from 'vxe-table';

export default function setupVexTable(app: App) {
  VXETable.setup({
    i18n(key: string, args?: any): string {
      return XEUtils.toFormatString(XEUtils.get(zhCN, key), args);
    },
    table: {
      showHeader: true
    }
  });

  app.use(Table).use(Column);
}
