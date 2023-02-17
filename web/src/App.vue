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
import { useDesignSettingStore } from "@/store/modules/designSetting.js";
import { useProjectSettingStore } from "@/store/modules/projectSetting.js";
import { zhCN, dateZhCN, darkTheme, enUS, dateEnUS } from "naive-ui";
import { lighten } from "@/utils/index";
import api from "@/api";
import cookies from "vue-cookies";

if (!cookies.get("mainuv")) {
  api.statistics.mainuv().then(() => {
    cookies.set("mainuv", "1", -1);
  });
}
const designStore = useDesignSettingStore();
provide("config", useConfig);
provide("designStore", designStore);
provide("projectStore", useProjectSettingStore());

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

* {
  margin: 0;
  padding: 0;
}
::-webkit-scrollbar {
  width: 5px;
  background: rgb(248, 230, 239);
}

::-webkit-scrollbar-thumb {
  background: rgb(255, 76, 162);
}
</style>
