<template>
  <div class="content">
    <n-tabs type="line" animated>
      <n-tab-pane
        name="author"
        :tab="designStore.getLocale ? '作者介绍' : 'author_introduce'"
      >
        <editor v-if="contentRef.author" :content="contentRef.author" />
      </n-tab-pane>
      <n-tab-pane
        name="project"
        :tab="designStore.getLocale ? '项目介绍' : 'project_introduce'"
      >
        <editor v-if="contentRef.project" :content="contentRef.project" />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import editor from "@/components/editor.vue";

const designStore = inject("designStore");
const contentRef = ref({
  author: "",
  project: "",
});
onMounted(() => {
  axios.get("static/about.md").then((res) => {
    contentRef.value.author = res.data;
  });
  axios
    .get(
      "https://gh.api.99988866.xyz/https://raw.githubusercontent.com/qiu-lzsnmb/QiuBlog/master/README.md"
    )
    .then((res) => {
      contentRef.value.project = res.data;
    });
});
</script>

<style lang="scss" scoped></style>
