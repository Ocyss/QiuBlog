<template>
  <n-card
    class="comment"
    :title="designStore.getLocale ? '近期热门' : 'recently popular'"
    size="small"
  ></n-card>
</template>

<script setup lang="ts">
import { ref, inject } from "vue";
const designStore: any = inject("designStore");
</script>

<style scoped lang="scss"></style>
