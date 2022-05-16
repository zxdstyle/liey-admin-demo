<template>
  <n-card class="rounded-xl">
    <BasicTable @register="registerTable">
      <template #toolbar>
        <n-button type="primary" @click="openModal"><Icon icon="ph:plus-bold"></Icon>新增菜单</n-button>
      </template>
    </BasicTable>

    <BasicModal style="width: 1000px" @register="registerModal">
      <MenuForm></MenuForm>
    </BasicModal>
  </n-card>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import { Icon } from '@iconify/vue';
import { BasicTable, useTable } from '@/components/basic/table';
import { BasicModal, useModal } from '@/components/basic/modal';
import MenuForm from './form.vue';

defineProps({});

const [registerModal, { openModal }] = useModal();

const [registerTable] = useTable({
  columns: [{ key: 'id', title: 'ID' }]
});

onMounted(() => openModal());
</script>

<style lang="less" scoped></style>
