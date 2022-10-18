<template>
  <h1>留言板:</h1>
  <n-grid x-gap="12" y-gap="12" cols="1 s:2 m:3 l:4" responsive="screen">
    <n-gi v-for="item in 20">
      <n-card hoverable embedded>
        留言:不错，这个项目很棒，投了
        <template #action>点赞:{{ item * 3 }} 2022-10-18 21:18</template>
      </n-card>
    </n-gi>
  </n-grid>
</template>

<script setup>
import { ref } from "vue";
</script>

<style scoped lang="scss"></style>
