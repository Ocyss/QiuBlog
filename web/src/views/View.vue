<template>
  <frontVue>
    <rightColumnVue>
      <RouterView :key="$route.fullPath" />
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
</script>

<style lang="scss" scoped></style>
