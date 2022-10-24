<template>
  <layoutVue>
    <template #sider>
      <div class="leftMenu" style="position: absolute">
        <LogoVue :collapsed="collapsed" />
        <MenuVue @clickMenuItem="" />
      </div>
    </template>
    <template #drawer>
      <div class="leftMenu" style="position: absolute">
        <LogoVue :collapsed="false" />
        <MenuVue @clickMenuItem="" />
      </div>
    </template>
    <template #header>
      <HeaderVue
        style="height: 7vh"
        @collapsed="collapsed = !collapsed"
        :collapsed="collapsed"
      />
    </template>
    <div style="min-height: 93vh"><slot name="default"></slot></div>
  </layoutVue>
</template>

<script setup>
import MenuVue from "./components/front/Menu.vue";
import LogoVue from "./components/Logo.vue";
import HeaderVue from "./components/front/Header.vue";
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
