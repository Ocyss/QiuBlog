<template>
  <h1>问答:</h1>
  <n-grid x-gap="12" y-gap="12" cols="1 m:2 l:3" responsive="screen">
    <n-gi v-for="item in 20" :key="item">
      <n-card :title="`这个提问是用来测试的嘛?? ${item}`" hoverable embedded>
        是的，这是一个测试提问 {{ item * 5 }}
        <template #action>2022-10-18 21:18 点赞:{{ item * 3 }}</template>
      </n-card>
    </n-gi>
  </n-grid>
</template>

<script setup>
import { ref } from "vue";
</script>

<style scoped lang="scss"></style>
