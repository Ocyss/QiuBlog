<template>
  <frontVue>
    <RouterView :key="$route.fullPath" />
  </frontVue>
</template>

<script setup>
import frontVue from "@/layout/front.vue";
</script>

<style lang="scss" scoped></style>
