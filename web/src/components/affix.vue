<template>
  <n-el tag="div" class="affix">
    <div class="button"><n-icon :component="Receipt"></n-icon></div>
    <n-back-top
      ref="backTopRef"
      v-if="scrollableEl"
      :listen-to="scrollableEl"
      right="15%"
      bottom="6%"
      style="z-index: 5000"
    />
  </n-el>
</template>

<script setup lang="ts">
import { ref, inject, Ref } from "vue";
import { Receipt } from "@vicons/ionicons5";
const scrollableEl: Ref<HTMLElement> = inject("scrollableEl");
const el = ref([]);
</script>

<style scoped lang="scss">
.affix {
  position: fixed;
  right: 15%;
  bottom: 15%;
  z-index: 999;
}
.button {
  width: 44px;
  height: 44px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: var(--popover-color);
  border-radius: 50%;
  cursor: pointer;
  font-size: 20px;
  &:hover {
    color: var(--primary-color);
  }
}
</style>
