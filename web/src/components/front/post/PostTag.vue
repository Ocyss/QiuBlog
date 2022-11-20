<template>
  <div>
    <div class="title">{{ data.name }} :</div>

    <div class="PostSpin" v-if="data.Article != false">
      <PostVue
        v-for="(item, index) in data.Article"
        :class="index % 2 === 0 ? 'left' : 'right'"
        :key="item"
        :item="item"
      />
    </div>
    <n-empty v-else description="没有东西。。。" />
  </div>
</template>

<script setup>
import { ref } from "vue";
import PostVue from "./Post.vue";
const props = defineProps(["data"]);
</script>

<style scoped lang="scss">
.title {
  font-size: 30px;
  font-weight: 600;
}
</style>
