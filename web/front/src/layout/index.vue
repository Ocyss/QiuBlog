<template>
  <n-layout has-sider>
    <n-layout-sider
      show-trigger="bar"
      @collapse="collapsed = true"
      @expand="collapsed = false"
      :collapsed="collapsed"
      collapse-mode="width"
      :collapsed-width="64"
      :width="leftMenuWidth"
      :native-scrollbar="false"
      class="layout-sider"
    >
      <LogoVue :collapsed="collapsed" />
      <MenuVue />
    </n-layout-sider>
    <n-layout>
      <n-layout-header><HeaderVue /></n-layout-header>
      <n-layout-content>
        <n-space>
          <n-button>Default</n-button>
          <n-button type="tertiary">Tertiary</n-button>
          <n-button type="primary">Primary</n-button>
          <n-button type="info">Info</n-button>
          <n-button type="success">Success</n-button>
          <n-button type="warning">Warning</n-button>
          <n-button type="error">Error</n-button>
        </n-space>
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<script setup>
import MenuVue from "./components/Menu.vue";
import LogoVue from "./components/Logo.vue";
import HeaderVue from "./components/Header.vue";
import { ref } from "vue";
const collapsed = ref(false);

const leftMenuWidth = computed(() => {
  return collapsed.value ? 64 : 200;
});
</script>

<style lang="scss" scoped>
.n-scrollbar {
  display: flex;
  justify-content: center;
  align-items: center;
}

.layout-sider {
  min-height: 100vh;
  box-shadow: 2px 0 8px 0 rgb(29 35 41 / 5%);
  background-color: antiquewhite;
  position: relative;
  z-index: 13;
  transition: all 0.2s ease-in-out;
}

.layout-sider-fix {
  position: fixed;
  top: 0;
  left: 0;
}
</style>
