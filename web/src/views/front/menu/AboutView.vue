<template>
  <div class="content">
    <div v-highlight v-html="htmlMD" id="content"></div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import { marked } from "marked";
const htmlMD = ref("");
onMounted(() => {
  axios
    .get(
      "https://gh.api.99988866.xyz/https://raw.githubusercontent.com/qiu-lzsnmb/QiuBlog/master/README.md"
    )
    .then((res) => {
      htmlMD.value = marked.parse(res.data);
    });
});
</script>

<style lang="scss">
.content {
  a {
    color: rgb(88, 166, 255);
  }
}
</style>
