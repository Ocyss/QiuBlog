<template>
  <div class="content">
    <n-tabs type="line" animated>
      <n-tab-pane
        name="author"
        :tab="designStore.getLocale ? '作者介绍' : 'author_introduce'"
      >
        <md-editor
          v-model="authorRef"
          preview-only
          :theme="designStore.getDarkTheme ? `dark` : `light`"
        />
      </n-tab-pane>
      <n-tab-pane
        name="project"
        :tab="designStore.getLocale ? '项目介绍' : 'project_introduce'"
      >
        <md-editor
          v-model="projectRef"
          preview-only
          :theme="designStore.getDarkTheme ? `dark` : `light`"
        />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const authorRef = ref("");
const projectRef = ref("");
const designStore = inject("designStore");

onMounted(() => {
  axios.get("static/about.md").then((res) => {
    authorRef.value = res.data;
  });
  axios
    .get(
      "https://gh.api.99988866.xyz/https://raw.githubusercontent.com/qiu-lzsnmb/QiuBlog/master/README.md"
    )
    .then((res) => {
      projectRef.value = res.data;
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
