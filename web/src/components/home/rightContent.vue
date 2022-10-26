<template>
  <div class="main">
    <aboutVue />
    <commentVue />
    <popularVue />
    <statisticsVue />
    <tagCloudVue />
  </div>
</template>

<script setup>
import aboutVue from "./card/about.vue";
import commentVue from "./card/comment.vue";
import popularVue from "./card/popular.vue";
import statisticsVue from "./card/statistics.vue";
import tagCloudVue from "./card/tagCloud.vue";
</script>

<style lang="scss" scoped>
.main {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  :deep(.n-card) {
    width: 95%;
    border-radius: 8px;
    border-style: solid;
    border-width: 2px;
    margin: 10px 0px 10px 0px;
  }
}
</style>
