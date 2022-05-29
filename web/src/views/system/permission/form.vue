<template>
  <BasicForm @register="registerForm"></BasicForm>
</template>

<script lang="ts" setup>
import ApiPermission from '@/service/api/scaffold/permission';
import { BasicForm, useForm } from '@/components/basic/form';

const props = defineProps({
  model: {
    type: Object,
    default: () => {}
  }
});

const emit = defineEmits(['submit']);

const [registerForm, { resetFields }] = useForm({
  model: props.model,
  rowProps: {
    gutter: 24
  },
  schemas: [
    { field: 'name', label: '权限名称', component: 'Input', span: 20 },
    { field: 'slug', label: '权限唯一标识', component: 'Input', span: 20 }
  ],
  rules: {
    name: { required: true, trigger: 'blur', message: '请输入权限名称' },
    slug: { required: true, trigger: 'blur', message: '请输入权限唯一标识' }
  },
  onSubmit: async e => {
    let res;
    if (e.id && e.id > 0) {
      res = await ApiPermission.Update(e.id, e);
    } else {
      res = await ApiPermission.Create(e);
    }
    await resetFields();
    emit('submit', res);
  }
});
</script>

<style lang="less" scoped></style>
