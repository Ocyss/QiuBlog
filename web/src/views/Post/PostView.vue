<template>
  <frontVue>
    <div class="main">
      <div class="cover" :style="{ backgroundImage: `url(${imgSrc})` }">
        <div class="info">
          <div class="info-head">
            <div>
              <n-icon-wrapper :size="18" :border-radius="15">
                <n-icon :size="13" :component="Calendar" />
              </n-icon-wrapper>
              <TimerVue :t="postData.created_at" />
            </div>
            <div>
              <n-icon-wrapper
                :size="18"
                :border-radius="15"
                color="rgb(192,202,51)"
              >
                <n-icon :size="13" :component="Book" />
              </n-icon-wrapper>
              {{ uv }}阅读
            </div>
            <div v-if="category">
              <n-icon-wrapper
                color="rgb(67,160,71)"
                :size="18"
                :border-radius="15"
              >
                <n-icon :size="13" :component="PricetagsSharp" />
              </n-icon-wrapper>
              {{ category }}
            </div>
          </div>

          <div class="info-text">
            <h1 class="title">
              {{ postData.title }}
            </h1>
            <span>{{ postData.desc }}</span>
          </div>
          <div class="info-tags">
            <n-space :wrap="false">
              <n-tag
                v-for="tag in postData.tags"
                :key="tag.id"
                size="small"
                round
                :type="
                  ['primary', 'info', 'success', 'warning', 'error'][
                    Math.floor(Math.random() * 5)
                  ]
                "
              >
                {{ tag.name }}
              </n-tag>
            </n-space>
          </div>
        </div>
      </div>
      <n-divider />
      <div class="content">
        <editorVue v-if="postData.content" :content="postData.content" />
      </div>
      <n-divider />
      <div class="copyright">
        <strong>本文链接：</strong>
        <a :href="url" :title="url" target="_blank" rel="noopener">
          {{ url }}
        </a>
        <br />
        <strong>版权声明：</strong>
        本文采用
        <a
          href="https://creativecommons.org/licenses/by-nc-sa/3.0/cn/deed.zh"
          target="_blank"
          rel="external nofollow noopener noreferrer"
        >
          CC BY-NC-SA 3.0 CN
        </a>
        协议进行许可
      </div>
      <n-divider />
    </div>
  </frontVue>
</template>

<script setup lang="ts">
import frontVue from "@/layout/front.vue";
import { ref, inject, computed, Ref, onServerPrefetch, onMounted } from "vue";
import api from "@/api";
import { useRoute } from "vue-router";
import { useMessage } from "naive-ui";
import { Calendar, Book, PricetagsSharp } from "@vicons/ionicons5";
import TimerVue from "@/components/Timer.vue";
import editorVue from "@/components/editor.vue";
import { useHead } from "@unhead/vue";

const route = useRoute();
const message = useMessage();
const category = ref("");
const url = ref("");
if (!import.meta.env.SSR) {
  url.value = window.location.href;
}

const postData = ref({
  id: 0,
  created_at: "2022-09-07T12:39:08+08:00",
  updated_at: "2022-09-07T12:39:08+08:00",
  title: "",
  img: "",
  desc: "",
  content: "",
  cid: 1,
  tags: [
    {
      id: 0,
      name: "",
    },
  ],
});
const uv = ref(void 0);

const imgSrc = computed(() => {
  return postData.value.img ? postData.value.img : "/static/img/fc.jpg";
});

async function getArticle() {
  const article_res = await api.article.get(
    route.params.pid as unknown as number
  );
  postData.value = article_res.data;
  uv.value = article_res.uv;
  category.value = article_res.category;
  useHead({
    title: postData.value.title,
    meta: [{ name: "description", content: postData.value.desc }],
  });
}

onServerPrefetch(async () => {
  await getArticle();
  postData.value.content = "";
});

onMounted(async () => {
  await getArticle();
});
</script>

<style lang="scss" scoped>
.main {
  padding: 20px;
  max-width: 1000px;
  width: 80%;
}
.cover {
  width: 100%;
  position: relative;
  overflow: hidden;
  border-radius: 16px;
  min-height: 230px;
  background-color: #eee;
  text-shadow: 0 0.1875rem 0.3125rem #1c1f21;
  background-position: 50%;
  background-size: cover;
  &::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 1;
    background-color: rgba(0, 0, 0, 0.4);
  }
  .info {
    min-height: 230px;
    position: relative;
    color: #fff;
    width: 100%;
    box-sizing: border-box;
    text-transform: none;
    margin: 0;
    padding: 24px;
    z-index: 1;
    display: flex;
    align-items: center;
    justify-content: space-evenly;
    flex-direction: column;
    .info-text {
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      .title {
        font-weight: 700;
        font-size: 1.6rem;
        text-align: center;
        letter-spacing: 0.25rem;
      }
    }
    .info-head {
      display: flex;
      align-items: center;
      justify-content: space-between;
      font-size: 0.3rem;
      > div {
        display: flex;
        align-items: center;
        margin: 5px;
      }
    }
  }
}

.content {
  margin-top: 20px;
  padding: 10px;
}

.copyright {
  a {
    text-decoration: none;
    color: rgb(255, 47, 47);
  }
}
</style>
