<template>
  <PostListVue :cdata="cdata" v-if="cdata.id != -1" />
</template>

<script setup>
import PostListVue from "@/components/PostList.vue";
import { useRoute } from "vue-router";
import { ref, onBeforeMount } from "vue";
import axios from "axios";

const route = useRoute();
const cdata = ref({ id: -1 });

onBeforeMount(() => {
  axios
    .get("/api/v1/menuchild", { params: { link: route.params.menuName } })
    .then((res) => {
      if (res.data.status == 200) {
        if (res.data.data.id == 0) {
          console.log(404);
        } else {
          cdata.value = res.data.data;
          cdata.value.cids.unshift({ id: -1, name: "全部", homeshow: true });
        }
      }
    });
});
</script>
