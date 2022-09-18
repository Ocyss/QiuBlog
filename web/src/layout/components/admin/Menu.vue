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

<script setup>
import { ref, h, onBeforeMount } from "vue";
import { NIcon, useMessage } from "naive-ui";
import { RouterLink, useRouter } from "vue-router";
import axios from "axios";
import { renderIcon } from "@/utils/index.js";
import { Home, ReaderSharp, Settings } from "@vicons/ionicons5";
const collapsed = ref(false);
const router = useRouter();
const message = useMessage();

function clickMenuItem(key) {
  // if (/http(s)?:/.test(key)) {
  //   window.open(key);
  // } else {
  //   router.push({ name: key });
  // }
  // emit("clickMenuItem", key);
  message.error("还没写好嘞");
}

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
            name: "other",
          },
        },
        { default: () => "其他设置" }
      ),
    key: "go-other",
    icon: renderIcon(Settings),
  },
]);

const menusKey = ref("主页");
onBeforeMount(() => {});
</script>

<style scoped></style>
