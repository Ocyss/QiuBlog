<template>
  <div>
    <n-button @click="showModal = true">清除缓存</n-button>
    <n-modal
      v-model:show="showModal"
      preset="dialog"
      title="缓存清理"
      positive-text="确认"
      negative-text="算了"
      @positive-click="submitCallback"
    >
      <n-checkbox v-model:checked="p.message">留言</n-checkbox>
      <n-checkbox v-model:checked="p.question">问答</n-checkbox>
    </n-modal>

    <n-tabs
      default-value="signin"
      size="large"
      justify-content="space-evenly"
      class="main"
      :key="key"
    >
      <n-tab-pane name="signin" tab="留言">
        <messageVue />
      </n-tab-pane>
      <n-tab-pane name="signup" tab="问答">
        <questionVue />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup lang="ts">
import api from "@/api";
import { ref, h } from "vue";
import messageVue from "@/components/admin/message/message.vue";
import questionVue from "@/components/admin/message/question.vue";
import { useMessage } from "naive-ui";
const message = useMessage();
const showModal = ref(false);
const key = ref(0);
const p = ref({ question: true, message: true });
function submitCallback() {
  api.message.clearMessage(p.value).then((res) => {
    if (res.status == 200) {
      res.data;
      message.success("清理成功!!!");
    } else {
      message.success(`清理失败。${res.message}`);
    }
  });
  key.value++;
}
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
