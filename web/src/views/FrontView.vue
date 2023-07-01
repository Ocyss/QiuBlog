<template>
  <frontVue>
    <rightColumnVue>
      <div class="content">
        <RouterView />
        <n-back-top :right="100" />
      </div>
      <template #column>
        <rightContentVue />
      </template>
    </rightColumnVue>
  </frontVue>
</template>

<script setup lang="ts">
import frontVue from "@/layout/front.vue";
import rightColumnVue from "@/layout/rightColumn.vue";
import rightContentVue from "@/components/home/rightContent.vue";
import { ref } from "vue";
const affixRef = ref(void 0);
</script>

<style lang="scss" scoped>
.content {
  width: 100%;
  max-width: 768px;
  min-height: 140vh;
}
</style>
