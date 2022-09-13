<template>
  <n-loading-bar-provider>
    <n-dialog-provider>
      <n-notification-provider>
        <n-message-provider>
          <layoutVue><slot name="default"></slot></layoutVue>
        </n-message-provider>
      </n-notification-provider>
    </n-dialog-provider>
  </n-loading-bar-provider>
</template>

<script setup>
import layoutVue from "@/layout/index.vue";
</script>

<style scoped></style>
