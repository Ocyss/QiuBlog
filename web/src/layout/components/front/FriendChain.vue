<template>
  <div class="main" v-if="!collapsed">
    <h2>{{ $t("component.chain.name") }}</h2>
    <!-- <a target="_blank" v-for="fc in config.friendChain" :href="fc.href">
      {{ fc.name }}
    </a> -->
    友链维护中
  </div>
</template>

<script setup lang="ts">
defineProps(["collapsed"]);
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
