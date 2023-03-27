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

<script setup lang="ts">
import { ref, onMounted, inject } from "vue";
import axios from "axios";
import editor from "@/components/editor.vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";

const designStore = useDesignSettingStore();
const contentRef = ref({
  author: "",
  project: "",
});
onMounted(() => {
  axios.get("/about.md").then((res) => {
    contentRef.value.author = res.data;
  });
  axios
    .get("https://raw.githubusercontent.com/Ocyss/QiuBlog/master/README.md")
    .then((res) => {
      contentRef.value.project = res.data;
    });
});
</script>

<style lang="scss" scoped></style>
