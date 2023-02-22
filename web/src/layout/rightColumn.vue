<template>
  <n-layout has-sider sider-placement="right">
    <n-layout-content
      class="layoutContent"
      :content-style="{ padding: '1rem' }"
    >
      <slot name="default"></slot>
    </n-layout-content>

    <n-layout-sider
      class="layoutSide"
      :content-style="{ padding: '0.5rem' }"
      width=""
      v-if="!settingStore.isMobile"
    >
      <slot name="column"></slot>
    </n-layout-sider>
  </n-layout>
</template>

<script setup lang="ts">
import { ref, onMounted, inject } from "vue";

const settingStore: any = inject("projectStore");
</script>

<style lang="scss" scoped>
.layoutContent {
  width: 100%;
  display: flex;
  align-items: center;
  :deep(.n-layout-scroll-container) {
    width: 100%;
    display: flex;
    align-items: center;
    flex-direction: column;
  }
}

.layoutSide {
  max-width: 320px;
  width: 23%;
}
</style>
