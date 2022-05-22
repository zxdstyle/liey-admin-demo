<template>
  <n-radio-group>
    <n-radio v-for="item in getOptions" :key="`${item.value}`" :value="item.value" :disabled="item.disabled">
      {{ item.label }}
    </n-radio>
  </n-radio-group>
</template>
<script lang="ts">
import { defineComponent, computed } from 'vue';
import type { PropType } from 'vue';
import { isString } from 'lodash-es';
import { useRuleFormItem } from '@/hooks/core/useFormItem';

type OptionsItem = { label: string; value: string | number; disabled?: boolean };
type RadioItem = string | OptionsItem;

export default defineComponent({
  name: 'RadioButtonGroup',
  components: {},
  props: {
    value: {
      type: [String, Number] as PropType<string | number>
    },
    options: {
      type: Array as PropType<RadioItem[]>,
      default: () => []
    }
  },
  setup(props) {
    // Embedded in the form, just use the hook binding to perform form verification
    const [state] = useRuleFormItem(props);

    // Processing options value
    const getOptions = computed((): OptionsItem[] => {
      const { options } = props;
      if (!options || options?.length === 0) return [];

      const isStringArr = options.some(item => isString(item));
      if (!isStringArr) return options as OptionsItem[];

      return options.map(item => ({ label: item, value: item })) as OptionsItem[];
    });

    return { state, getOptions };
  }
});
</script>
