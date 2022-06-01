<template>
  <n-form ref="permissionForm" label-placement="left" :rules="rules" :model="modelValue" @submit="onSubmit">
    <n-grid cols="24">
      <n-form-item-gi span="12" label="父级权限" path="parent_id">
        <ApiTreeSelect v-model:value="modelValue.parent_id" v-bind="getTreeSelect" />
      </n-form-item-gi>
      <n-form-item-gi span="18" label="名称" path="name">
        <n-input v-model:value="modelValue.name" />
      </n-form-item-gi>
      <n-form-item-gi span="18" label="标识" path="slug">
        <n-input v-model:value="modelValue.slug" />
      </n-form-item-gi>
      <n-form-item-gi span="20" label="排序值" path="sort_num">
        <n-slider
          v-model:value="modelValue.sort_num"
          :max="1000"
          :marks="{ 100: '靠后', 500: '居中', 700: '靠前', 950: '置顶' }"
        />
      </n-form-item-gi>

      <n-grid-item span="24">
        <n-divider title-placement="left">权限规则</n-divider>
        <n-grid v-for="(rule, index) in modelValue.rules" :key="index" cols="24" x-gap="5">
          <n-form-item-gi
            label="HTTP方法"
            span="10"
            :path="`rules[${index}].http_methods`"
            :rule="[{ required: true, trigger: 'blur', type: 'array', message: '请选择HTTP方法' }]"
          >
            <n-select v-model:value="rule.http_methods" :options="methods" multiple clearable />
          </n-form-item-gi>
          <n-form-item-gi
            label="URL"
            span="10"
            :path="`rules[${index}].http_path`"
            :rule="[{ required: true, trigger: 'blur', message: '请输入URL' }]"
          >
            <n-input v-model:value="rule.http_path" clearable />
          </n-form-item-gi>
          <n-grid-item span="4">
            <BasicButton icon="ic:sharp-plus" class="mr-2" size="small" type="primary" circle @click="addRule" />
            <BasicButton icon="ic:sharp-minus" size="small" circle @click="delRule(index)" />
          </n-grid-item>
        </n-grid>
      </n-grid-item>
      <n-grid-item span="22" class="w-full flex justify-end gap-5">
        <BasicButton icon="ion:paper-plane" type="primary" attr-type="submit" :loading="loading">提交</BasicButton>
      </n-grid-item>
    </n-grid>
  </n-form>
</template>

<script lang="ts" setup>
import { computed, PropType, ref, unref } from 'vue';
import { useMessage } from 'naive-ui';
import { useLoading } from '@/hooks';
import ApiPermission from '@/service/api/scaffold/permission';
import { FormActionType } from '@/components/basic/form/src/types/form';
import { ApiTreeSelect } from '@/components/basic/form';
import ApiMenu from '@/service/api/scaffold/menu';

const message = useMessage();
const permissionForm = ref<Nullable<FormActionType>>(null);
const props = defineProps({
  model: {
    type: Object as PropType<Api.Permission>,
    default: () => {}
  }
});

const emit = defineEmits(['submit', 'update:model']);
const { startLoading, endLoading, loading } = useLoading();

const methods = [
  { value: 'ANY', label: '所有方法' },
  { value: 'GET', label: 'GET' },
  { value: 'POST', label: 'POST' },
  { value: 'PUT', label: 'PUT' },
  { value: 'DELETE', label: 'DELETE' },
  { value: 'PATCH', label: 'PATCH' },
  { value: 'HEAD', label: 'HEAD' },
  { value: 'OPTIONS', label: 'OPTIONS' }
];

const addRule = () => {
  const val = props.model;
  if (!val.rules) {
    val.rules = [];
  }
  val.rules.push({ http_methods: [], http_path: '' });
  emit('update:model', val);
};

const delRule = (idx: number) => {
  const val = props.model;
  if (val.rules && val.rules.length === 1) {
    message.warning('至少保留一条权限规则');
    return;
  }
  val.rules && val.rules.splice(idx, 1);
  emit('update:model', val);
};

const rules = {
  name: { required: true, trigger: 'blur', message: '请输入权限名称' },
  slug: { required: true, trigger: 'blur', message: '请输入权限唯一标识' },
  rules: { required: true, message: '请输入权限规则' }
};

const modelValue = computed<Api.Permission>({
  get: () => {
    const val = props.model;
    if (!val.rules) {
      val.rules = [{ http_methods: [], http_path: '' }];
    }
    return val;
  },
  set: val => emit('update:model', val)
});

async function resetFields() {}

async function onSubmit(e: Event) {
  e.preventDefault();
  try {
    startLoading();
    await unref(permissionForm)?.validate();
    let res;
    if (props.model.id && props.model.id > 0) {
      res = await ApiPermission.Update(props.model.id, props.model);
    } else {
      res = await ApiPermission.Create(props.model);
    }
    await resetFields();
    emit('submit', res);
  } catch (e) {
    console.warn(e);
  } finally {
    endLoading();
  }
}

const getTreeSelect = {
  api: ApiPermission.TreeData,
  keyField: 'id',
  labelField: 'name',
  afterFetch: (data: Api.Permission[]): Api.Permission[] => {
    return [
      {
        id: 0,
        name: '顶级权限',
        children: data
      }
    ];
  }
};
</script>

<style lang="less" scoped></style>
