import type { DataTableProps, DropdownDividerOption, DropdownRenderOption } from 'naive-ui';
import { DropdownGroupOption, DropdownOption } from 'naive-ui';
import { TableBaseColumn, TableColumn } from 'naive-ui/es/data-table/src/interface';
import { VueNode } from '@/utils/common/propType';

export interface FetchParams {
  searchInfo?: Recordable;
  page?: number;
  sortInfo?: Recordable;
  filterInfo?: Recordable;
}

export type BasicColumn<T> = TableColumn<T> | BasicTableColumn<T>;

export interface BasicTableColumn<T> extends TableBaseColumn<T> {
  // 业务控制是否显示
  ifShow?: boolean | ((column: BasicColumn<T>) => boolean);
}

export interface BasicTableProps<T> extends DataTableProps {
  api?: (...arg: any) => Promise<any>;
  // 列配置
  columns: BasicColumn<T>[];

  // 额外的请求参数
  searchInfo?: Recordable;

  // 是否立即请求远程数据
  immediate?: boolean;
}

export interface TableActionType<T> {
  setProps: (props: Partial<BasicTableProps<T>>) => void;
  reload: (opt?: FetchParams) => Promise<void>;
  setTableData: (values: Recordable[]) => void;
  setRowData: (row: number, value: Recordable) => void;
  setFieldData: (row: number, field: string, value: any) => void;
}

export type TableActionHandler<T> = (action: string, row: T) => void;

export type TableActionOption = Array<
  DropdownOption | DropdownGroupOption | DropdownDividerOption | DropdownRenderOption
>;

export interface BatchOperation {
  key: string;
  label: string;
  icon: VueNode;
  on: (keys: Array<string | number>) => void;
}
