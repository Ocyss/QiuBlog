<template>
  <adminVue>
    <div class="article" style="border: 1px solid #ccc">
      <Toolbar
        style="border-bottom: 1px solid #ccc"
        :editor="editorRef"
        :defaultConfig="toolbarConfig"
        :mode="mode"
      />
      <Editor
        style="height: 500px; overflow-y: hidden"
        v-model="valueHtml"
        :defaultConfig="editorConfig"
        :mode="mode"
        @onCreated="handleCreated"
      />
      <n-button type="success">发布</n-button>
    </div>
  </adminVue>
</template>

<script setup>
import "@wangeditor/editor/dist/css/style.css";
import { onBeforeUnmount, ref, shallowRef } from "vue";
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";

import adminVue from "@/layout/admin.vue";
// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef();
const mode = ref("default");
// 内容 HTML
const valueHtml = ref("");

const toolbarConfig = {};
const editorConfig = {
  placeholder: "请输入内容...",
  MENU_CONF: {
    uploadImage: {
      fieldName: "file",
      // 上传图片的配置
      server: "/api/v1/upload/image",
      maxNumberOfFiles: 15,
      maxFileSize: 10 * 1024 * 1024, // 10M
      meta: {
        class: "Article",
      },
      withCredentials: true,
      timeout: 5 * 1000, // 5 秒
      base64LimitSize: 5 * 1024, // 5kb
    },
  },
};

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value;
  if (editor == null) return;
  editor.destroy();
});

const handleCreated = (editor) => {
  editorRef.value = editor; // 记录 editor 实例，重要！
};
</script>

<style lang="scss" scoped>
.article {
  width: 80%;
  margin: auto;
}
</style>
