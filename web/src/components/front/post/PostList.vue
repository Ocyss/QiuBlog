<template>
  <n-spin class="postlist" :show="PostSpinShow" style="min-height: 300px">
    <!-- <n-tabs
      :bar-width="28"
      type="line"
      class="custom-tabs"
      animated
      @update:value="changeCategory"
    >
      <n-tab-pane
        v-for="citem in cdata.cids"
        :key="citem"
        :name="String(citem.id)"
        :tab="citem.name"
      >
        <div class="PostSpin" v-if="PostData[citem.id] != false">
          <div class="PostPage">
            <div>{{ page[citem.id] }}/{{ pageCount[citem.id] }}</div>

            <n-checkbox
              v-if="settingStore.autuLoad"
              v-model:checked="settingStore.autuLoad"
            >
              自动加载
            </n-checkbox>
          </div>

          <PostVue
            v-for="(item, index) in PostData[citem.id]"
            :class="index % 2 === 0 ? 'left' : 'right'"
            :key="item"
            :item="item"
          />

          <Pagination
            :page="page"
            :cid="citem.id"
            :pageCount="pageCount"
            @upage="upPage"
            @load="load"
          />
        </div>
        <n-empty v-else description="没有东西。。。" />
      </n-tab-pane>
    </n-tabs> -->
    <div style="height: 2000px">{{ PostSpinShow }}</div>
    <template #description>加载中~~~~</template>
  </n-spin>
</template>

<script setup lang="ts">
import { inject, ref, Ref, onMounted } from "vue";
import PostVue from "./Post.vue";
import api from "@/api";
import Pagination from "./Pagination.vue";
import { useProjectSettingStore } from "@/store/modules/projectSetting";
const PostSpinShow = ref(true);
const props = defineProps(["cdata"]);

const page = ref({ "-1": 1 });
const pageCount = ref({ "-1": 1 });
const cid = ref(-1);

const settingStore = useProjectSettingStore();
const backTop: Ref<any> = inject("BackTop");
//各分类下的文章
const PostData = ref({});

//请求主页文章列表
// c是分类请求,m是菜单请求
function getPosts() {
  PostSpinShow.value = true;
  const params: {
    pagesize: number;
    pagenum: number;
    cid?: number;
    mid?: number;
  } = {
    pagesize: 10,
    pagenum: page.value[cid.value],
  };
  if (cid.value == -1) {
    params.mid = props.cdata.id;
  } else {
    params.cid = cid.value;
  }
  api.article.getList(params).then((res) => {
    if (!PostData.value[cid.value]) {
      PostData.value[cid.value] = [];
    }
    PostData.value[cid.value].push(
      ...res.data.map((item) => {
        item.cname = props.cdata.cids.find((citem) => {
          return citem.id == item.cid;
        });
        return item;
      })
    );
    pageCount.value[cid.value] = Math.ceil(res.total / params.pagesize);
    PostSpinShow.value = false;
  });
}

function upPage(p) {
  //返回顶部
  backTop.value?.backTop();
  PostData.value[cid.value] = [];
  getPosts();
}

function load() {
  page.value[cid.value]++;
  getPosts();
}

function changeCategory(val) {
  PostSpinShow.value = true;
  cid.value = val;
  if (PostData.value[val] == undefined) {
    getPosts();
  } else {
    PostSpinShow.value = false;
  }
}

onMounted(() => {
  getPosts();
});
</script>

<style scoped lang="scss">
.PostSpin {
  display: flex;
  flex-direction: column;
  min-width: 200px;
  align-items: center;
  .PostPage {
    margin-right: auto;
    display: flex;
    justify-content: left;
    > div {
      margin-right: 8px;
    }
  }
}

.postlist {
  :deep(.n-pagination) {
    width: 80%;
    margin: 0 auto;
    justify-content: center;
  }
}
</style>
