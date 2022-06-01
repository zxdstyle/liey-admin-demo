<template>
  <div class="flex flex-col gap-2">
    <div class="pb-3">
      <BasicButton quaternary icon="ion:md-arrow-round-back" @click="back">Back</BasicButton>
    </div>
    <n-card class="rounded-xl" hoverable>
      <header class="flex items-center">
        <h3 class="text-3xl">{{ role?.name }}</h3>
      </header>
    </n-card>
    <n-card title="权限" class="rounded-xl" hoverable />
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import ApiRole from '@/service/api/scaffold/role';

const route = useRoute();
const router = useRouter();
const back = () => router.back();

const role = ref<Nullable<Api.Role>>(null);

async function fetch() {
  const { data } = await ApiRole.Show(route.params.id as string);
  role.value = data;
  console.log(data);
}

onMounted(() => fetch());
</script>

<style lang="less" scoped></style>
