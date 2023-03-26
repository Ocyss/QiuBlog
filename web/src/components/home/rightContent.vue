<template>
  <div class="main">
    <aboutVue />

    <n-affix v-if="scrollableEl" :trigger-top="60" :listen-to="scrollableEl">
      <commentVue />
      <popularVue />
      <statisticsVue />
      <tagCloudVue />
    </n-affix>
  </div>
</template>

<script setup lang="ts">
import { inject, Ref } from "vue";
import aboutVue from "./card/cardAbout.vue";
import commentVue from "./card/cardComment.vue";
import popularVue from "./card/cardPopular.vue";
import statisticsVue from "./card/cardStatistics.vue";
import tagCloudVue from "./card/cardTagCloud.vue";

const props = defineProps(["affixRef"]);
const scrollableEl: Ref<HTMLElement> = inject("scrollableEl");
</script>

<style lang="scss" scoped>
.main,
.n-affix {
  display: flex;
  align-items: center;
  flex-direction: column;
  max-width: 280px;
  margin: auto;
}

.main {
  position: relative;

  .n-affix {
  }
  .n-card,
  .n-affix > .n-card {
    max-width: 280px;
    width: 20vw;
    border-radius: 8px;
    border-style: solid;
    border-width: 2px;
    margin: 10px 0px 10px 0px;
  }
}
</style>
