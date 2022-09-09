<template>
  <NMenu
    :options="menus"
    :collapsed="collapsed"
    :collapsed-width="64"
    :collapsed-icon-size="20"
    :indent="24"
    @update:value="clickMenuItem"
  />
</template>

<script setup>
import { ref, h } from "vue";
import { NIcon, useMessage } from "naive-ui";
import { RouterLink, useRouter } from "vue-router";
import { AddCircleSharp } from "@vicons/ionicons5";
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

function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menus = ref(
  ["主页", "后端", "前端", "生活", "未来", "关于"].map((item) => {
    return {
      label: () =>
        h(
          RouterLink,
          {
            to: {
              name: "home",
            },
          },
          { default: () => item }
        ),
      key: item,
      icon: renderIcon(AddCircleSharp),
    };
  })
);
</script>

<style scoped></style>
