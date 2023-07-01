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
          <n-icon-wrapper :size="19.2" :border-radius="10" style="margin-right: 0.1rem">
            <n-icon :size="14.4" :component="Calendar" />
          </n-icon-wrapper>
          <TimerVue :t="item.created_at" />
        </div>
        <div class="rightInfo">
          <div style="margin-right: 0.5rem">
            <n-icon-wrapper :size="19.2" :border-radius="10" color="rgb(192,202,51)" style="margin-right: 0.1rem">
              <n-icon :size="14.4" :component="Book" />
            </n-icon-wrapper>
            &nbsp;{{ item.uv }}
            {{ designStore.getLocale ? "阅读" : "RDG" }}
          </div>
          <div v-if="item.cname">
            <n-icon-wrapper color="rgb(67,160,71)" :size="19.2" :border-radius="10" style="margin-right: 0.1rem">
              <n-icon :size="14.4" :component="PricetagsSharp" />
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
          <n-tag v-for="tag in item.tags" @click="router.push({ name: 'menuTag', query: { id: tag.id } })"
            style="cursor: pointer" :key="tag.id" size="small" round :type="['primary', 'info', 'success', 'warning', 'error'][
              Math.floor(Math.random() * 5)
            ]
              ">
            {{ tag.name }}
          </n-tag>
        </n-space>
      </div>
    </div>
  </n-config-provider>
</template>

<script setup lang="ts">
import { Calendar, Book, PricetagsSharp } from "@vicons/ionicons5";
import { inject, ref, Ref } from "vue";
import { useRouter } from "vue-router";
import TimerVue from "@/components/Timer.vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";

const designStore = useDesignSettingStore();
const router = useRouter();
const imgSrc = ref("");
const props = defineProps({
  item: Object,
  index: Number,
});
//渲染判断，无头图则采用随机api
if (props.item.img == "") {
  imgSrc.value = "/static/img/fc.jpg";
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
    * {
      white-space: nowrap;
    }

    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 12px;

    >div {
      display: flex;
      align-items: center;
    }
  }

  .rightInfo>div {
    display: flex;

    align-items: center;
  }

  .title {
    font-size: 19px;
    font-weight: 600;
  }

  .contentMain {
    font-size: 14px;
    -webkit-line-clamp: 2;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    text-indent: 1rem;

    a {
      white-space: break-spaces;
    }
  }

  .tags {
    overflow: hidden;
  }
}

.main {
  margin: 13px 0px;
  display: flex;
  width: 95%;
  max-width: 800px;
  height: 12.5rem;
  background-color: #fff;
  border-radius: 0.5rem;
  overflow: hidden;
  position: relative;
  box-sizing: border-box;
  transition: all 0.3s;
  box-shadow: 0px 0px 3px var(--n-tab-text-color),
    0px 0px 3px var(--n-tab-text-color), 0px 0px 3px var(--n-tab-text-color),
    0px 0px 3px var(--n-tab-text-color);

  &:hover {
    box-shadow: 0px 0px 12px var(--n-tab-text-color),
      0px 0px 12px var(--n-tab-text-color), 0px 0px 12px var(--n-tab-text-color),
      0px 0px 12px var(--n-tab-text-color);

    .img {
      img {
        transform: translate(3px, 3px) scale(1.05) rotate(0deg);
      }
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
  width: 33%;
  -webkit-clip-path: polygon(0 0, 100% 0, 94% 100%, 0 100%);
  clip-path: polygon(0 0, 100% 0, 94% 100%, 0 100%);

  img {
    transition: all 0.4s;
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
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
    transform: translate(0, 0) rotate(0) skewX(0) skewY(0) scaleX(1.25) scaleY(1.25);
  }
}
</style>
