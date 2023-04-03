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
      <div class="article" v-for="data in datas" :key="data">
        <n-image
          preview-disabled
          class="carousel-img"
          :src="data"
          object-fit="cover"
        />
      </div>
    </n-carousel>
  </div>
  <PostListVue v-if="load" :cdata="cdata" />
</template>

<script setup lang="ts">
import PostListVue from "@/components/front/post/PostList.vue";
import { ref, computed, onMounted } from "vue";
import api from "@/api";
import { useHead } from "@unhead/vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";

const designStore = useDesignSettingStore();
const datas = ref([
  "https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1676282445541866800.webp",
  "https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1676282445542375600.webp",
  "https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1676282445542950200.webp",
  "https://qiu-blog.oss-cn-hangzhou.aliyuncs.com/Article/1676282445544555700.webp",
]);
//帖子分类数据
const cdata = ref({
  ename: "home",
  id: 0,
  link: "home",
  name: "主页",
  cids: [
    {
      id: -1,
      name: computed(() => {
        return designStore.getLocale ? "全部" : "All";
      }),
      homeshow: true,
    },
  ],
});
const load = ref(false);

useHead({
  link: [
    {
      rel: "alternate",
      type: "application/rss+xml",
      title: "RSS 2.0 Feed",
      href: "/rss/rss",
    },
    {
      rel: "alternate",
      type: "application/atom+xml",
      title: "RSS Atom Feed",
      href: "/rss/atom",
    },
    {
      rel: "alternate",
      type: "application/rss+json",
      title: "RSS JSON Feed",
      href: "/rss/json",
    },
  ],
});

onMounted(() => {
  api.category.get().then((res) => {
    res.data.map((item) => {
      if (item.homeshow) {
        cdata.value.cids.push(item);
      }
    });
    load.value = true;
  });
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
