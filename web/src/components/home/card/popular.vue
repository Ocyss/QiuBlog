<template>
  <n-card
    class="comment"
    :title="designStore.getLocale ? '近期热门' : 'recently popular'"
    size="small"
  ></n-card>
</template>

<script setup>
import { ref } from "vue";
const designStore = inject("designStore");
</script>

<style scoped lang="scss"></style>
