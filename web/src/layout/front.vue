<template>
  <layoutVue>
    <template #sider>
      <div class="leftMenu" style="position: absolute">
        <LogoVue :collapsed="collapsed" />
        <MenuVue />
      </div>
    </template>
    <template #drawer>
      <div class="leftMenu" style="position: absolute">
        <LogoVue :collapsed="false" />
        <MenuVue />
      </div>
    </template>
    <template #header>
      <HeaderVue @collapsed="collapsed = !collapsed" />
    </template>
    <slot name="default"></slot>
  </layoutVue>
</template>

<script setup>
import MenuVue from "./components/front/Menu.vue";
import LogoVue from "./components/Logo.vue";
import HeaderVue from "./components/front/Header.vue";
import layoutVue from "./index.vue";

import { projectSetting } from "@/store/modules/projectSetting";
const settingStore = projectSetting();
const collapsed = computed({
  get: () => settingStore.getCollapsed,
  set: (val) => {
    settingStore.setCollapsed(val);
  },
});
</script>

<style lang="scss" scoped></style>
