<template>
  <layoutVue>
    <template #sider>
      <div class="leftMenu" style="position: absolute">
        <LogoVue :collapsed="collapsed" />
        <MenuVue />
        <FriendChainVue :collapsed="collapsed" />
      </div>
    </template>
    <template #drawer>
      <div class="leftMenu" style="position: absolute">
        <LogoVue :collapsed="false" />
        <MenuVue />
        <FriendChainVue :collapsed="false" />
      </div>
    </template>
    <template #header>
      <HeaderVue @collapsed="collapsed = !collapsed" :collapsed="collapsed" />
    </template>
    <div class="content" style="min-height: 95vh">
      <slot name="default"></slot>
    </div>
  </layoutVue>
</template>

<script setup lang="ts">
import MenuVue from "./components/front/Menu.vue";
import LogoVue from "./components/Logo.vue";
import HeaderVue from "./components/Header.vue";
import layoutVue from "./index.vue";
import FriendChainVue from "./components/front/FriendChain.vue";
import { inject, computed, ref } from "vue";
import { useProjectSettingStore } from "@/store/modules/projectSetting";

const settingStore = useProjectSettingStore();
//是否折叠
const collapsed = computed({
  get: () => settingStore.collapsed,
  set: (val) => {
    settingStore.collapsed = val;
  },
});
</script>

<style lang="scss" scoped>
.leftMenu {
  width: 100%;
  display: flex;
  flex-direction: column;
}
.content {
  display: flex;
  justify-content: center;
}
</style>
