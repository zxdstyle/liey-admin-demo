<template>
  <NSwitch v-bind="bindValues"></NSwitch>
</template>

<script lang="ts" setup>
import { computed, PropType, unref, useAttrs } from 'vue';
import { NSwitch } from 'naive-ui';
import { useLoading } from '@/hooks';

const emit = defineEmits(['update:value']);

interface Api {
  (params: any): Promise<any>;
}

const props = defineProps({
  api: {
    type: Function as PropType<Api>,
    required: true
  },
  field: {
    type: String,
    default: 'status'
  },
  value: {
    type: [Boolean, Number],
    default: false
  },
  round: {
    type: Boolean,
    default: false
  }
});
const attrs = useAttrs();

const { startLoading, endLoading, loading } = useLoading(false);

const onUpdate = async (value: boolean) => {
  startLoading();
  props
    .api({ [props.field]: value })
    .then(() => {
      emit('update:value', value);
    })
    .catch(() => {
      emit('update:value', !value);
    })
    .finally(() => endLoading());
};

const bindValues = computed(() => {
  return {
    ...attrs,
    round: props.round,
    value: props.value,
    loading: unref(loading),
    'onUpdate:value': onUpdate
  };
});
</script>

<style lang="less" scoped></style>
