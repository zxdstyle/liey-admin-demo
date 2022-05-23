<template>
  <div ref="wrapRef" class="h-full py-2 flex flex-col">
    <TableHeader>
      <template #left><slot name="left"></slot></template>
      <template #toolbar><slot name="toolbar"></slot></template>
    </TableHeader>
    <div class="h-full">
      <n-data-table ref="tableElRef" v-bind="getBindValues" :style="resizeStyle" />
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref, toRaw, unref } from 'vue';
import { useDataSource } from '@/components/basic/table/src/hooks/useDataSource';
import { usePagination } from '@/components/basic/table/src/hooks/usePagination';
import TableHeader from './components/TableHeader.vue';
import { useLoading } from './hooks/useLoading';
import type { BasicTableProps, TableActionType } from './types/table';
import { createTableContext } from './hooks/useTableContext';

export default defineComponent({
  name: 'BasicTable',
  components: { TableHeader },
  emits: ['register'],
  setup(props, { attrs, emit }) {
    const wrapRef = ref(null);
    const tableElRef = ref(null);
    const tableData = ref<Recordable[]>([]);

    const innerPropsRef = ref<Partial<BasicTableProps<any>>>();

    const getProps = computed(() => {
      return { ...props, ...unref(innerPropsRef) } as BasicTableProps<any>;
    });

    function setProps(props: Partial<BasicTableProps<any>>) {
      innerPropsRef.value = { ...unref(innerPropsRef), ...props };
    }

    const { setLoading, getLoading } = useLoading(getProps);

    const { getPaginationInfo, setPagination } = usePagination(getProps);

    const {
      reload,
      setRowData,
      setTableData,
      setFieldData,
      getDataSourceRef,
      handleSorterChange,
      handlePageChange,
      handlePageSizeChange,
      handleFilterChange
    } = useDataSource(getProps, {
      setLoading,
      tableData,
      setPagination,
      getPaginationInfo
    });

    const getBindValues = computed(() => {
      const data = unref(getDataSourceRef);
      const propsData: Recordable = {
        rowKey: (row: Recordable) => row.id,
        pagination: toRaw(unref(getPaginationInfo)),
        loading: unref(getLoading),
        remote: true,
        tableLayout: 'fixed',
        striped: true,
        flexHeight: true,
        'onUpdate:page': handlePageChange,
        'onUpdate:pageSize': handlePageSizeChange,
        'onUpdate:sorter': handleSorterChange,
        'onUpdate:filters': handleFilterChange,
        ...attrs,
        ...unref(getProps),
        data
      };
      return propsData;
    });

    const tableAction: TableActionType<any> = {
      reload,
      setProps,
      setTableData,
      setRowData,
      setFieldData
    };

    const resizeStyle = computed(() => {
      if (unref(getBindValues).flexHeight) {
        return { height: '100%' };
      }
      return {};
    });

    createTableContext({ ...tableAction, wrapRef, getBindValues });

    emit('register', tableAction);

    return { wrapRef, tableElRef, getBindValues, resizeStyle };
  }
});
</script>

<style lang="less" scoped></style>
