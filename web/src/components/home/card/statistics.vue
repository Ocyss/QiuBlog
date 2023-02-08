<template>
  <n-card class="statistics" title="统计信息" size="small">
    <div class="main">
      <div class="mainUv">
        <div>浏览量:</div>
        <div>{{ data.main_uv }}</div>
      </div>
      <div class="wordsTotal">
        <div>总字数:</div>
        <div>{{ data.words_total }}</div>
      </div>
      <div class="articleCount">
        <div>文章数量:</div>
        <div>{{ data.article_count }}</div>
      </div>
      <div class="lastUpdated">
        <div>最后更新于:</div>
        <n-time unix :time="data.last_updated" type="relative" />
      </div>
      <div class="elapsedTime">
        <div>已稳点运行:</div>
        <div>{{ dates }}</div>
      </div>
    </div>
  </n-card>
</template>

<script setup>
import { ref } from "vue";
import api from "@/api";
const data = ref({
  article_count: 0,
  words_total: 0,
  main_uv: 0,
  last_updated: 0,
});
let date = Math.round(new Date().getTime() / 1000);
const dates = ref("");
api.statistics.statistics().then((res) => {
  if (res.status == 200) {
    data.value = res.data;
    date = date - res.data.elapsed_time;
  }
});
setInterval(() => {
  date++;
  let day = parseInt(date / 86400);
  let hour = parseInt((date - day * 86400) / 3600);
  let minute = parseInt((date - day * 86400 - hour * 3600) / 60);
  let second = date - day * 86400 - hour * 3600 - minute * 60;
  dates.value = `${day}天${hour}时${minute}分${second}秒`;
}, 1000);
</script>

<style scoped lang="scss">
.main > div {
  display: flex;
  justify-content: space-between;
}
</style>
