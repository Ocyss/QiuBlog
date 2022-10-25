<template>
  <div id="app" v-wechat-title="$route.meta.title">
    <n-config-provider
      :locale="getLocale"
      :theme="getDarkTheme"
      :theme-overrides="getThemeOverrides"
      :date-locale="getDateLocale"
    >
      <n-loading-bar-provider>
        <n-dialog-provider>
          <n-notification-provider>
            <n-message-provider>
              <RouterView />
            </n-message-provider>
          </n-notification-provider>
        </n-dialog-provider>
      </n-loading-bar-provider>
    </n-config-provider>
  </div>
</template>

<script setup>
import { useDesignSettingStore } from "@/store/modules/designSetting.js";
import { zhCN, dateZhCN, darkTheme, enUS, dateEnUS } from "naive-ui";
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

const getLocale = computed(() => (designStore.locale ? zhCN : enUS));
const getDateLocale = computed(() =>
  designStore.locale ? dateZhCN : dateEnUS
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
