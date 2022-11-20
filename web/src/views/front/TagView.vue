<template>
  <PostTagtVue :data="data" v-if="data" />
</template>

<script setup>
import PostTagtVue from "@/components/front/post/PostTag.vue";
import { useRoute, useRouter } from "vue-router";
import { ref } from "vue";
import api from "@/api";

const route = useRoute();
const router = useRouter();
const data = ref();

api.article.tagGetList(route.query.id).then((res) => {
  if (res.data.Article == null) {
    router.push({ name: "exception-404" });
  } else {
    data.value = res.data;
  }
});
</script>
