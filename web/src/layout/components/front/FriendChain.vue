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
import { Ref, inject } from "vue";
import type { Config } from "@/types";

const props = defineProps(["collapsed"]);
const config: Ref<Config> = inject("config");
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
