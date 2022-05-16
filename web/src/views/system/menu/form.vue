<template>
  <BasicForm @register="registerForm"></BasicForm>
</template>

<script lang="ts" setup>
import { BasicForm, useForm } from '@/components/basic/form';

const props = defineProps({
  model: {
    type: Object,
    default: () => {}
  }
});

const [registerForm] = useForm({
  model: props.model,
  rowProps: {
    gutter: 24
  },
  schemas: [
    { field: 'parent', label: '父级菜单', component: 'Select', span: 12 },
    { field: 'name', label: '菜单名称', component: 'Input', span: 12 },
    { field: 'slug', label: '唯一标识', component: 'Input', helpMessage: '由客户端定义', span: 8 },
    { field: 'path', label: '路由Path', component: 'Input', span: 16, dynamicDisabled: true },
    { field: 'keepalive', label: 'Keepalive', component: 'Switch', componentProps: { round: false }, span: 6 },
    {
      field: 'require_auth',
      label: '是否公开',
      component: 'Switch',
      helpMessage: '新建的菜单需要绑定权限才可以显示，如果选择公开则无需绑定权限',
      componentProps: { round: false },
      span: 6
    }
  ]
});
</script>

<style lang="less" scoped></style>
