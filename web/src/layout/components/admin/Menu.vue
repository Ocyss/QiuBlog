<template>
  <NMenu
    :options="menus"
    :collapsed="collapsed"
    :value="menusKey"
    :collapsed-width="64"
    :collapsed-icon-size="20"
    :indent="24"
    @update:value="clickMenuItem"
  />
</template>

<script setup lang="ts">
import { ref, h, onBeforeMount } from "vue";
import { NIcon, useMessage } from "naive-ui";
import { RouterLink, useRoute, useRouter } from "vue-router";
import { renderIcon } from "@/utils";
import {
  Home,
  ReaderSharp,
  Settings,
  Timer,
  Chatbubbles,
} from "@vicons/ionicons5";
const router = useRouter();
const route = useRoute();

const collapsed = ref(false);
const message = useMessage();
const menusKey = ref("go-" + (route.name as string).split("-")[0]);
function clickMenuItem(key) {
  menusKey.value = key;
}

//后台菜单项
const menus = ref([
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "home",
          },
        },
        { default: () => "返回主页" }
      ),
    key: "go-back-home",
    icon: renderIcon(Home),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "dashboard",
          },
        },
        { default: () => "仪表盘" }
      ),
    key: "go-dashboard",
    icon: renderIcon(Timer),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "article",
          },
        },
        { default: () => "文章管理" }
      ),
    key: "go-article",
    icon: renderIcon(ReaderSharp),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "message",
          },
        },
        { default: () => "留言/问答" }
      ),
    key: "go-message",
    icon: renderIcon(Chatbubbles),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: "other",
          },
        },
        { default: () => "分类管理" }
      ),
    key: "go-other",
    icon: renderIcon(Settings),
  },
]);
</script>

<style scoped></style>
