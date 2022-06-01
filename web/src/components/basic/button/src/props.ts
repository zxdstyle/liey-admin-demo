import { PropType } from 'vue';

export const buttonProps = {
  color: { type: String, validator: (v: any) => ['error', 'warning', 'success', ''].includes(v) },
  loading: { type: Boolean },
  disabled: { type: Boolean },
  icon: { type: String },
  iconSize: { type: Number, default: 14 },
  onClick: { type: Function as PropType<(...args: any) => any>, default: null }
};
