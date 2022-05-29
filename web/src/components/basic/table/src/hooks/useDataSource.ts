import type { ComputedRef, Ref } from 'vue';
import { computed, onMounted, ref, unref, watch, watchEffect } from 'vue';
import { merge, isBoolean } from 'lodash-es';
import { PAGE_SIZE } from '@/components/basic/table/src/const';
import type { PaginationProps } from '../types/pagination';
import type { BasicTableProps, FetchParams } from '../types/table';
import useSorter from './useSorter';
import useFilter from './useFilter';

interface ActionType {
  getPaginationInfo: ComputedRef<boolean | PaginationProps>;
  setPagination: (info: Partial<PaginationProps>) => void;
  setLoading: (loading: boolean) => void;
  // getFieldsValue: () => Recordable;
  // clearSelectedRowKeys: () => void;
  tableData: Ref<Recordable[]>;
}

export function useDataSource(
  propsRef: ComputedRef<BasicTableProps<any>>,
  { setLoading, tableData, setPagination, getPaginationInfo }: ActionType
) {
  const dataSourceRef = ref<Recordable[]>([]);

  watchEffect(() => {
    tableData.value = unref(dataSourceRef);
  });

  watch(
    () => unref(propsRef).data,
    () => {
      const { data, api } = unref(propsRef);
      !api && data && (dataSourceRef.value = data);
    },
    {
      immediate: true
    }
  );

  async function reload(opt?: FetchParams) {
    await fetch(opt);
  }

  const { handleSorterChange } = useSorter(reload);
  const { handleFilterChange } = useFilter(reload);

  async function fetch(opt?: Recordable) {
    try {
      setLoading(true);

      let pageParams: Recordable = {};

      const { api, pagination, searchInfo } = unref(propsRef);
      if (!api) return;

      const { page = 1, pageSize = PAGE_SIZE } = unref(getPaginationInfo) as PaginationProps;
      if ((isBoolean(pagination) && !pagination) || isBoolean(getPaginationInfo)) {
        pageParams = {};
      } else {
        pageParams.page = (opt && opt.page) || page;
        pageParams.pageSize = pageSize;
      }

      const params: Recordable = merge(searchInfo, pageParams, unref(opt));

      const { data, meta } = await api(params);
      if (!isBoolean(pagination) || pagination) {
        setPagination({ itemCount: meta.pagination?.total });
      }

      dataSourceRef.value = data;
    } finally {
      setLoading(false);
    }
  }

  const getDataSourceRef = computed(() => {
    return unref(dataSourceRef);
  });

  onMounted(() => {
    unref(propsRef).immediate && fetch();
  });

  async function handlePageSizeChange(pageSize: number) {
    setPagination({ pageSize });
    await reload();
  }

  async function handlePageChange(page: number) {
    setPagination({ page });
    await reload();
  }

  function setTableData<T = Recordable>(values: T[]) {
    dataSourceRef.value = values;
  }

  function setRowData<T = Recordable>(row: number, values: T) {
    dataSourceRef.value[row] = values;
  }

  function setFieldData(row: number, field: string, value: any) {
    dataSourceRef.value[row][field] = value;
  }

  return {
    reload,
    setTableData,
    setRowData,
    setFieldData,
    getDataSourceRef,
    handleSorterChange,
    handlePageChange,
    handlePageSizeChange,
    handleFilterChange
  };
}
