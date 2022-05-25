import { ref, toRaw, unref, reactive, onUnmounted, getCurrentInstance, computed, watchEffect, nextTick } from 'vue';
import type { ModalProps } from 'naive-ui';
import { isEqual, isFunction } from 'lodash-es';
import { tryOnUnmounted } from '@vueuse/core';
import { isProdMode } from '@/utils/common/env';
import type { UseModalReturnType, ModalMethods, ReturnMethods, UseModalInnerReturnType } from '../typing';

const dataTransfer = reactive<any>({});

const visibleData = reactive<{ [key: number]: boolean }>({});

export function useModal(): UseModalReturnType {
  const modal = ref<Nullable<ModalMethods>>(null);
  const loaded = ref<Nullable<boolean>>(false);
  const uid = ref<string>('');

  function register(modalMethod: ModalMethods, uuid: string) {
    if (!getCurrentInstance()) {
      throw new Error('useModal() can only be used inside setup() or functional components!');
    }
    uid.value = uuid;

    onUnmounted(() => {
      modal.value = null;
      loaded.value = false;
      dataTransfer[unref(uid)] = null;
    });
    if (unref(loaded) && modalMethod === unref(modal)) return;

    modal.value = modalMethod;
    loaded.value = true;
    // eslint-disable-next-line no-param-reassign
    modalMethod.emitVisible = (visible: boolean, id: number) => {
      visibleData[id] = visible;
    };
  }

  const getInstance = () => {
    const instance = unref(modal);
    if (!instance) {
      console.error('useModal instance is undefined!');
    }
    return instance;
  };

  const methods: ReturnMethods = {
    setModalProps: (props: Partial<ModalProps>): void => {
      getInstance()?.setModalProps(props);
    },

    getVisible: computed((): boolean => {
      // eslint-disable-next-line no-bitwise
      return visibleData[~~unref(uid)];
    }),

    redoModalHeight: () => {
      getInstance()?.redoModalHeight?.();
    },

    // eslint-disable-next-line default-param-last
    openModal: <T = any>(show = true, data?: T, openOnSet = true): void => {
      getInstance()?.setModalProps({
        show
      });

      if (!data) return;
      const id = unref(uid);
      if (openOnSet) {
        dataTransfer[id] = null;
        dataTransfer[id] = toRaw(data);
        return;
      }
      const equal = isEqual(toRaw(dataTransfer[id]), toRaw(data));
      if (!equal) {
        dataTransfer[id] = toRaw(data);
      }
    },

    closeModal: () => {
      getInstance()?.setModalProps({ show: false });
    }
  };

  return [register, methods];
}

export const useModalInner = (callbackFn?: Fn): UseModalInnerReturnType => {
  const modalInstanceRef = ref<Nullable<ModalMethods>>(null);
  const currentInstance = getCurrentInstance();
  const uidRef = ref<string>('');

  const getInstance = () => {
    const instance = unref(modalInstanceRef);
    if (!instance) {
      import.meta.env.PROD;
      console.error('useModalInner instance is undefined!');
    }
    return instance;
  };

  const register = (modalInstance: ModalMethods, uuid: string) => {
    isProdMode() &&
      tryOnUnmounted(() => {
        modalInstanceRef.value = null;
      });
    uidRef.value = uuid;
    modalInstanceRef.value = modalInstance;
    currentInstance?.emit('register', modalInstance, uuid);
  };

  watchEffect(() => {
    const data = dataTransfer[unref(uidRef)];
    if (!data) return;
    if (!callbackFn || !isFunction(callbackFn)) return;
    nextTick(() => {
      callbackFn(data);
    });
  });

  return [
    register,
    {
      changeLoading: (loading = true) => {
        getInstance()?.setModalProps({ loading });
      },
      getVisible: computed((): boolean => {
        return visibleData[~~unref(uidRef)];
      }),

      closeModal: () => {
        getInstance()?.setModalProps({ show: false });
      },

      setModalProps: (props: Partial<ModalProps>) => {
        getInstance()?.setModalProps(props);
      },

      redoModalHeight: () => {
        const callRedo = getInstance()?.redoModalHeight;
        callRedo && callRedo();
      }
    }
  ];
};
