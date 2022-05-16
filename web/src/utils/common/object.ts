/** 设置对象数据 */
import { ComputedRef, Ref, unref } from 'vue';

export type DynamicProps<T> = {
  [P in keyof T]: Ref<T[P]> | T[P] | ComputedRef<T[P]>;
};

export function objectAssign<T extends Record<string, any>>(target: T, source: Partial<T>) {
  Object.assign(target, source);
}

export function getDynamicProps<T, U>(props: T): Partial<U> {
  const ret: Recordable = {};

  Object.keys(props).forEach(key => {
    ret[key] = unref((props as Recordable)[key]);
  });

  return ret as Partial<U>;
}
