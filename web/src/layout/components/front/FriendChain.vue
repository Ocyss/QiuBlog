<template>
  <div class="main" v-if="!collapsed">
    <h2>友链</h2>
    <a target="_blank" href="http://www.vcabin.cn/#/home/index">如沐春风</a>
  </div>
</template>

<script setup>
import { ref } from "vue";
const props = defineProps(["collapsed"]);
</script>

<style scoped lang="scss">
.main {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
a,
a:focus {
  color: aliceblue;
  text-decoration: none;
}

a:hover {
  cursor: pointer;
  color: rgb(138, 192, 239);
}
</style>
