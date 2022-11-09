<template>
  <PostListVue :cdata="cdata" v-if="cdata" />
</template>

<script setup>
import PostListVue from "@/components/PostList.vue";
import { useRoute, useRouter } from "vue-router";
import { ref, onBeforeMount } from "vue";
import Api from "@/api";

const route = useRoute();
const router = useRouter();
const cdata = ref();

Api.menuchild.get({ link: route.params.menuName }).then((res) => {
  if (res.data.id == 0) {
    router.push({ name: "exception-404" });
  } else {
    cdata.value = res.data;
    cdata.value.cids.unshift({ id: -1, name: "全部", homeshow: true });
  }
});
</script>
