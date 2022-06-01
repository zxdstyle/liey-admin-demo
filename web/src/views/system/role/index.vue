<template>
  <n-card class="rounded-xl h-full">
    <BasicTable @register="registerTable">
      <template #toolbar>
        <n-button type="primary" @click="openCreateModal"><Icon icon="ph:plus-bold" />新增角色</n-button>
      </template>
    </BasicTable>

    <BasicModal style="width: 700px" @register="registerModal">
      <RoleForm v-bind="getFormBind" />
    </BasicModal>
  </n-card>
</template>

<script lang="tsx" setup>
import { computed, onActivated, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { Icon } from '@iconify/vue';
import { TableActionHandler } from '@/components/basic/table';
import ApiRole from '@/service/api/scaffold/role';
import { useModal } from '@/components/basic/modal';
import useBasicDialog from '@/hooks/common/useDialog';
import RoleForm from './form.vue';
import { useIndex } from './hooks';

const [registerModal, { openModal, setModalProps, closeModal }] = useModal();

const state = reactive({
  model: {}
});

const router = useRouter();
const { warning } = useBasicDialog();
const openEditModal = (row: Api.Role) => {
  state.model = row;
  setModalProps({ title: '编辑信息' });
  openModal();
};

const openCreateModal = () => {
  state.model = {};
  openModal();
  setModalProps({ title: '新增角色' });
};

const actionHandler: TableActionHandler<Api.Role> = async (action, row) => {
  switch (action) {
    case 'edit':
      openEditModal(row);
      break;
    case 'detail':
      router.push(`/system/role/${row.id}`);
      break;
    case 'delete':
      warning('是否确认删除该角色？', () => handleDelete(row.id));
      break;
    default:
  }
};

const { registerTable, reload } = useIndex(actionHandler);

async function handleDelete(key: number) {
  await ApiRole.Destroy(key);
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
