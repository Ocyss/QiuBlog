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
        <n-icon size="25"><ColorPalette /></n-icon>
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
const shareShow = ref(false);

const designStore = useDesignSettingStore();
const url = window.location.href;
//暗黑模式
const darkMode = computed({
  get: () => designStore.getDarkTheme,
  set: (val) => designStore.setDarkTheme(val),
});
//国际化语言
const locale = computed({
  get: () => designStore.getLocale,
  set: (val) => designStore.setLocale(val),
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
</style>
