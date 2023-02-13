<template>
  <n-card
    class="comment"
    :title="designStore.getLocale ? '近期热门' : 'recently popular'"
    size="small"
  ></n-card>
</template>

<script setup>
import { ref } from "vue";
import { useDesignSettingStore } from "@/store/modules/designSetting.js";
const designStore = useDesignSettingStore();
</script>

<style scoped lang="scss"></style>
