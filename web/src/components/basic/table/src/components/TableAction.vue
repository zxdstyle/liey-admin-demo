<template>
  <NDropdown class="px-2" v-bind="bindValue">
    <NButton size="small" type="primary" quaternary>
      操作&nbsp;
      <Icon icon="ph:caret-down-bold" />
    </NButton>
  </NDropdown>
</template>

<script lang="ts" setup>
import { computed, PropType, useAttrs } from 'vue';
import { DropdownOption, NButton, NDropdown } from 'naive-ui';
import { Icon } from '@iconify/vue';

const props = defineProps({
  actions: {
    type: Array as PropType<DropdownOption[]>,
    required: true
  }
});

const emit = defineEmits(['select']);

const attrs = useAttrs();

const bindValue = computed(() => {
  return {
    options: props.actions,
    onSelect: (key: string) => emit('select', key),
    ...attrs
  };
});
</script>
