<template>
  <n-card
    class="comment"
    :title="designStore.getLocale ? '近期热门' : 'recently popular'"
    size="small"
  ></n-card>
</template>

<script setup lang="ts">
import { ref, inject } from "vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";

const designStore = useDesignSettingStore();
</script>

<style scoped lang="scss"></style>
