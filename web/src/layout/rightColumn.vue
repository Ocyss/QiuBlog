<template>
  <n-layout has-sider sider-placement="right">
    <n-layout-content
      class="layoutContent"
      :content-style="{ padding: '0.5rem' }"
    >
      <div class="layoutContentMain"><slot name="content"></slot></div>
    </n-layout-content>

    <n-layout-sider
      bordered
      class="layoutSide"
      :content-style="{ padding: '0.5rem' }"
      width=""
      v-if="!settingStore.isMobile"
    >
      <slot name="column"></slot>
    </n-layout-sider>
  </n-layout>
</template>

<script setup>
import { ref } from "vue";

import { projectSetting } from "@/store/modules/projectSetting";
const settingStore = projectSetting();
</script>

<style lang="scss" scoped>
.layoutContent {
  display: flex;
  justify-content: center;
  align-items: center;
  :deep(.n-layout-scroll-container) {
    width: 100%;
    max-width: 768px;
  }
}
.layoutSide :deep(.n-layout-sider-scroll-container) {
  display: flex;
  align-items: center;
  flex-direction: column;
}
</style>
