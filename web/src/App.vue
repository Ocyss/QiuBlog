<template>
  <div id="app" v-wechat-title="$route.meta.title">
    <n-config-provider
      :locale="designStore.getLocale ? zhCN : enUS"
      :theme="designStore.darkTheme ? darkTheme : undefined"
      :theme-overrides="getThemeOverrides"
      :date-locale="designStore.getLocale ? dateZhCN : dateEnUS"
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
      <n-global-style />
    </n-config-provider>
  </div>
</template>

<script setup>
import { provide } from "vue";
import { useDesignSettingStore } from "@/store/modules/designSetting.js";
import { zhCN, dateZhCN, darkTheme, enUS, dateEnUS } from "naive-ui";
import { lighten } from "@/utils/index";
import api from "@/api";
import cookies from "vue-cookies";
import axios from "axios";

if (!cookies.get("mainuv")) {
  api.statistics.mainuv().then(() => {
    cookies.set("mainuv", "1", -1);
  });
}
const designStore = useDesignSettingStore();
provide("config", useConfig);
provide("designStore", designStore);

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
