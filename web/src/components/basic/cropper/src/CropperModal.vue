<template>
  <BasicModal
    v-bind="$attrs"
    style="width: 800px; position: fixed; left: 0; right: 0; top: 100px"
    :can-fullscreen="false"
    preset="dialog"
    negative-text="取消"
    positive-text="确认上传"
    @register="register"
    @positive-click="handleOk"
  >
    <div class="flex">
      <div class="w-1/2 h-340px">
        <div class="cropper">
          <CropperImage
            v-if="src"
            :src="src"
            height="300px"
            :circled="circled"
            @cropend="handleCropend"
            @ready="handleReady"
          />
        </div>
        <div class="flex justify-between items-center mt-10px">
          <div>
            <NUpload :file-list="[]" accept="image/*" @before-upload="handleBeforeUpload">
              <NTooltip placement="bottom">
                <template #trigger>
                  <BasicButton size="small" pre-icon="ant-design:upload-outlined" type="primary" />
                </template>
                选择图片
              </NTooltip>
            </NUpload>
          </div>

          <NSpace>
            <NTooltip placement="bottom">
              <template #trigger>
                <BasicButton
                  type="primary"
                  pre-icon="ant-design:reload-outlined"
                  size="tiny"
                  :disabled="!src"
                  @click="handlerToolbar('reset')"
                />
              </template>
              重置
            </NTooltip>

            <NTooltip placement="bottom">
              <template #trigger>
                <BasicButton
                  type="primary"
                  pre-icon="ant-design:rotate-left-outlined"
                  size="tiny"
                  :disabled="!src"
                  @click="handlerToolbar('rotate', -45)"
                />
              </template>
              逆时针旋转
            </NTooltip>

            <NTooltip placement="bottom">
              <template #trigger>
                <BasicButton
                  type="primary"
                  pre-icon="ant-design:rotate-right-outlined"
                  size="tiny"
                  :disabled="!src"
                  @click="handlerToolbar('rotate', 45)"
              /></template>

              顺时针旋转
            </NTooltip>
            <NTooltip placement="bottom">
              <template #trigger>
                <BasicButton
                  type="primary"
                  pre-icon="vaadin:arrows-long-h"
                  size="tiny"
                  :disabled="!src"
                  @click="handlerToolbar('scaleX')"
                />
              </template>
              水平翻转
            </NTooltip>
            <NTooltip placement="bottom">
              <template #trigger>
                <BasicButton
                  type="primary"
                  pre-icon="vaadin:arrows-long-v"
                  size="tiny"
                  :disabled="!src"
                  @click="handlerToolbar('scaleY')"
                />
              </template>
              垂直翻转
            </NTooltip>
            <NTooltip placement="bottom">
              <template #trigger>
                <BasicButton
                  type="primary"
                  pre-icon="ant-design:zoom-in-outlined"
                  size="tiny"
                  :disabled="!src"
                  @click="handlerToolbar('zoom', 0.1)"
                />
              </template>
              放大
            </NTooltip>
            <NTooltip placement="bottom">
              <template #trigger>
                <BasicButton
                  type="primary"
                  pre-icon="ant-design:zoom-out-outlined"
                  size="tiny"
                  :disabled="!src"
                  @click="handlerToolbar('zoom', -0.1)"
                />
              </template>
              缩小
            </NTooltip>
          </NSpace>
        </div>
      </div>
      <div class="w-1/2 h-340px">
        <div class="cropper-preview">
          <img v-if="previewSource" :src="previewSource" alt="预览" />
        </div>
        <template v-if="previewSource">
          <div class="flex justify-between items-center pl-3 mt-5">
            <img
              v-for="size in previewSizes"
              :key="size"
              :src="previewSource"
              alt="preview"
              class="border rounded-full border-gray-300"
              :style="{ width: size, height: size }"
            />
          </div>
        </template>
      </div>
    </div>
  </BasicModal>
</template>

<script lang="ts" setup>
import { PropType, ref, unref } from 'vue';
import { NSpace, NTooltip, NUpload, UploadFileInfo } from 'naive-ui';
import { isFunction } from 'lodash-es';
import { useModalInner } from '@/components/basic/modal/src/hooks/useModal';
import { dataURLtoBlob } from '@/utils/file/base64Conver';
import type { CropendResult, Cropper } from './typing';
import CropperImage from './Cropper.vue';

type apiFunParams = { file: Blob; name: string; filename: string };

const props = defineProps({
  circled: { type: Boolean, default: true },
  uploadApi: {
    type: Function as PropType<(params: apiFunParams) => Promise<any>>,
    required: true
  }
});

const emit = defineEmits(['uploadSuccess', 'register']);

let filename = '';
const src = ref('');
const previewSource = ref('');
const cropper = ref<Cropper>();
let scaleX = 1;
let scaleY = 1;

const [register, { closeModal, setModalProps }] = useModalInner();

// Block upload
function handleBeforeUpload({ file }: { file: UploadFileInfo; fileList: UploadFileInfo[] }) {
  if (!file.file) return;
  const reader = new FileReader();
  reader.readAsDataURL(file.file);
  src.value = '';
  previewSource.value = '';
  reader.onload = function (e) {
    src.value = (e.target?.result as string) ?? '';
    filename = file.name;
  };
  return false;
}

function handleCropend({ imgBase64 }: CropendResult) {
  previewSource.value = imgBase64;
}

function handleReady(cropperInstance: Cropper) {
  cropper.value = cropperInstance;
}

function handlerToolbar(event: string, arg?: number) {
  if (event === 'scaleX') {
    scaleX = arg = scaleX === -1 ? 1 : -1;
  }
  if (event === 'scaleY') {
    scaleY = arg = scaleY === -1 ? 1 : -1;
  }
  // @ts-ignore
  cropper?.value?.[event]?.(arg);
}

async function handleOk() {
  if (!unref(previewSource)) return;
  const { uploadApi } = props;
  if (!uploadApi || !isFunction(uploadApi)) {
    return Promise.reject(new Error('missing upload api'));
  }

  const blob = dataURLtoBlob(previewSource.value);
  const result = await uploadApi({ name: 'file', file: blob, filename });
  emit('uploadSuccess', { source: previewSource.value, data: result.data });
  closeModal();
  console.log(222);
}

const previewSizes = ['32px', '48px', '64px', '80px'];
</script>

<style scoped>
.cropper {
  height: 300px;
  background-image: linear-gradient(45deg, rgb(0 0 0 / 25%) 25%, transparent 0, transparent 75%, rgb(0 0 0 / 25%) 0),
    linear-gradient(45deg, rgb(0 0 0 / 25%) 25%, transparent 0, transparent 75%, rgb(0 0 0 / 25%) 0);
  background-position: 0 0, 12px 12px;
  background-size: 24px 24px;
}

.cropper-preview {
  width: 220px;
  height: 220px;
  margin: 0 auto;
  overflow: hidden;
  border: 1px solid rgb(0, 0, 0, 0.15);
  border-radius: 50%;
}

.cropper-preview img {
  width: 100%;
  height: 100%;
}
</style>
