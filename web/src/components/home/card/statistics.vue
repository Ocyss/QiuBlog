<template>
  <n-card class="statistics" title="统计信息" size="small">
    <div v-for="item in data" class="block">
      <div>{{ item.key }}:</div>
      <div>{{ item.value }}</div>
    </div>
  </n-card>
</template>

<script setup>
import { ref } from "vue";
const data = ref([
  { key: "文章数目", value: "543" },
  { key: "已运行时间", value: "45天" },
  { key: "本站总字数", value: "123.5w" },
  { key: "当前在线人数", value: "12" },
  { key: "本站总访问量", value: "34.3w" },
  { key: "最后更新时间", value: "1天前" },
]);
</script>

<style scoped lang="scss">
.block {
  display: flex;
  justify-content: space-between;
}
</style>
