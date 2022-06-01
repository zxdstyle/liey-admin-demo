<template>
  <BasicForm @register="registerForm" />
</template>

<script lang="ts" setup>
import { transformRouteNameToRoutePath } from '@/utils';
import ApiMenu from '@/service/api/scaffold/menu';
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
    {
      field: 'parent_id',
      label: '父级菜单',
      component: 'ApiTreeSelect',
      span: 10,
      componentProps: {
        api: ApiMenu.TreeData,
        keyField: 'id',
        labelField: 'title',
        afterFetch: formatTreeData
      }
    },
    { field: 'title', label: '菜单名称', component: 'Input', span: 14 },
    {
      field: 'name',
      label: '唯一标识',
      component: 'Input',
      helpMessage: '由客户端定义',
      span: 8,
      onChange({ model }) {
        model.path = transformRouteNameToRoutePath(model.name);
      }
    },
    { field: 'path', label: '路由Path', component: 'Input', span: 16, dynamicDisabled: true },
    {
      field: 'keepalive',
      label: 'Keepalive',
      component: 'Switch',
      componentProps: { round: false },
      span: 6,
      defaultValue: true
    },
    { field: 'hidden', label: '是否隐藏', component: 'Switch', componentProps: { round: false }, span: 6 },
    {
      field: 'require_auth',
      label: '是否需要权限',
      component: 'Switch',
      helpMessage: '新建的菜单需要绑定权限才可以显示，如果选择是则无需绑定权限',
      componentProps: { round: false },
      span: 6
    },
    {
      field: 'sort_num',
      label: '排序值',
      component: 'Slider',
      span: 16,
      componentProps: { max: 1000, marks: { 100: '靠后', 500: '居中', 700: '靠前', 950: '置顶' } },
      helpMessage: '值越大排序越靠前'
    },
    {
      field: 'icon',
      label: '图标',
      component: 'IconPicker',
      span: 16,
      helpMessage: 'http://icones.js.org 可以找到所有支持的图标'
    }
  ],
  rules: {
    parent_id: { required: true, type: 'number', trigger: 'blur', message: '请选择父级菜单' },
    title: { required: true, trigger: 'blur', message: '请输入菜单名称' },
    name: { required: true, trigger: 'blur', message: '请输入菜单标识' },
    path: { required: true, trigger: 'blur', message: '请输入菜单路径' }
  },
  onSubmit: async e => {
    let res;
    if (e.id && e.id > 0) {
      res = await ApiMenu.Update(e.id, e);
    } else {
      res = await ApiMenu.Create(e);
    }
    await resetFields();
    emit('submit', res);
  }
});
</script>

<style lang="less" scoped></style>
