<template>
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
              :src="`https://cdn.seovx.com/ha/?mom=302&a=${item}`"
              object-fit="cover"
            />
          </div>
        </n-carousel>
      </div>
      <n-tabs :bar-width="28" type="line" class="custom-tabs">
        <n-tab-pane name="oasis" tab="全部"><PostVue /></n-tab-pane>
        <n-tab-pane name="the beatles" tab="Python">Hey Jude</n-tab-pane>
        <n-tab-pane name="jay chou" tab="Golang">七里香</n-tab-pane>
      </n-tabs>
    </template>
    <template #column>
      <rightContentVue />
    </template>
  </rightColumnVue>
</template>

<script setup>
import { useMessage } from "naive-ui";
import rightColumnVue from "@/layout/rightColumn.vue";
import PostVue from "@/components/Post.vue";
import rightContentVue from "@/components/home/rightContent.vue";
const message = useMessage();
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
</style>
