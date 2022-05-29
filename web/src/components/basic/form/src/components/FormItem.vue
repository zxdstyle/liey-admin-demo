<script lang="tsx">
import { computed, defineComponent, unref } from 'vue';
import type { PropType } from 'vue';
import { isBoolean, isFunction, upperFirst } from 'lodash-es';
import { componentMap } from '@/components/basic/form/src/componentMap';
import type { FormProps, FormSchema } from '../types/form';
import HelpMessage from './HelpMessage.vue';

export default defineComponent({
  name: 'FormItem',
  props: {
    schema: {
      type: Object as PropType<FormSchema>,
      default: () => ({})
    },
    formProps: {
      type: Object as PropType<FormProps>,
      default: () => ({})
    },
    allDefaultValues: {
      type: Object as PropType<Recordable>,
      default: () => ({})
    },
    formModel: {
      type: Object as PropType<Recordable>,
      default: () => ({})
    },
    setFormModel: {
      type: Function as PropType<(key: string, value: any) => void>,
      default: null
    },
    componentProps: {
      type: Object as PropType<Recordable>,
      default: () => {}
    }
  },
  setup(props) {
    const getValues = computed(() => {
      const { allDefaultValues, formModel, schema } = props;
      const { mergeDynamicData } = props.formProps;
      return {
        field: schema.field,
        model: formModel,
        values: {
          ...mergeDynamicData,
          ...allDefaultValues,
          ...formModel
        } as Recordable,
        schema
      };
    });

    const getDisable = computed(() => {
      const { disabled: globDisabled } = props.formProps;
      const { dynamicDisabled } = props.schema;
      const { disabled: itemDisabled = false } = unref(props.componentProps || {});
      let disabled = !!globDisabled || itemDisabled;
      if (isBoolean(dynamicDisabled)) {
        disabled = dynamicDisabled;
      }
      if (isFunction(dynamicDisabled)) {
        disabled = dynamicDisabled(unref(getValues));
      }
      return disabled;
    });

    function getShow(): { isShow: boolean; isIfShow: boolean } {
      const { show, ifShow } = props.schema;

      let isShow = true;
      let isIfShow = true;

      if (isBoolean(show)) {
        isShow = show;
      }
      if (isBoolean(ifShow)) {
        isIfShow = ifShow;
      }
      if (isFunction(show)) {
        isShow = show(unref(getValues));
      }
      if (isFunction(ifShow)) {
        isIfShow = ifShow(unref(getValues));
      }
      return { isShow, isIfShow };
    }

    function renderComponent(Comp: ReturnType<typeof defineComponent>) {
      const { renderComponentContent, field, changeEvent = 'update:value', valueField } = props.schema;
      const eventKey = `on${upperFirst(changeEvent)}`;

      const on = {
        [eventKey]: (...args: Nullable<Recordable>[]) => {
          const [e] = args;
          const target = e ? e.target : null;
          const value = target ? target.value : e;
          props.setFormModel(field, value);

          if (props.schema && props.schema.onChange) {
            props.schema.onChange(unref(getValues));
          }
        }
      };

      const bindValue: Recordable = {
        [valueField || 'value']: props.formModel[field]
      };

      const compAttr: Recordable = {
        ...props.componentProps,
        ...on,
        ...bindValue,
        disabled: unref(getDisable)
      };

      const compSlot = isFunction(renderComponentContent)
        ? { ...renderComponentContent(unref(getValues)) }
        : {
            default: () => renderComponentContent
          };
      return <Comp {...compAttr}>{compSlot}</Comp>;
    }

    function renderHelpMessage() {
      const { helpMessage } = props.schema;
      const getHelpMessage = isFunction(helpMessage) ? helpMessage(unref(getValues)) : helpMessage;

      return <HelpMessage message={getHelpMessage} />;
    }

    function renderFormItem(Comp: ReturnType<typeof defineComponent>) {
      const formItemAttr: Recordable = { path: props.schema.field, ...props.schema };
      const formItemSlot = {
        default: () => renderComponent(Comp),
        label: () => {
          const { helpMessage, label } = props.schema;
          if (!helpMessage) {
            return label;
          }
          return (
            <div>
              <span>{label}</span>
              {renderHelpMessage()}
            </div>
          );
        }
      };

      return <n-form-item {...formItemAttr}>{formItemSlot}</n-form-item>;
    }

    function renderContent() {
      const { component } = props.schema;
      const Comp = componentMap.get(component) as ReturnType<typeof defineComponent>;
      if (component === 'Divider') {
        return renderComponent(Comp);
      }
      return renderFormItem(Comp);
    }

    return () => {
      const { isIfShow, isShow } = getShow();

      const getColAttr: Recordable = {
        span: props.schema?.span || 24
      };

      return (
        isIfShow && (
          <n-col v-show={isShow} {...getColAttr}>
            {renderContent()}
          </n-col>
        )
      );
    };
  }
});
</script>
