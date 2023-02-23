<template>
  <div id="app" v-wechat-title="$route.meta.title" v-if="useConfig">
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

<script setup lang="ts">
import { useDesignSettingStore } from "@/store/modules/designSetting";
import { zhCN, dateZhCN, darkTheme, enUS, dateEnUS } from "naive-ui";
import { lighten } from "@/utils/index";
import api from "@/api";
import { provide, ref, computed, inject, Ref } from "vue";
import axios from "axios";
import type { VueCookies } from "vue-cookies";
import type { Config } from "@/types";

const useConfig: Ref<Config> = ref(void 0);
const $cookies = inject<VueCookies>("$cookies");
const designStore = useDesignSettingStore();

provide("config", useConfig);

axios.get("static/config.json5").then((res) => {
  useConfig.value = new Function("return " + res.data)();
});
if (!$cookies.get("mainuv")) {
  api.statistics.mainuv().then(() => {
    $cookies.set("mainuv", "1", -1);
  });
}

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
img {
  user-select: none;
  -webkit-user-drag: none;
}
::-webkit-scrollbar {
  width: 3px;
  background: rgb(248, 230, 239);
}
@media screen and (max-width: 700px) {
  ::-webkit-scrollbar {
    width: 0px;
  }
}
::-webkit-scrollbar-thumb {
  background: rgb(255, 76, 162);
}

.md-editor-dark,
.md-editor-light {
  transition: all 0.2s ease-in-out;
  --md-bk-color: var(--n-color) !important;
  --md-color: var(--n-title-text-color) !important;
  --md-theme-color: var(--n-title-text-color) !important;
}
.md-editor {
  .default-theme * {
    white-space: break-spaces;
  }
  .default-theme img {
    border-width: 0px !important;
  }
}
</style>
