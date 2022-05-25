<template>
  <div class="inline-block text-center" :style="getStyle">
    <div
      class="overflow-hidden cursor-pointer rounded-full border border-gray-200"
      :style="getImageWrapperStyle"
      @click="openModal"
    >
      <div v-if="sourceValue" class="rounded-full">
        <div class="cropper-image-mask" :style="getImageWrapperStyle">
          <Icon icon="ant-design:cloud-upload-outlined" color="#d6d6d6" class="mx-auto" :style="getIconStyle" />
        </div>
        <img :src="sourceValue" alt="avatar" class="w-full" />
      </div>

      <Icon v-else icon="ant-design:cloud-upload-outlined" color="#d6d6d6" class="mx-auto" :style="getIconStyle" />
    </div>
    <NButton v-if="showBtn" class="my-2" v-bind="btnProps" @click="openModal">
      {{ btnText ? btnText : '请选择图片' }}
    </NButton>

    <CropperModal
      :title="modalTitle"
      :upload-api="uploadApi"
      :src="sourceValue"
      @register="register"
      @upload-success="handleUploadSuccess"
    />
  </div>
</template>

<script lang="ts" setup>
import { computed, CSSProperties, PropType, ref, unref, watch, watchEffect } from 'vue';
import { ButtonProps, NButton } from 'naive-ui';
import { Icon } from '@iconify/vue';
import { useModal } from '@/components/basic/modal';
import CropperModal from './CropperModal.vue';

const emit = defineEmits(['update:value', 'change']);

const props = defineProps({
  width: { type: [String, Number], default: '200px' },
  value: { type: String },
  showBtn: { type: Boolean, default: true },
  btnProps: { type: Object as PropType<ButtonProps> },
  btnText: { type: String, default: '' },
  uploadApi: {
    type: Function as PropType<({ file, name }: { file: Blob; name: string }) => Promise<void>>
  },
  modalTitle: { type: String, default: '上传' }
});

const [register, { openModal }] = useModal();

const sourceValue = ref(props.value || 'https://avatars.dicebear.com/api/male/john.svg?background=%230000ff');

const getWidth = computed(() => `${`${props.width}`.replace(/px/, '')}px`);

const getIconStyle = computed(() => {
  return {
    width: `${parseInt(`${props.width}`.replace(/px/, ''), 10) / 2}px`,
    height: unref(getWidth)
  };
});

const getStyle = computed((): CSSProperties => ({ width: unref(getWidth) }));

const getImageWrapperStyle = computed((): CSSProperties => ({ width: unref(getWidth), height: unref(getWidth) }));

watchEffect(() => {
  sourceValue.value = props.value || '';
});

watch(
  () => sourceValue.value,
  (v: string) => {
    emit('update:value', v);
  }
);

function handleUploadSuccess({ source }: { source: string }) {
  sourceValue.value = source;
  emit('change', source);
}
</script>

<style scoped>
.cropper-image-mask {
  position: absolute;
  cursor: pointer;
  border-radius: inherit;
  transition: all 0.4s;
  opacity: 0;
}

.cropper-image-mask:hover {
  background: rgb(0, 0, 0, 0.4);
  opacity: 4000%;
}
</style>
