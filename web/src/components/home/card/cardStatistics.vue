<template>
  <n-card
    class="statistics"
    :title="designStore.getLocale ? '统计信息' : 'statistics'"
    size="small"
  >
    <div class="main">
      <div class="mainUv">
        <div>{{ designStore.getLocale ? "浏览量" : "Page_View" }}:</div>
        <div>{{ data.main_uv }}</div>
      </div>

      <div class="wordsTotal">
        <div>{{ designStore.getLocale ? "总字数" : "Words_Total" }}:</div>
        <div>
          {{ data.words_total }}
          w
        </div>
      </div>
      <div class="articleCount">
        <div>{{ designStore.getLocale ? "文章数量" : "Article_Count" }}:</div>
        <div>{{ data.article_count }}</div>
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

<script setup lang="ts">
import { ref, inject, Ref } from "vue";
import api from "@/api";
import { useDesignSettingStore } from "@/store/modules/designSetting";
import moment from "moment";
const designStore = useDesignSettingStore();

let date: Ref<any> = ref({
  cur: moment().unix(),
  d: 0,
  h: 0,
  m: 0,
  s: 0,
});

const res = await api.statistics.statistics();

const data: Ref<any> = ref(res.data);

date.value.cur = date.value.cur - res.data.elapsed_time;
date.value.d = Math.floor(date.value.cur / 86400);
date.value.h = Math.floor((date.value.cur - date.value.d * 86400) / 3600);
date.value.m = Math.floor(
  (date.value.cur - date.value.d * 86400 - date.value.h * 3600) / 60
);

date.value.s =
  date.value.cur -
  date.value.d * 86400 -
  date.value.h * 3600 -
  date.value.m * 60;

if (!import.meta.env.SSR) {
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
}
</script>

<style scoped lang="scss">
.main > div {
  display: flex;
  justify-content: space-between;
}
</style>
