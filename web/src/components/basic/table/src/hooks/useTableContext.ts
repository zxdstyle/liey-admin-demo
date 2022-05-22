import type { Ref, ComputedRef } from 'vue';
import { provide, inject } from 'vue';
import type { BasicTableProps, TableActionType } from '../types/table';

const key = Symbol('basic-table');

type Instance = TableActionType<any> & {
  wrapRef: Ref<Nullable<HTMLElement>>;
  getBindValues: ComputedRef<Recordable>;
};

type RetInstance = Omit<Instance, 'getBindValues'> & {
  getBindValues: ComputedRef<BasicTableProps<any>>;
};

export function createTableContext(instance: Instance) {
  provide(key, instance);
}

export function useTableContext(): RetInstance {
  return inject(key) as RetInstance;
}
