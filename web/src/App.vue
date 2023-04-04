<template>
  <div id="app">
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
              <loader v-show="settingStore.isLoad" />
              <Suspense>
                <RouterView />
              </Suspense>
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
import { useProjectSettingStore } from "@/store/modules/projectSetting";
import loader from "./components/loader.vue";
import { zhCN, enUS, dateZhCN, darkTheme, dateEnUS } from "naive-ui";
import { lighten } from "@/utils/index";
import { provide, ref, computed, inject, Ref, onMounted } from "vue";
import { useHead } from "@unhead/vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
let oldtitle: string;

const designStore = useDesignSettingStore();
const settingStore = useProjectSettingStore();

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

const titleTemplate = (title?: string) => {
  oldtitle = title;
  if (!title) {
    return t("info.title.load");
  }
  return `${title} - ${t("main.title")}`;
};

useHead({
  titleTemplate,
});

onMounted(() => {
  if (!import.meta.env.SSR) {
    settingStore.mainUV(); //统计访问量
    //调用原生接口判断是否离开了页面
    document.addEventListener("visibilitychange", function () {
      const state = document.visibilityState;
      if (state === "visible") {
        useHead({ title: oldtitle, titleTemplate });
      } else if (state === "hidden") {
        useHead({
          title: t("info.title.hidden"),
          titleTemplate: null,
        });
      }
    });
  }
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
.n-icon {
  width: 20px;
  height: 20px;
}
</style>
