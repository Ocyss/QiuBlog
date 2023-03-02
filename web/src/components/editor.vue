<template>
  <div class="editor" ref="editorRef">
    <md-editor
      editor-id="author-id"
      v-model="contentRef"
      preview-only
      :theme="getTheme"
    />
    <!-- <n-affix
      class="affix"
      v-if="scrollableEl"
      :top="60"
      :listen-to="scrollableEl"
    >
      <n-popover :overlap="overlap" placement="left-end" trigger="click">
        <template #trigger></template>
        <md-catalog
          editor-id="author-id"
          :scroll-element="scrollableEl"
          :theme="getTheme"
        />
      </n-popover>
      <n-collapse></n-collapse>
    </n-affix> -->
  </div>
</template>

<script setup lang="ts">
import { ref, computed, inject, onMounted, Ref } from "vue";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import { useDesignSettingStore } from "@/store/modules/designSetting";

const overlap = ref(false);
const props = defineProps(["content"]);
const editorRef = ref(void 0);
const contentRef = ref(props.content);
const MdCatalog = MdEditor.MdCatalog;
const scrollableEl: Ref<HTMLElement> = inject("scrollableEl");

const designStore = useDesignSettingStore();

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
  right: 2%;
}

.catalog {
  background-color: var(--n-color);
  width: 40vw;
}
</style>
