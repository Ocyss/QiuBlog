<template>
  <n-spin class="postlist" :show="PostSpinShow" style="min-height: 300px">
    <div class="PostSpin" v-if="PostData">
      <div class="PostPage">
        <div class="title">{{ tname }} :</div>
        <div>{{ page["0"] }}/{{ pageCount["0"] }}</div>

        <n-checkbox
          v-if="settingStore.autuLoad"
          v-model:checked="settingStore.autuLoad"
        >
          自动加载
        </n-checkbox>
      </div>
      <PostVue
        v-for="(item, index) in PostData"
        :class="index % 2 === 0 ? 'left' : 'right'"
        :key="item"
        :item="item"
      />
      <Pagination
        :page="page"
        cid="0"
        :pageCount="pageCount"
        @upage="upPage"
        @load="load"
      />
    </div>
    <n-empty v-else description="没有东西。。。" />
  </n-spin>
</template>

<script setup lang="ts">
import { ref, inject, Ref, onMounted } from "vue";
import PostVue from "@/components/front/post/Post.vue";
import Pagination from "@/components/front/post/Pagination.vue";
import api from "@/api";
import { useRoute, useRouter } from "vue-router";
import { useProjectSettingStore } from "@/store/modules/projectSetting";
import { useHead } from "@unhead/vue";
const settingStore = useProjectSettingStore();
const route = useRoute();
const router = useRouter();
const tid = route.query.id;
const tname = ref("");
const page = ref({ "0": 1 });
const pageCount = ref({ "0": 1 });
const PostData = ref([]);
const backTopRef: Ref<any> = inject("backTopRef");
const PostSpinShow = ref(false);
const cdata = ref([]);

settingStore.getAllTags().then((res) => {
  let tag = res.find((item) => {
    return item.id == tid;
  });
  if (tag == undefined) {
    router.push({ name: "exception-404" });
  }
  tname.value = tag.name;
  useHead({ title: tname.value });
});

function getPosts() {
  PostSpinShow.value = true;
  const params = {
    pagesize: 10,
    pagenum: page.value["0"],
    tid: tid,
  };
  api.article.getList(params).then((res) => {
    PostData.value.push(
      ...res.data.map((item) => {
        item.cname = cdata.value.find((citem) => {
          return citem.id == item.cid;
        });
        return item;
      })
    );
    pageCount.value["0"] = Math.ceil(res.total / params.pagesize);
    PostSpinShow.value = false;
  });
}

function upPage(p) {
  //返回顶部
  backTopRef.value.handleClick();
  PostData.value = [];
  getPosts();
}

function load() {
  page.value["0"]++;
  getPosts();
}

onMounted(() => {
  api.category.get().then((res) => {
    cdata.value = res.data;
  });
  getPosts();
});
</script>

<style scoped lang="scss">
.title {
  font-size: 30px;
  font-weight: 600;
}
</style>
