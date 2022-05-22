<template>
  <n-select
    v-bind="getBindVal"
    v-model:value="state"
    :options="getOptions"
    :loading="loading"
    @update:value="handleChange"
  ></n-select>
</template>

<script lang="ts">
import { computed, defineComponent, ref, unref, watch } from 'vue';
import type { PropType } from 'vue';
import type { SelectOption } from 'naive-ui';
import { isFunction, omit } from 'lodash-es';
import { useLoading } from '@/hooks';
import { useRuleFormItem } from '@/hooks/core/useFormItem';

// interface State {
//   options: SelectOption[];
// }

export default defineComponent({
  name: 'ApiSelect',
  components: {},
  props: {
    value: [String, Number],
    api: {
      type: Function as PropType<(arg?: Recordable) => Promise<SelectOption[]>>,
      default: null
    },
    params: {
      type: Object as PropType<Recordable>,
      default: () => ({})
    },
    labelField: { type: String, default: 'label' },
    valueField: { type: String, default: 'value' }
  },
  setup(props, { attrs }) {
    const options = ref<SelectOption[]>([]);
    const emitData = ref<any[]>([]);
    const { startLoading, endLoading, loading } = useLoading();

    async function fetch() {
      const { api } = props;
      if (!api || !isFunction(api)) return;
      options.value = [];
      try {
        startLoading();
        const res = await api(props.params);
        if (Array.isArray(res)) {
          options.value = res;
          return;
        }
      } catch (error) {
        // eslint-disable-next-line no-console
        console.warn(error);
      } finally {
        endLoading();
      }
    }

    watch(
      () => props.params,
      () => fetch(),
      { deep: true, immediate: true }
    );

    const getOptions = computed(() => {
      const { labelField, valueField } = props;

      return unref(options).reduce((prev, next: Recordable) => {
        if (next) {
          const value = next[valueField];
          prev.push({
            ...omit(next, [labelField, valueField]),
            label: next[labelField],
            value
          });
        }
        return prev;
      }, [] as SelectOption[]);
    });

    const getBindVal = computed(() => {
      return {
        ...props,
        ...attrs
      };
    });

    const [state] = useRuleFormItem(props, 'value', 'change', emitData);

    function handleChange(_: any, val: any) {
      emitData.value = val;
    }

    return { getBindVal, getOptions, loading, handleChange, state };
  }
});
</script>
