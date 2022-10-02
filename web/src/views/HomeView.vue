<template>
  <frontVue>
    <rightColumnVue>
      <template #content>
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
            <div class="article" v-for="(item, index) in 10">
              <n-image
                preview-disabled
                class="carousel-img"
                :src="随机美女API + `wcnm=${index}`"
                object-fit="cover"
              />
            </div>
          </n-carousel>
        </div>
        <n-tabs :bar-width="28" type="line" class="custom-tabs" animated>
          <n-tab-pane
            :name="citem.name"
            :tab="citem.name"
            v-for="citem in category"
            :key="citem"
          >
            <div
              class="content"
              v-for="(item, index) in PostData[citem.name]"
              :key="item"
            >
              <PostVue
                v-if="item.id > 0"
                :key="index"
                :index="index"
                :item="item"
                :category="category[item.cid]"
              />
              <n-skeleton v-else class="skeleton" :sharp="false" />
            </div>
          </n-tab-pane>
        </n-tabs>
      </template>
      <template #column>
        <rightContentVue />
      </template>
    </rightColumnVue>
  </frontVue>
</template>

<script setup>
import frontVue from "@/layout/front.vue";
import { ref, onMounted } from "vue";
import { useMessage } from "naive-ui";
import rightColumnVue from "@/layout/rightColumn.vue";
import PostVue from "@/components/Post.vue";
import rightContentVue from "@/components/home/rightContent.vue";
import axios from "axios";
import { globalData } from "@/store/modules/globalData.js";
import { 随机美女API } from "@/settings/config.js";
const dataStore = globalData();

const message = useMessage();
const skeleton = [
  { id: -1 },
  { id: -2 },
  { id: -3 },
  { id: -4 },
  { id: -5 },
  { id: -6 },
  { id: -7 },
  { id: -8 },
  { id: -9 },
  { id: -10 },
];
const PostData = ref({
  全部: [...skeleton],
});

const category = computed(() => dataStore.getCategory);
onMounted(() => {
  dataStore.askCategory();
  axios
    .get("/api/v1/article/list", {
      params: { pagesize: 10, pagenum: 1 },
    })
    .then((res) => {
      if (res.data.status == 200) {
        PostData.value.全部 = res.data.data.articles;
      }
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
