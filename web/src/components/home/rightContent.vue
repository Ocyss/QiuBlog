<template>
  <div ref="mainRef" class="main">
    <aboutVue />
    <n-affix :trigger-top="200" position="absolute" :listen-to="() => mainRef">
      <commentVue />
      <popularVue />
      <statisticsVue />
      <tagCloudVue />
    </n-affix>
  </div>
</template>

<script setup>
import aboutVue from "./card/about.vue";
import commentVue from "./card/comment.vue";
import popularVue from "./card/popular.vue";
import statisticsVue from "./card/statistics.vue";
import tagCloudVue from "./card/tagCloud.vue";
import { ref } from "vue";
const props = defineProps(["affixRef"]);
const mainRef = ref(null);
</script>

<style lang="scss" scoped>
.main,
.n-affix {
  display: flex;
  align-items: center;
  flex-direction: column;
  max-width: 280px;
}

.main {
  position: relative;
  .n-affix {
    width: 100%;
  }
  .n-card,
  .n-affix > .n-card {
    width: 95%;
    border-radius: 8px;
    border-style: solid;
    border-width: 2px;
    margin: 10px 0px 10px 0px;
  }
}
</style>
