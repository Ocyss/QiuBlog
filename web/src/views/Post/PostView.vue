<template>
  <frontVue>
    <div style="padding: 20px">
      <div class="cover">
        <img :src="imgSrc" :alt="postData.title" loading="lazy" />
        <div class="info">
          <div class="info-head">
            <div>
              <n-icon-wrapper :size="18" :border-radius="15">
                <n-icon :size="13" :component="Calendar" />
              </n-icon-wrapper>
              <n-tooltip trigger="hover">
                <template #trigger>
                  <n-time
                    time-zone="Asia/Shanghai"
                    :time="new Date(postData.created_at)"
                    type="relative"
                  />
                </template>
                <n-time
                  time-zone="Asia/Shanghai"
                  :time="new Date(postData.created_at)"
                  format="yyyy-MM-dd HH:mm:ss"
                />
              </n-tooltip>
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
                1.5K阅读
              </div>
              <div v-if="category">
                <n-icon-wrapper
                  color="rgb(67,160,71)"
                  :size="18"
                  :border-radius="15"
                >
                  <n-icon :size="13" :component="PricetagsSharp" />
                </n-icon-wrapper>
                {{ category.name }}
              </div>
            </div>
          </div>

          <div class="info-text">
            <h1>{{ postData.title }}</h1>
            <span>{{ postData.desc }}</span>
          </div>
          <div class="info-tags">
            <n-space :wrap="false">
              <n-tag
                v-for="tag in postData.tags"
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
      </div>
      <n-divider />
      <div class="content">
        <Editor
          style="min-height: 320px"
          v-model="postData.content"
          :defaultConfig="editorConfig"
          mode="simple"
        />
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

<script setup>
import frontVue from "@/layout/front.vue";
import { ref, onMounted, reactive } from "vue";
import axios from "axios";
import { useRoute } from "vue-router";
import { useMessage } from "naive-ui";
import { 随机风景API } from "@/settings/config.js";
import { globalData } from "@/store/modules/globalData.js";
import { Calendar, Book, PricetagsSharp } from "@vicons/ionicons5";
import { computed } from "@vue/reactivity";
import "@wangeditor/editor/dist/css/style.css"; // 引入 css
import { Editor } from "@wangeditor/editor-for-vue";
const url = window.location.href;
const dataStore = globalData();
const route = useRoute();
const message = useMessage();
const postData = ref({
  id: -1,
  created_at: "2000-01-01T00:00:00+08:00",
  updated_at: "2022-01-01T00:00:00+08:00",
  deleted_at: null,
  title: "",
  img: "",
  desc: "",
  content: "",
  cid: -1,
  tags: [],
});

const editorConfig = { readOnly: true, scroll: false };

const imgSrc = computed(() => {
  return postData.value.img ? postData.value.img : 随机风景API;
});

const category = computed(() => {
  return dataStore.getCategory().find((item) => {
    return item.id == postData.value.cid;
  });
});

axios.get(`/api/v1/article/${route.params.pid}`).then((res) => {
  if (res.data.status == 200) {
    postData.value = res.data.data;
  } else {
    message.error(res.data.message);
  }
});
</script>

<style lang="scss" scoped>
.cover {
  width: 100%;
  position: relative;
  overflow: hidden;
  border-radius: 16px;
  min-height: 230px;
  background-color: #eee;
  img {
    display: block;
    position: absolute;
    left: 0;
    top: 0;
    object-fit: cover;
    width: 100%;
    height: 100%;
    filter: brightness(50%);
  }
  .info {
    position: absolute;
    color: #fff;
    width: 100%;
    height: 100%;
    box-sizing: border-box;
    text-transform: none;
    margin: 0;
    padding: 24px;
    z-index: 1;
    text-shadow: 0 1px 3px rgb(0 0 0 / 25%);
    display: flex;
    flex-direction: column;
    justify-content: space-between;
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
