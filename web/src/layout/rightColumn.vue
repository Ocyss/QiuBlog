<template>
  <n-layout has-sider sider-placement="right">
    <n-layout-content
      class="layoutContent"
      :content-style="{ padding: '0.5rem' }"
    >
      <div class="layoutContentMain">
        <slot name="default"></slot>
        <Footer style="width: 100%" />
      </div>
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

<script setup>
import { ref } from "vue";
import Footer from "./components/Footer.vue";
const settingStore = inject("projectStore");
</script>

<style lang="scss" scoped>
.layoutContent {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  :deep(.n-layout-scroll-container) {
    width: 100%;
    padding: 0px !important;
  }
}

.layoutContentMain {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
</style>
