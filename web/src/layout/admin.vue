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
import MenuVue from "./components/admin/Menu.vue";
import HeaderVue from "./components/admin/Header.vue";
import LogoVue from "./components/Logo.vue";
import layoutVue from "./index.vue";

import { projectSetting } from "@/store/modules/projectSetting";
const settingStore = projectSetting();
//是否折叠
const collapsed = computed({
  get: () => settingStore.getCollapsed,
  set: (val) => {
    settingStore.setCollapsed(val);
  },
});
</script>

<style lang="scss" scoped>
.leftMenu {
  width: 100%;
  display: flex;
  flex-direction: column;
}
</style>
