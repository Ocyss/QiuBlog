<template>
  <div class="layout-header">
    <div class="menuLeft">
      <n-icon size="27" @click="() => $emit('collapsed')" style="">
        <MenuSharp />
      </n-icon>
      <n-icon size="25"><ShareSocialOutline /></n-icon>
    </div>
    <div class="menuRight">
      <n-switch>
        <template #checked>傍晚六点下班</template>
        <template #unchecked>午夜零点下班</template>
      </n-switch>
      <n-icon size="25"><Language /></n-icon>
      <n-icon size="25"><ColorPalette /></n-icon>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import {
  ColorPalette,
  Language,
  ShareSocialOutline,
  MenuSharp,
  Airplane,
} from "@vicons/ionicons5";

import { projectSetting } from "@/store/modules/projectSetting";
const settingStore = projectSetting();
</script>

<style lang="scss" scoped>
.layout-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0;
  height: 7vh;
  transition: all 0.2s ease-in-out;
  width: 100%;
  z-index: 500;
}

.menuLeft {
  display: flex;
  align-items: center;
  padding-left: 15px;
}

.menuRight {
  display: flex;
  align-items: center;
  padding-right: 25px;
}

.n-icon {
  cursor: pointer;
}
</style>
