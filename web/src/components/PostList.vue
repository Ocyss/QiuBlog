<template>
  <n-spin :show="PostSpinShow" style="min-height: 300px">
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
        <div class="PostSpin" v-if="PostData[citem.id] != false">
          <PostVue
            v-for="(item, index) in PostData[citem.id]"
            :class="index % 2 === 0 ? 'left' : 'right'"
            :key="item"
            :item="item"
          />
        </div>
        <n-empty v-else description="没有东西。。。" />
      </n-tab-pane>
    </n-tabs>
    <template #description>加载中~~~~</template>
  </n-spin>
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

//请求主页文章列表
// c是分类请求,m是菜单请求
function getPosts(id, ty = "c") {
  const params = { pagesize: 10, pagenum: 1 };
  if (ty == "m" || id == "-1") {
    params.mid = props.cdata.id;
  } else if (ty == "c") {
    params.cid = id;
  }
  axios
    .get("/api/v1/article/list", {
      params: params,
    })
    .then((res) => {
      if (res.data.status == 200) {
        PostData.value[id] = res.data.data.articles.map((item) => {
          item.cname = props.cdata.cids.find((citem) => {
            return citem.id == item.cid;
          })["name"];
          return item;
        });
      }
      PostSpinShow.value = false;
    });
}

getPosts("-1", "m");

function changeCategory(val) {
  PostSpinShow.value = true;
  if (PostData.value[val] == undefined) {
    getPosts(val, "c");
  } else {
    PostSpinShow.value = false;
  }
}
</script>

<style scoped lang="scss">
.PostSpin {
  display: flex;
  flex-direction: column;
  min-width: 200px;
  align-items: center;
}
</style>
