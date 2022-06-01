<template>
  <BasicForm @register="registerForm" />
</template>

<script lang="ts" setup>
import { PropType } from 'vue';
import ApiRole from '@/service/api/scaffold/role';
import { BasicForm, useForm } from '@/components/basic/form';

const props = defineProps({
  model: {
    type: Object as PropType<Api.Role>,
    default: () => {}
  }
});

const emit = defineEmits(['submit']);

const [registerForm, { resetFields }] = useForm({
  model: props.model,
  rowProps: { gutter: 24 },
  schemas: [
    { field: 'name', label: '角色名称', component: 'Input', span: 20 },
    { field: 'slug', label: '角色标识', component: 'Input', span: 20 }
  ],
  rules: {
    name: { required: true, trigger: 'blur', message: '请输入角色名称' },
    slug: { required: true, trigger: 'blur', message: '请输入角色标识' }
  },
  onSubmit: async e => {
    let res;
    if (e.id && e.id > 0) {
      res = await ApiRole.Update(e.id, e);
    } else {
      res = await ApiRole.Create(e);
    }
    await resetFields();
    emit('submit', res);
  }
});
</script>

<style lang="less" scoped></style>
