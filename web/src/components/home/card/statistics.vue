<template>
  <n-card
    class="statistics"
    :title="designStore.getLocale ? '统计信息' : 'statistics'"
    size="small"
  >
    <div class="main">
      <div class="mainUv">
        <div>{{ designStore.getLocale ? "浏览量" : "Page_View" }}:</div>
        <div><n-number-animation :to="data.main_uv" /></div>
      </div>
      <div class="wordsTotal">
        <div>{{ designStore.getLocale ? "总字数" : "Words_Total" }}:</div>
        <div>
          <n-number-animation :to="data.words_total" :precision="2" />
          w
        </div>
      </div>
      <div class="articleCount">
        <div>{{ designStore.getLocale ? "文章数量" : "Article_Count" }}:</div>
        <div><n-number-animation :to="data.article_count" /></div>
      </div>
      <div class="lastUpdated">
        <div>{{ designStore.getLocale ? "最后更新于" : "Last_Updated" }}:</div>
        <n-time unix :time="data.last_updated" type="relative" />
      </div>
      <div class="elapsedTime">
        <div>{{ designStore.getLocale ? "已稳点运行" : "Run_Time" }}:</div>
        <div>
          {{
            designStore.getLocale
              ? `${date.d}天${date.h}时${date.m}分${date.s}秒`
              : `${date.d} d ${date.h} h ${date.m} m ${date.s} s`
          }}
        </div>
      </div>
    </div>
  </n-card>
</template>

<script setup>
import { ref } from "vue";
import api from "@/api";
const designStore = inject("designStore");
const data = ref({
  article_count: 0,
  words_total: 0,
  main_uv: 0,
  last_updated: 0,
});
let date = ref({
  cur: Math.round(new Date().getTime() / 1000),
  d: 0,
  h: 0,
  m: 0,
  s: 0,
});

const dates = ref("");

api.statistics.statistics().then((res) => {
  if (res.status == 200) {
    data.value = res.data;
    date.value.cur = date.value.cur - res.data.elapsed_time;
    date.value.d = parseInt(date.value.cur / 86400);
    date.value.h = parseInt((date.value.cur - date.value.d * 86400) / 3600);
    date.value.m = parseInt(
      (date.value.cur - date.value.d * 86400 - date.value.h * 3600) / 60
    );
    date.value.s =
      date.value.cur -
      date.value.d * 86400 -
      date.value.h * 3600 -
      date.value.m * 60;
  }
});
setInterval(() => {
  date.value.s++;
  if (date.value.s >= 60) {
    date.value.s = 0;
    date.value.m++;
    if (date.value.m >= 60) {
      date.value.m = 0;
      date.value.h++;
      if (date.value.h >= 60) {
        date.value.h = 0;
        date.value.d++;
      }
    }
  }
}, 1000);
</script>

<style scoped lang="scss">
.main > div {
  display: flex;
  justify-content: space-between;
}
</style>
