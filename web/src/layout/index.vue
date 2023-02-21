<template>
  <n-back-top
    ref="backTopRef"
    v-if="scrollableEl"
    :listen-to="scrollableEl"
    right="20%"
  />
  <n-layout class="layout" has-sider position="absolute">
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
    <n-layout ref="layoutRef">
      <n-layout-header class="layout-header" bordered position="absolute">
        <slot name="header" :collapsed="collapsed"></slot>
      </n-layout-header>
      <n-layout-content class="layout-content">
        <slot name="default"></slot>
      </n-layout-content>
      <n-layout-footer>
        <Footer style="width: 100%" />
      </n-layout-footer>
    </n-layout>
  </n-layout>
</template>

<script setup>
import { ref, onMounted } from "vue";
import Footer from "./components/Footer.vue";
import { NBackTop } from "naive-ui";
const settingStore = inject("projectStore");
//获取内layout元素
const layoutRef = ref(void 0);
//获取可滚动元素，注入方便其他组件监听
const scrollableEl = ref(void 0);
const backTopRef = ref(void 0);
provide("scrollableEl", scrollableEl);
provide("backTopRef", backTopRef);

//是否手机模式，宽度小于700
const isMobile = computed({
  get: () => settingStore.isMobile,
  set: (val) => {
    settingStore.isMobile = val;
  },
});
//是否折叠
const collapsed = computed({
  get: () => settingStore.collapsed,
  set: (val) => {
    settingStore.collapsed = val;
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
  //将可滚动元素赋值
  scrollableEl.value = layoutRef.value?.scrollableElRef;
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
  min-height: 95vh;
  padding-top: 5vh;
}

.layout-header {
  z-index: 20;
}
</style>
