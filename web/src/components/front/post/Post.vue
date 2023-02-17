<template>
  <n-config-provider :theme="null" class="main">
    <div class="bgimg">
      <img :src="imgSrc" alt="" />
    </div>
    <div class="img">
      <router-link :to="`/post/${item.id}`" target="_blank">
        <img :src="imgSrc" alt="" />
      </router-link>
    </div>
    <div class="content">
      <div class="information">
        <div>
          <n-icon-wrapper :size="18" :border-radius="15">
            <n-icon :size="13" :component="Calendar" />
          </n-icon-wrapper>
          <TimerVue :t="item.created_at" />
        </div>
        <div class="rightInfo">
          <div style="margin-right: 0.5rem">
            <n-icon-wrapper
              :size="18"
              :border-radius="15"
              color="rgb(192,202,51)"
            >
              <n-icon :size="13" :component="Book" />
            </n-icon-wrapper>
            &nbsp;{{ item.uv }}
            {{ designStore.getLocale ? "阅读" : "pageview" }}
          </div>
          <div v-if="item.cname">
            <n-icon-wrapper
              color="rgb(67,160,71)"
              :size="18"
              :border-radius="15"
            >
              <n-icon :size="13" :component="PricetagsSharp" />
            </n-icon-wrapper>
            {{ item.cname.name }}
          </div>
        </div>
      </div>
      <div class="title">
        <router-link :to="`/post/${item.id}`" target="_blank">
          {{ item.title }}
        </router-link>
      </div>
      <div class="contentMain">
        <router-link :to="`/post/${item.id}`" target="_blank">
          {{ item.desc }}
        </router-link>
      </div>
      <div class="tags">
        <n-space :wrap="false">
          <n-tag
            v-for="tag in item.tags"
            @click="router.push({ name: menuTag, query: { id: tag.id } })"
            style="cursor: pointer"
            :key="tag"
            size="small"
            round
            :color="tag.color"
            :type="
              tag.color
                ? ''
                : ['primary', 'info', 'success', 'warning', 'error'][
                    Math.floor(Math.random() * 5)
                  ]
            "
          >
            {{ tag.name }}
          </n-tag>
        </n-space>
      </div>
    </div>
  </n-config-provider>
</template>

<script setup>
import { Calendar, Book, PricetagsSharp } from "@vicons/ionicons5";
import { 随机风景API } from "@/settings/config.js";
import { ref } from "vue";
import { useRouter } from "vue-router";
import TimerVue from "@/components/Timer.vue";

const designStore = inject("designStore");
const router = useRouter();
const imgSrc = ref("");
const props = defineProps(["item"]);
//渲染判断，无头图则采用随机api
if (props.item.img == "") {
  imgSrc.value = 随机风景API + `wcnm=${props.index}`;
} else {
  imgSrc.value = props.item.img;
}
</script>

<style lang="scss" scoped>
a {
  text-decoration-line: none;
  color: #fff;
}

.content {
  color: #fff;
  z-index: 10;
  padding: 24px 18px;
  width: 66.6666666%;
  text-shadow: 0 0.1875rem 0.3125rem #0000004d;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  letter-spacing: 0.06rem;
  .information {
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 0.3rem;
    > div {
      display: flex;
      align-items: center;
    }
  }
  .rightInfo > div {
    display: flex;

    align-items: center;
  }
  .title {
    font-size: 1.5rem;
    font-weight: 600;
  }
  .contentMain {
    font-size: 1rem;
    -webkit-line-clamp: 4;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    text-indent: 2rem;
  }
  .tags {
    overflow: hidden;
  }
}
.main {
  margin: 20px 0px;
  display: flex;
  width: 95%;
  max-width: 800px;
  height: 14.5rem;
  background-color: #fff;
  border-radius: 0.5rem;
  overflow: hidden;
  position: relative;
  box-sizing: border-box;
  transition: all 0.3s;
}
.main:hover {
  box-shadow: 0px 0px 6px var(--n-tab-text-color),
    0px 0px 6px var(--n-tab-text-color), 0px 0px 6px var(--n-tab-text-color),
    0px 0px 6px var(--n-tab-text-color);
  .img {
    img {
      transform: translate(3px, 3px) scale(1.05) rotate(0deg);
    }
  }
}
.right {
  flex-direction: row-reverse;
  .img {
    clip-path: polygon(6% 0, 100% 0, 100% 100%, 0 100%);
    -webkit-clip-path: polygon(6% 0, 100% 0, 100% 100%, 0 100%);
  }
}
.img {
  width: 33.333333%;
  -webkit-clip-path: polygon(0 0, 100% 0, 94% 100%, 0 100%);
  clip-path: polygon(0 0, 100% 0, 94% 100%, 0 100%);
  img {
    transition: all 0.4s;
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  z-index: 2;
}
.bgimg {
  z-index: 0;
  box-sizing: border-box;
  border: 0 solid #e5e7eb;
  position: absolute;
  width: 100%;
  height: 14.5rem;
  overflow: hidden;
  img {
    object-fit: cover;
    background-position: 50%;
    height: 100%;
    width: 100%;
    filter: blur(1.875rem) brightness(0.75);
    transform: translate(0, 0) rotate(0) skewX(0) skewY(0) scaleX(1.25)
      scaleY(1.25);
  }
}
</style>
