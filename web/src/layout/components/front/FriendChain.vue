<template>
  <div class="main" v-if="!collapsed">
    <h2>友链</h2>
    <a
      v-if="config"
      target="_blank"
      v-for="fc in config.friendChain"
      :href="fc.href"
    >
      {{ fc.name }}
    </a>
  </div>
</template>

<script setup lang="ts">
import { ref, inject } from "vue";
const props = defineProps(["collapsed"]);
const config: any = inject("config");
</script>

<style scoped lang="scss">
.main {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
a,
a:focus {
  color: var(--n-title-text-color);
  text-decoration: none;
  &:hover {
    text-decoration: underline;
  }
}
</style>
