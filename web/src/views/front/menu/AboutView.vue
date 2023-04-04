<template>
  <div class="content">
    <n-tabs type="line" animated>
      <n-tab-pane name="author" :tab="t('component.about.author')">
        <editor v-if="contentRef.author" :content="contentRef.author" />
      </n-tab-pane>
      <n-tab-pane name="project" :tab="t('component.about.project')">
        <editor v-if="contentRef.project" :content="contentRef.project" />
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup lang="ts">
import { ref, onServerPrefetch, inject } from "vue";
import request from "@/utils/request";
import editor from "@/components/editor.vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";
import { useI18n } from "vue-i18n";

const { t } = useI18n();
const designStore = useDesignSettingStore();
const contentRef = ref({
  author: "",
  project: "",
});
onServerPrefetch(async () => {
  request.get("/about.md").then((res) => {
    contentRef.value.author = res.data;
  });
  request
    .get("https://raw.githubusercontent.com/Ocyss/QiuBlog/master/README.md")
    .then((res) => {
      contentRef.value.project = res.data;
    });
});
</script>

<style lang="scss" scoped></style>
