<template>
  {{ data }}
  <NTransfer v-bind="getBindValue" />
</template>

<script lang="ts" setup>
import { computed, onMounted, PropType, ref, useAttrs } from 'vue';
import { NTransfer, TransferProps } from 'naive-ui';
import { isFunction } from 'lodash-es';

const props = defineProps({
  api: {
    type: Function as PropType<(arg?: Recordable) => Promise<Recordable>>,
    required: true
  },
  params: { type: Object, default: () => {} },
  afterFetch: { type: Function, default: (val: Recordable) => val }
});

const attrs = useAttrs();
const data = ref([]);

async function fetch() {
  const { api } = props;
  if (!api || !isFunction(api)) return;

  try {
    const res = await api(props.params);
    if (props.afterFetch) {
      data.value = props.afterFetch(res.data);
    } else {
      data.value = res.data;
    }
  } catch (e) {
    console.error(e);
  }
}

onMounted(() => fetch());

const getBindValue = computed<TransferProps>(() => {
  return {
    ...attrs,
    options: data.value
  };
});
</script>

<style lang="less" scoped></style>
