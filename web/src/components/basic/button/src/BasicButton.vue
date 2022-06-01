<template>
  <NButton v-bind="getBindValue" :class="getButtonClass" @click="onClick">
    <slot></slot>
    <template #icon>
      <Icon v-if="icon" :icon="icon" :size="iconSize" />
    </template>
  </NButton>
</template>

<script lang="ts" setup>
import { computed, PropType, unref, useAttrs } from 'vue';
import { NButton } from 'naive-ui';
import { Icon } from '@iconify/vue';

const props = defineProps({
  color: { type: String, validator: (v: string) => ['error', 'warning', 'success', ''].includes(v) },
  loading: { type: Boolean },
  disabled: { type: Boolean },
  /**
   * Text before icon.
   */
  icon: { type: String },
  /**
   * preIcon and postIcon icon size.
   * @default: 14
   */
  iconSize: { type: Number, default: 14 },
  onClick: { type: Function as PropType<(...args: any) => any>, default: null }
});

const attrs = useAttrs();
const getButtonClass = computed(() => {
  const { color, disabled } = props;
  return [
    {
      [`ant-btn-${color}`]: !!color,
      [`is-disabled`]: disabled
    }
  ];
});

const getBindValue = computed(() => ({ ...unref(attrs), ...props }));
</script>

<style lang="less" scoped></style>
