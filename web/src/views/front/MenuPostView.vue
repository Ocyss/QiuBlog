<template>
  <PostListVue :cdata="cdata" v-if="cdata" />
</template>

<script setup>
import PostListVue from "@/components/PostList.vue";
import { useRoute } from "vue-router";
import { ref, onBeforeMount } from "vue";
import Api from "@/api";

const route = useRoute();
const cdata = ref();

Api.menuchild.get({ link: route.params.menuName }).then((res) => {
  if (res.data.id == 0) {
    console.log(404);
  } else {
    cdata.value = res.data;
    cdata.value.cids.unshift({ id: -1, name: "全部", homeshow: true });
  }
});
</script>
