<template>
  <n-card
    class="comment"
    :title="designStore.getLocale ? '最新留言' : 'latest message'"
    size="small"
  >
    <div class="subitem" v-for="item in data" :key="item.id">
      <n-avatar
        round
        :size="35"
        :src="`http://q.qlogo.cn/headimg_dl?dst_uin=${item.qq}&spec=640&img_type=jpg`"
      />
      <div class="text">
        <div class="name">
          {{ item.name }}
        </div>
        <div
          class="content"
          :style="{ backgroundColor: randomRgb(130, 250, 0.5) }"
        >
          {{ item.content }}
        </div>
      </div>
    </div>
  </n-card>
</template>

<script setup>
import { ref } from "vue";
import { randomRgb } from "@/utils";
import api from "@/api";
const designStore = inject("designStore");
const data = ref([]);

api.message.getMessage({ pagesize: 6, pagenum: 1 }).then((res) => {
  res.data.map((item) => {
    if (item.show == true || item.show == undefined) {
      data.value.push(item);
    }
  });
});
</script>

<style scoped lang="scss">
.comment {
  :deep(.n-card__content) {
    max-height: 350px;
    overflow-y: auto;
  }
}

.subitem {
  display: flex;
  line-height: 1.3;
  .name {
    font-size: 11px;
    line-height: 25px;
  }
}
.text {
  width: 85%;
}

.content {
  min-height: 15px;
  border-radius: 5px;
  padding: 5px;
}
</style>
