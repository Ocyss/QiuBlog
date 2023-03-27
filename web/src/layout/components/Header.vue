<template>
  <div class="layout-header">
    <div class="menuLeft">
      <div>
        <n-icon size="27" @click="() => $emit('collapsed')" style="">
          <MenuSharp />
        </n-icon>
      </div>
      <div>
        <n-icon size="25" @click="shareShow = true">
          <ShareSocialOutline />
        </n-icon>
      </div>
    </div>
    <div class="menuRight">
      <div>
        <n-switch v-model:value="darkMode" :rail-style="railStyle">
          <template #checked>
            {{ designStore.getLocale ? "下班" : "LowerClass" }}
          </template>
          <template #unchecked>
            {{ designStore.getLocale ? "上班" : "UpperClass" }}
          </template>
          <template #icon>{{ darkMode ? "🌛" : "🌞" }}</template>
        </n-switch>
      </div>
      <div>
        <n-icon size="25" @click="locale = !locale"><Language /></n-icon>
      </div>
      <div>
        <n-popover placement="bottom" trigger="click">
          <template #trigger>
            <n-icon size="25"><ColorPalette /></n-icon>
          </template>
          <n-space class="themsBottom">
            <div
              class="themsRadio"
              :class="designStore.appTheme == thems ? 'selected' : ''"
              :style="{ backgroundColor: thems }"
              v-for="thems in appThemeList"
              @click="designStore.appTheme = thems"
            ></div>
          </n-space>
        </n-popover>
      </div>
    </div>
    <n-modal v-model:show="shareShow">
      <vue-qrcode :value="url" :options="{ width: 300 }"></vue-qrcode>
    </n-modal>
  </div>
</template>

<script setup lang="ts">
import { ref, inject, computed } from "vue";
import {
  ColorPalette,
  Language,
  ShareSocialOutline,
  MenuSharp,
} from "@vicons/ionicons5";
import { railStyle } from "@/utils";
import { useDesignSettingStore } from "@/store/modules/designSetting";
import { useI18n } from "vue-i18n";
const i18n = useI18n();
const shareShow = ref(false);
const appThemeList = ref([
  "#FC9D99",
  "#D24D57",
  "#26A65B",
  "#EB7347",
  "#AEDD81",
  "#00CCFF",
  "#D0D0D0",
  "#ff9800",
  "#FF3D68",
  "#00C1D4",
  "#71EFA3",
  "#78DEC7",
  "#1768AC",
  "#FB9300",
  "#FC5404",
  "#3B99D4",
  "#8ED14B",
  "#F06B49",
  "#ECC2F1",
  "#82C7C3",
  "#E3698A",
  "#1776EB",
  "#F5B2AC",
]);
const designStore = useDesignSettingStore();
let url = "";
if (!import.meta.env.SSR) {
  url = window.location.href;
}

//暗黑模式
const darkMode = computed({
  get: () => designStore.getDarkTheme,
  set: (val) => designStore.setDarkTheme(val),
});
//国际化语言
const locale = computed({
  get: () => designStore.getLocale,
  set: (val) => {
    i18n.locale.value = val ? "zhCN" : "enUS";
    designStore.setLocale(val);
  },
});

defineProps(["collapsed"]);
</script>

<style lang="scss" scoped>
.layout-header {
  height: 5vh;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0;
  transition: all 0.2s ease-in-out;
  width: 100%;
  z-index: 10000;
}

.menuLeft {
  display: flex;
  align-items: center;
  padding-left: 15px;
  div {
    margin: 0 3px;
  }
}

.menuRight {
  display: flex;
  align-items: center;
  padding-right: 25px;
  div {
    margin: 0 3px;
  }
}

.n-icon {
  cursor: pointer;
}
.themsBottom {
  width: 10vw;
  justify-content: center !important;
  .selected {
    border: solid 2px var(--n-text-color);
  }
}
.themsRadio {
  width: 20px;
  height: 20px;
  box-sizing: border-box;
  &:hover {
    border: solid 3px var(--n-text-color);
  }
}
</style>
