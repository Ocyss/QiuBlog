<template>
  <PostListVue :cdata="cdata" v-if="cdata" />
</template>

<script setup lang="ts">
import PostListVue from "@/components/front/post/PostList.vue";
import { useRoute, useRouter } from "vue-router";
import { ref, onMounted, onServerPrefetch } from "vue";
import api from "@/api";

const route = useRoute();
const router = useRouter();
const cdata = ref(void 0);

async function getMenus() {
  const res = await api.menuchild.get({ link: route.params.menuName });
  if (res.data.id == 0) {
    router.push({ name: "exception-404" });
  } else {
    cdata.value = res.data;
    cdata.value.cids.unshift({ id: -1, name: "全部", homeshow: true });
  }
}

onServerPrefetch(() => {
  getMenus();
});
onMounted(() => {
  if (!cdata.value) {
    getMenus();
  }
});
</script>
