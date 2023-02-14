<template>
  <n-layout class="layout" has-sider>
    <n-layout-sider
      v-if="!isMobile"
      show-trigger="arrow-circle"
      collapse-mode="width"
      class="layout-sider"
      :collapsed="collapsed"
      :collapsed-width="64"
      :width="leftMenuWidth"
      :native-scrollbar="false"
      @collapse="collapsed = true"
      @expand="collapsed = false"
      bordered
    >
      <slot name="sider" :collapsed="collapsed"></slot>
    </n-layout-sider>

    <n-drawer
      v-model:show="showSideDrawder"
      width="35%"
      placement="left"
      class="layout-side-drawer"
    >
      <slot name="drawer" :collapsed="collapsed"></slot>
    </n-drawer>
    <n-layout>
      <n-layout-header class="layout-header" bordered>
        <slot name="header" :collapsed="collapsed"></slot>
      </n-layout-header>
      <n-layout-content class="layout-content">
        <slot name="default"></slot>
      </n-layout-content>
    </n-layout>
  </n-layout>
  <Footer />
</template>

<script setup>
import { ref, onMounted } from "vue";
import { projectSetting } from "@/store/modules/projectSetting";
import Footer from "./components/Footer.vue";
const settingStore = projectSetting();

//是否手机模式，宽度小于700
const isMobile = computed({
  get: () => settingStore.getIsMobile,
  set: (val) => settingStore.setIsMobile(val),
});
//是否折叠
const collapsed = computed({
  get: () => settingStore.getCollapsed,
  set: (val) => {
    settingStore.setCollapsed(val);
  },
});
//是否显示侧抽屉
const showSideDrawder = computed({
  get: () => isMobile.value && collapsed.value,
  set: (val) => (collapsed.value = val),
});
//左侧菜单宽度
const leftMenuWidth = ref("140");

//宽度改变自适应
const watchWidth = () => {
  const Width = document.body.clientWidth;
  if (Width <= 700) {
    isMobile.value = true;
    leftMenuWidth.value = "25%";
  } else {
    isMobile.value = false;
    leftMenuWidth.value = "150";
  }
};

onMounted(() => {
  //监听大小改动，自适应页面
  watchWidth();
  window.addEventListener("resize", watchWidth);
  window["$loading"] = useLoadingBar();
  window["$loading"].finish();
});
</script>

<style lang="scss" scoped>
.layout {
  display: flex;
  flex-direction: row;
  flex: auto;
}
.n-scrollbar {
  display: flex;
  justify-content: center;
  align-items: center;
}

.layout-sider {
  height: 100vh;
  box-shadow: 2px 0 8px 0 rgb(29 35 41 / 5%);
  position: relative;
  z-index: 16;
  transition: all 0.2s ease-in-out;
}

.layout-content {
  flex: auto;
  height: 93vh;
}
</style>
