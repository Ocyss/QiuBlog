<template>
  <div>
    <n-tabs
      :bar-width="28"
      type="line"
      class="custom-tabs"
      animated
      @update:value="changeCategory"
    >
      <n-tab-pane
        :name="String(citem.id)"
        :tab="citem.name"
        v-for="citem in cdata.cids"
      >
        <n-spin :show="PostSpinShow" v-if="PostData[citem.id]">
          <div
            class="content"
            v-if="PostData[citem.id].length > 0"
            v-for="(item, index) in PostData[citem.id]"
            :key="item"
          >
            <PostVue
              v-if="item.created_at"
              :index="index"
              :item="item"
              :category="getCategory(item.cid)"
            />
          </div>
          <n-empty v-else description="没有东西。。。" />
          <template #description>加载中~~~~</template>
        </n-spin>
      </n-tab-pane>
    </n-tabs>
  </div>
</template>

<script setup>
import { ref } from "vue";
import PostVue from "./Post.vue";
import axios from "axios";
const PostSpinShow = ref(true);
const props = defineProps(["cdata"]);

//各分类下的文章
const PostData = ref({
  "-1": [{ id: -1 }, { id: -2 }, { id: -3 }],
});

const getCategory = (cid) => {
  return props.cdata.cids.find((item) => {
    return item.id == cid;
  });
};

//请求主页文章列表
axios
  .get("/api/v1/article/list", {
    params: { pagesize: 10, pagenum: 1, mid: props.cdata.id },
  })
  .then((res) => {
    if (res.data.status == 200) {
      PostData.value["-1"] = res.data.data.articles;
    }
  });
PostSpinShow.value = false;
function changeCategory(val) {
  PostSpinShow.value = true;
  if (PostData.value[val] == undefined) {
    axios
      .get("/api/v1/article/list", {
        params: { pagesize: 10, pagenum: 1, cid: val },
      })
      .then((res) => {
        if (res.data.status == 200) {
          PostData.value[val] = res.data.data.articles;
        }
      });
  }
  PostSpinShow.value = false;
}
</script>

<style scoped lang="scss"></style>
