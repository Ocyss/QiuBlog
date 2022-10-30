<template>
  <n-card class="comment" title="最新留言" size="small">
    <div class="subitem" v-for="item in 5" :key="item">
      <n-avatar round size="small">{{ item }}</n-avatar>
      <div class="text">
        <div class="name">{{ item }}(时间)</div>
        <div
          class="content"
          :style="{ backgroundColor: randomRgb(130, 250, 0.5) }"
        >
          评论内容
        </div>
      </div>
    </div>
  </n-card>
</template>

<script setup>
import { ref } from "vue";
import { randomRgb } from "@/utils";
</script>

<style scoped lang="scss">
.subitem {
  display: flex;
  line-height: 1;
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
