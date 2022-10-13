<template>
  <div>
    <n-tabs :bar-width="28" type="line" class="custom-tabs" animated>
      <n-tab-pane
        :name="citem.name"
        :tab="citem.name"
        v-for="citem in TabCategory"
        :key="citem"
      >
        <div
          class="content"
          v-for="(item, index) in PostData[citem.name]"
          :key="item"
        >
          <PostVue
            v-if="item.id > 0"
            :key="index"
            :index="index"
            :item="item"
            :category="getCategory(item.cid)"
          />
          <n-skeleton v-else class="skeleton" :sharp="false" />
        </div>
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useMessage } from "naive-ui";
import PostVue from "./Post.vue";
import axios from "axios";
import { globalData } from "@/store/modules/globalData.js";
import { useRouter } from "vue-router";
const router = useRouter();
const dataStore = globalData();
const message = useMessage();
//文章列表
const skeleton = [
  { id: -1 },
  { id: -2 },
  { id: -3 },
  { id: -4 },
  { id: -5 },
  { id: -6 },
  { id: -7 },
  { id: -8 },
  { id: -9 },
  { id: -10 },
];
//各分类下的文章
const PostData = ref({
  全部: [...skeleton],
});
//获取全部分类列表
const category = computed(() => dataStore.getCategory);
//筛选主页显示分类
const TabCategory = computed(() =>
  category.value.filter((item) => {
    return item.homeshow;
  })
);
const getCategory = (cid) => {
  return category.value.find((item) => {
    return item.id == cid;
  });
};

onMounted(() => {
  //初始化分类列表
  dataStore.askCategory(false);
  //请求主页文章列表
  axios
    .get("/api/v1/article/list", {
      params: { pagesize: 10, pagenum: 1 },
    })
    .then((res) => {
      if (res.data.status == 200) {
        PostData.value.全部 = res.data.data.articles;
      }
    });
});
</script>

<style scoped lang="scss"></style>
