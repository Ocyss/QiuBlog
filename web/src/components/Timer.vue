<template>
  <n-tooltip trigger="hover" v-if="load">
    <template #trigger>
      <n-time time-zone="Asia/Shanghai" :time="new Date(t)" type="relative" />
    </template>
    <n-time
      time-zone="Asia/Shanghai"
      :time="new Date(t)"
      format="yyyy-MM-dd HH:mm:ss"
    />
  </n-tooltip>
  <n-time
    time-zone="Asia/Shanghai"
    :time="new Date(t)"
    format="yyyy-MM-dd HH:mm:ss"
  />
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
const props = defineProps(["t"]);

const load = computed(() => {
  return !import.meta.env.SSR && props.t;
});
</script>

<style scoped lang="scss"></style>
