<template>
  <n-tabs
    default-value="signin"
    size="large"
    justify-content="space-evenly"
    class="main"
  >
    <n-tab-pane name="signin" tab="留言">
      <messageVue />
    </n-tab-pane>
    <n-tab-pane name="signup" tab="问答">
      <questionVue />
    </n-tab-pane>
  </n-tabs>
</template>

<script setup lang="ts">
import api from "@/api";
import { ref, h } from "vue";
import messageVue from "@/components/admin/message/message.vue";
import questionVue from "@/components/admin/message/question.vue";
</script>

<style scoped lang="scss"></style>

<style lang="scss">
.switch {
  .n-switch {
    margin-bottom: 5px;
  }
}
.operation {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
