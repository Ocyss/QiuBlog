<template>
  <n-tooltip trigger="hover" v-if="t">
    <template #trigger>
      <n-time time-zone="Asia/Shanghai" :time="new Date(t)" type="relative" />
    </template>
    <n-time
      time-zone="Asia/Shanghai"
      :time="new Date(t)"
      format="yyyy-MM-dd HH:mm:ss"
    />
  </n-tooltip>
</template>

<script setup>
import { ref } from "vue";
defineProps(["t"]);
</script>

<style scoped lang="scss"></style>
