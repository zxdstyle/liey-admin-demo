<template>
  <BasicForm @register="registerForm"></BasicForm>
</template>

<script lang="ts" setup>
import ApiAdmin from '@/service/api/scaffold/admin';
import { BasicForm, useForm } from '@/components/basic/form';

const props = defineProps({
  model: {
    type: Object,
    default: () => {}
  }
});

const emit = defineEmits(['submit']);

const formatTreeData = (data: Api.Menu[]): Api.Menu[] => {
  return [
    {
      id: 0,
      title: '顶级菜单',
      children: data
    }
  ];
};

const [registerForm, { resetFields }] = useForm({
  model: props.model,
  rowProps: {
    gutter: 24
  },
  schemas: [
    { field: 'username', label: '用户名', component: 'Input', span: 20 },
    { field: 'email', label: '邮箱', component: 'Input', span: 24 },
    { field: 'avatar', label: '头像', component: 'Upload', span: 24 },
    {
      field: 'password',
      label: '密码',
      component: 'Input',
      span: 20,
      componentProps: { type: 'password', placeholder: '请输入初始密码，默认与邮箱一致' }
    },
    { field: 'is_active', label: '是否启用', component: 'Switch', span: 16 }
  ],
  rules: {
    username: { required: true, trigger: 'blur', message: '请选择账户名称' },
    email: [
      { required: true, trigger: 'blur', message: '请输入邮箱' },
      { type: 'email', message: '邮箱格式错误' }
    ]
  },
  onSubmit: async e => {
    let res;
    if (e.id && e.id > 0) {
      res = await ApiAdmin.Update(e.id, e);
    } else {
      res = await ApiAdmin.Create(e);
    }
    await resetFields();
    emit('submit', res);
  }
});
</script>

<style lang="less" scoped></style>
