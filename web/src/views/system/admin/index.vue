<template>
  <n-card class="rounded-xl h-full">
    <BasicTable @register="registerTable">
      <template #toolbar>
        <n-button type="primary" @click="openCreateModal"><Icon icon="ph:plus-bold"></Icon>新增管理员</n-button>
      </template>
    </BasicTable>

    <BasicModal style="width: 700px" @register="registerModal">
      <AdminForm v-bind="getFormBind"></AdminForm>
    </BasicModal>
  </n-card>
</template>

<script lang="tsx" setup>
import { computed, onActivated, reactive } from 'vue';
import { Icon } from '@iconify/vue';
import { TableActionHandler } from '@/components/basic/table';
import ApiAdmin from '@/service/api/scaffold/admin';
import { useModal } from '@/components/basic/modal';
import useBasicDialog from '@/hooks/common/useDialog';
import AdminForm from './form.vue';
import { useIndex } from './hooks';

const { warning } = useBasicDialog();
const [registerModal, { openModal, setModalProps, closeModal }] = useModal();

const state = reactive({
  model: {}
});

const openEditModal = (row: Api.Admin) => {
  state.model = row;
  setModalProps({ title: '编辑信息' });
  openModal();
};

const openCreateModal = () => {
  state.model = {};
  openModal();
  setModalProps({ title: '新增管理员' });
};

const actionHandler: TableActionHandler<Api.Admin> = async (action, row) => {
  switch (action) {
    case 'edit':
      openEditModal(row);
      break;
    case 'delete':
      warning('是否确定删除该管理员？', () => handleDelete(row.id));
      break;
    default:
  }
};

const { registerTable, reload } = useIndex(actionHandler);

async function handleDelete(key: number) {
  await ApiAdmin.Destroy(key);
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
