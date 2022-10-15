<template>
  <div class="popular">
    <n-carousel
      autoplay
      draggable
      show-arrow
      :space-between="20"
      effect="custom"
      :transition-props="{ name: 'creative' }"
      trigger="hover"
      :interval="2000"
    >
      <div class="article" v-for="(item, index) in 3">
        <n-image
          preview-disabled
          class="carousel-img"
          :src="随机美女API + `wcnm=${index}`"
          object-fit="cover"
        />
      </div>
    </n-carousel>
  </div>
  <PostListVue :cdata="cdata" v-if="cdata.id == 0" />
</template>

<script setup>
import PostListVue from "@/components/PostList.vue";
import { ref, onBeforeMount, onMounted } from "vue";
import { 随机美女API } from "@/settings/config.js";
import { globalData } from "@/store/modules/globalData.js";
const dataStore = globalData();

//获取全部分类列表
const category = computed(() => dataStore.getCategory);

//筛选主页显示分类
// const TabCategory = computed(() =>
//   category.value.filter((item) => {
//     return item.homeshow;
//   })
// );

const cdata = ref({
  ename: "home",
  id: -1,
  link: "home",
  name: "主页",
  cids: [{ id: -1, name: "全部", homeshow: true }],
});

onBeforeMount(() => {
  //初始化全部分类
  dataStore.askCategory();
  cdata.value.cids = category.value.filter((item) => {
    return item.homeshow;
  });
  cdata.value.id = 0;
});
</script>

<style lang="scss" scoped>
$popularHeight: 320px;
.popular {
  width: 100%;
  height: $popularHeight;
  vertical-align: middle;
  .article {
    display: flex;
    height: $popularHeight;
  }
}
.carousel-img {
  width: 100%;
}
.carousel-img :deep(img) {
  width: 100%;
}

:deep(.creative-enter-from),
:deep(.creative-leave-to) {
  opacity: 0;
  transform: scale(0.8);
}

:deep(.creative-enter-active),
:deep(.creative-leave-active) {
  transition: all 0.3s ease;
}
.content {
  display: flex;
  justify-content: center;
  align-items: center;
}
.skeleton {
  z-index: 10;
  margin-top: 20px;
  padding: 24px 18px;
  width: 88%;
  height: 200px;
}
</style>
