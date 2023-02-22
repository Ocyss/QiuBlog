<template>
  <div class="editor" ref="editorRef">
    <md-editor
      editor-id="author-id"
      v-model="contentRef"
      preview-only
      :theme="getTheme"
    />
    <n-affix
      class="affix"
      v-if="scrollableEl"
      :top="60"
      :listen-to="scrollableEl"
    >
      <md-catalog
        editor-id="author-id"
        :scroll-element="scrollableEl"
        :theme="getTheme"
      />
    </n-affix>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, inject } from "vue";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
const props = defineProps(["content"]);
const editorRef = ref(void 0);
const contentRef = ref(props.content);
const MdCatalog = MdEditor.MdCatalog;
const scrollableEl: any = inject("scrollableEl");
const designStore: any = inject("designStore");

const getTheme = computed(() => {
  return designStore.getDarkTheme ? `dark` : `light`;
});
</script>

<style scoped lang="scss">
.editor {
  display: flex;
}

.affix {
  position: fixed;
  right: 30%;
}
</style>
