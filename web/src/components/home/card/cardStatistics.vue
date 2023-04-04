<template>
  <n-card
    class="statistics"
    :title="t('info.card.title.statistics')"
    size="small"
    v-if="data"
  >
    <div class="main">
      <div class="mainUv">
        <div>{{ $t("info.card.statistics.mainUv") }}:</div>
        <div>{{ data.main_uv }}</div>
      </div>

      <div class="wordsTotal">
        <div>{{ $t("info.card.statistics.wordsTotal") }}:</div>
        <div>
          {{ data.words_total }}
          w
        </div>
      </div>
      <div class="articleCount">
        <div>{{ $t("info.card.statistics.articleCount") }}:</div>
        <div>{{ data.article_count }}</div>
      </div>
      <div class="lastUpdated">
        <div>{{ $t("info.card.statistics.lastUpdated") }}:</div>
        <n-time unix :time="data.last_updated" type="relative" />
      </div>
      <div class="elapsedTime">
        <div>{{ $t("info.card.statistics.elapsedTime") }}:</div>
        <div>
          {{ $t("info.card.statistics.date", date) }}
        </div>
      </div>
    </div>
  </n-card>
</template>

<script setup lang="ts">
import { ref, inject, Ref, onMounted } from "vue";
import api from "@/api";
import { useDesignSettingStore } from "@/store/modules/designSetting";
import moment from "moment";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const designStore = useDesignSettingStore();

let date: Ref<any> = ref({
  cur: moment().unix(),
  d: 0,
  h: 0,
  m: 0,
  s: 0,
});

const data: Ref<any> = ref({});

onMounted(async () => {
  const res = await api.statistics.statistics();
  data.value = res.data;
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
});
</script>

<style scoped lang="scss">
.main > div {
  display: flex;
  justify-content: space-between;
}
</style>
