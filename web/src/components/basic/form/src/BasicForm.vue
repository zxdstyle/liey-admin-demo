<template>
  <n-form ref="formElRef" v-bind="getBindValue" :model="formModel">
    <n-row v-bind="getRow">
      <FormItem
        v-for="schema in getSchema"
        :key="schema.field"
        :schema="schema"
        :form-props="getProps"
        :form-model="formModel"
        :set-form-model="setFormModel"
        :component-props="schema.componentProps"
      />
    </n-row>
    <div class="flex justify-end items-center">
      <n-button class="mr-2" @click="resetFields">
        <template #icon>
          <Icon icon="carbon:reset" class="text-xl"></Icon>
        </template>
        重置
      </n-button>
      <n-button type="primary" attr-type="submit" :loading="loading">
        <template #icon>
          <Icon icon="ion:paper-plane" class="text-xl" />
        </template>
        提交
      </n-button>
    </div>
  </n-form>
</template>

<script lang="ts">
import { computed, defineComponent, nextTick, onMounted, reactive, ref, unref, watch } from 'vue';
import { NRow } from 'naive-ui';
import { merge } from 'lodash-es';
import dayjs from 'dayjs';
import { Icon } from '@iconify/vue';
import { useLoading } from '@/hooks';
import useFormValues from './hooks/useFormValues';
import useFormEvent from './hooks/useFormEvent';
import type { FormActionType, FormSchema, FormProps } from './types/form';
import { dateItemType } from './helper';
import { basicProps } from './props';
import { FormItem } from './components';

export default defineComponent({
  name: 'BasicForm',
  components: { FormItem, Icon, NRow },
  props: basicProps,
  emits: ['register'],
  setup(props, { attrs, emit }) {
    const formModel = reactive<Recordable>({});

    const defaultValueRef = ref<Recordable>({});
    const propsRef = ref<Partial<FormProps>>({});
    const schemaRef = ref<Nullable<FormSchema[]>>(null);
    const formElRef = ref<Nullable<FormActionType>>(null);

    const getProps = computed((): FormProps => {
      // @ts-ignore
      return { ...props, ...unref(propsRef) } as FormProps;
    });

    const { startLoading, endLoading, loading } = useLoading();
    const handleSubmit = async (e: Event) => {
      e.preventDefault();
      try {
        startLoading();
        await unref(formElRef)?.validate();
        const p = unref(propsRef);
        if (p && p.onSubmit) {
          await p.onSubmit(formModel);
        }
      } catch (e) {
        console.warn(e);
      } finally {
        endLoading();
      }
    };

    const getBindValue = computed(() => {
      const p = unref(propsRef);
      return { ...p, onSubmit: handleSubmit, ...attrs };
    });

    function setFormModel(key: string, value: any) {
      formModel[key] = value;
    }

    const getRow = computed((): Recordable => {
      const { baseRowStyle = {}, rowProps } = unref(getProps);
      return {
        style: baseRowStyle,
        gutter: 12,
        ...rowProps
      };
    });

    const getSchema = computed((): FormSchema[] => {
      const schemas: FormSchema[] = unref(schemaRef) || (unref(getProps).schemas as any);
      for (const schema of schemas) {
        const { defaultValue, component } = schema;
        // handle date type
        if (defaultValue && dateItemType.includes(component)) {
          if (!Array.isArray(defaultValue)) {
            schema.defaultValue = dayjs(defaultValue);
          } else {
            const def: any[] = [];
            defaultValue.forEach(item => {
              def.push(dayjs(item));
            });
            schema.defaultValue = def;
          }
        }
      }

      return schemas as FormSchema[];
    });

    async function setProps(formProps: Partial<FormProps>): Promise<void> {
      propsRef.value = merge(unref(propsRef) || {}, formProps);
    }

    const { initDefault } = useFormValues({ defaultValueRef, getSchema, formModel, getProps });

    const { resetFields, setFieldsValue } = useFormEvent({ emit, getProps, formModel, getSchema, defaultValueRef });

    watch(
      () => unref(getProps).model,
      () => {
        const { model } = unref(getProps);
        if (!model) return;
        nextTick(() => setFieldsValue(model));
      },
      {
        immediate: true,
        deep: true
      }
    );

    const formActionType: Partial<FormActionType> = {
      setProps,
      resetFields,
      setFieldsValue
    };

    onMounted(() => {
      nextTick(() => initDefault());
      emit('register', formActionType);
    });

    return {
      getBindValue,
      formModel,
      defaultValueRef,
      getRow,
      getProps,
      formElRef,
      getSchema,
      setFormModel,
      loading,
      formActionType: formActionType as any,
      ...formActionType
    };
  }
});
</script>
