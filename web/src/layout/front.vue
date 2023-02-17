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
    <div class="content" style="min-height: 95vh; padding-top: 5vh">
      <slot name="default"></slot>
    </div>
  </layoutVue>
</template>

<script setup>
import MenuVue from "./components/front/Menu.vue";
import LogoVue from "./components/Logo.vue";
import HeaderVue from "./components/Header.vue";
import layoutVue from "./index.vue";
import FriendChainVue from "./components/front/FriendChain.vue";
const settingStore = inject("projectStore");
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
