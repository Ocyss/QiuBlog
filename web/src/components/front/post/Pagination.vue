<template>
  <div class="main">
    <n-button
      v-if="page[cid] < pageCount[cid]"
      ref="button"
      type="primary"
      @click="emit('load')"
    >
      加载···
    </n-button>
    <n-pagination
      v-model:page="page[cid]"
      :page-count="pageCount[cid]"
      @update:page="(p) => emit('upage', p)"
    >
      <template #suffix>
        <n-checkbox v-model:checked="settingStore.autuLoad">
          自动加载
        </n-checkbox>
      </template>
    </n-pagination>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, inject } from "vue";
import { useIntersectionObserver } from "@vueuse/core";
import { useProjectSettingStore } from "@/store/modules/projectSetting";
const emit = defineEmits(["upage", "load"]);
const button = ref(void 0);
const props = defineProps(["page", "cid", "pageCount"]);
const settingStore = useProjectSettingStore();

onMounted(() => {
  setTimeout(() => {
    //进行自动加载按钮监听,按钮进入可视区域自动加载
    useIntersectionObserver(button, ([{ isIntersecting }]) => {
      if (settingStore.autuLoad && isIntersecting) {
        emit("load");
      }
    });
  }, 2000);
});
</script>

<style scoped lang="scss">
.main {
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .n-pagination {
    margin: 20px auto !important;
    align-items: center;
  }
}
</style>
