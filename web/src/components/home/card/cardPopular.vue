<template>
  <n-card
    class="comment"
    :title="t('info.card.title.popular')"
    size="small"
  ></n-card>
</template>

<script setup lang="ts">
import { ref, inject } from "vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const designStore = useDesignSettingStore();
</script>

<style scoped lang="scss"></style>
