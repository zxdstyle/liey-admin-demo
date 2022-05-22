import { onUnmounted, ref, unref, watch } from 'vue';
import type { WatchStopHandle } from 'vue';
import { getDynamicProps, DynamicProps } from '@/utils';
import type { BasicTableProps, FetchParams, TableActionType } from '../types/table';

// type Props = Partial<DynamicProps<BasicTableProps<T>>>;

type UseTableMethod<T> = TableActionType<T>;

export function useTable<T>(
  tableProps?: Partial<DynamicProps<BasicTableProps<T>>>
): [(instance: TableActionType<T>, formInstance: UseTableMethod<T>) => void, TableActionType<T>] {
  const tableRef = ref<Nullable<TableActionType<T>>>(null);
  const loadedRef = ref<Nullable<boolean>>(false);

  let stopWatch: WatchStopHandle;

  function register(instance: TableActionType<T>) {
    onUnmounted(() => {
      tableRef.value = null;
      loadedRef.value = null;
    });

    if (unref(loadedRef) && instance === unref(tableRef)) return;

    tableRef.value = instance;
    tableProps && instance.setProps(getDynamicProps(tableProps));
    loadedRef.value = true;

    stopWatch?.();

    stopWatch = watch(
      () => tableProps,
      () => {
        tableProps && instance.setProps(getDynamicProps(tableProps));
      },
      {
        immediate: true,
        deep: true
      }
    );
  }

  function getTableInstance(): TableActionType<T> {
    const table = unref(tableRef);
    if (!table) {
      console.error(
        'The table instance has not been obtained yet, please make sure the table is presented when performing the table operation!'
      );
    }
    return table as TableActionType<T>;
  }

  const methods: TableActionType<T> = {
    reload: async (opt?: FetchParams) => {
      await getTableInstance().reload(opt);
    },
    setProps: (props: Partial<BasicTableProps<T>>) => getTableInstance().setProps(props),
    setTableData: values => getTableInstance().setTableData(values),
    setRowData: (row, value) => getTableInstance().setRowData(row, value),
    setFieldData: (row, field, value) => getTableInstance().setFieldData(row, field, value)
  };

  return [register, methods];
}
