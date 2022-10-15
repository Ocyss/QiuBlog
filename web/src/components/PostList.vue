<template>
  <div>
    <n-tabs :bar-width="28" type="line" class="custom-tabs" animated>
      <n-tab-pane
        :name="citem.name"
        :tab="citem.name"
        v-for="citem in cdata.cids"
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
import PostVue from "./Post.vue";
import axios from "axios";

const props = defineProps(["cdata"]);

//各分类下的文章
const PostData = ref({
  全部: [{ id: -1 }, { id: -2 }, { id: -3 }],
});

const getCategory = (cid) => {
  return props.cdata.cids.find((item) => {
    return item.id == cid;
  });
};

onMounted(() => {
  //请求主页文章列表
  const params = { pagesize: 10, pagenum: 1 };
  if (props.cdata.id > 0) {
    params.mid = props.cdata.id;
  }
  axios
    .get("/api/v1/article/list", {
      params: params,
    })
    .then((res) => {
      if (res.data.status == 200) {
        PostData.value.全部 = res.data.data.articles;
      }
    });
});
</script>

<style scoped lang="scss"></style>
