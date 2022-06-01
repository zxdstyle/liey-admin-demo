<template>
  <n-card class="rounded-xl h-full">
    <BasicTable @register="registerTable">
      <template #toolbar>
        <n-button type="primary" @click="openCreateModal"><Icon icon="ph:plus-bold"></Icon>新增权限</n-button>
      </template>
    </BasicTable>

    <BasicModal style="width: 1000px" @register="registerModal">
      <RoleForm v-bind="getFormBind"></RoleForm>
    </BasicModal>
  </n-card>
</template>

<script lang="tsx" setup>
import { computed, onActivated, reactive } from 'vue';
import { Icon } from '@iconify/vue';
import { TableActionHandler } from '@/components/basic/table';
import ApiPermission from '@/service/api/scaffold/permission';
import { useModal } from '@/components/basic/modal';
import useBasicDialog from '@/hooks/common/useDialog';
import RoleForm from './form.vue';
import { useIndex } from './hooks';

const { warning } = useBasicDialog();
const [registerModal, { openModal, setModalProps, closeModal }] = useModal();

const state = reactive({
  model: {}
});

const openEditModal = (row: Api.Permission) => {
  state.model = row;
  setModalProps({ title: '编辑信息' });
  openModal();
};

const openCreateModal = () => {
  state.model = {};
  openModal();
  setModalProps({ title: '新增权限' });
};

const actionHandler: TableActionHandler<Api.Permission> = async (action, row) => {
  switch (action) {
    case 'edit':
      openEditModal(row);
      break;
    case 'delete':
      warning('是否确认删除该权限？', () => handleDelete(row.id));
      break;
    default:
  }
};

const { registerTable, reload } = useIndex(actionHandler);

async function handleDelete(key: number) {
  await ApiPermission.Destroy(key);
  await reload();
}

const getFormBind = computed(() => {
  return {
    model: state.model,
    onSubmit: () => {
      closeModal();
      reload();
    }
  };
});

onActivated(() => reload());
</script>

<style lang="less" scoped></style>
