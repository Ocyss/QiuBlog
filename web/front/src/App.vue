<template>
  <n-config-provider
    :locale="zhCN"
    :theme="getDarkTheme"
    :theme-overrides="getThemeOverrides"
    :date-locale="dateZhCN"
  >
    <AppProvider>
      <RouterView />
    </AppProvider>
  </n-config-provider>
</template>

<script setup>
import AppProvider from "@/components/Application.vue";
import { useDesignSettingStore } from "@/store/modules/designSetting.js";
import { zhCN, dateZhCN, darkTheme } from "naive-ui";
import { lighten } from "@/utils/index";
const designStore = useDesignSettingStore();

const getThemeOverrides = computed(() => {
  const appTheme = designStore.appTheme;
  const lightenStr = lighten(designStore.appTheme, 6);
  return {
    common: {
      primaryColor: appTheme,
      primaryColorHover: lightenStr,
      primaryColorPressed: lightenStr,
    },
    LoadingBar: {
      colorLoading: appTheme,
    },
  };
});

const getDarkTheme = computed(() =>
  designStore.darkTheme ? darkTheme : undefined
);
</script>

<style lang="scss" scoped>
#main {
  display: flex;
}

.content {
  width: 100%;
}
</style>

<style lang="scss">
@import "styles/index.scss";

body {
  overflow-y: hidden;

  overflow-x: hidden;
}
* {
  margin: 0;
  padding: 0;
}
::-webkit-scrollbar {
  width: 5px;
}

::-webkit-scrollbar-thumb {
  background: rgb(255, 76, 162);
}
</style>
