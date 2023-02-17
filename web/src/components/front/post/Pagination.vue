<template>
  <div class="main">
    <n-button ref="button" type="primary" @click="emit('load')">
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

<script setup>
import { ref, onMounted } from "vue";
import { useIntersectionObserver } from "@vueuse/core";
import { defineEmits } from "vue";

const emit = defineEmits(["upage", "load"]);
const button = ref(null);
const props = defineProps(["page", "cid", "pageCount"]);
const settingStore = inject("projectStore");

onMounted(() => {
  setTimeout(() => {
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
