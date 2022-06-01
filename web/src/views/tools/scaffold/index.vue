<template>
  <n-card class="rounded-xl" title="脚手架" hoverable>
    <n-form v-bind="getFormBindValue">
      <n-grid cols="24">
        <n-form-item-gi label="Table" span="8">
          <n-input />
          <n-select class="ml-4" />
        </n-form-item-gi>
      </n-grid>
    </n-form>
  </n-card>
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import { FormProps } from 'naive-ui';

const getFormBindValue = computed<FormProps>(() => {
  return {
    labelPlacement: 'left'
  };
});
</script>

<style lang="less" scoped></style>
