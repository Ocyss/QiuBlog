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
const collapsed = ref(false);
const router = useRouter();
const message = useMessage();

//切换左侧菜单
function clickMenuItem(key) {
  if (/http(s)?:/.test(key)) {
    window.open(key);
  } else {
    router.push({ name: "home", query: { menu: key } });
    menusKey.value = key;
  }
}
//图标赋值
function renderIcon(icon) {
  return () => h(NIcon, { innerHTML: icon });
}
//主页菜单
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
        { default: () => "主页" }
      ),
    key: "home",
    icon: renderIcon(
      `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512"><path d="M261.56 101.28a8 8 0 0 0-11.06 0L66.4 277.15a8 8 0 0 0-2.47 5.79L63.9 448a32 32 0 0 0 32 32H192a16 16 0 0 0 16-16V328a8 8 0 0 1 8-8h80a8 8 0 0 1 8 8v136a16 16 0 0 0 16 16h96.06a32 32 0 0 0 32-32V282.94a8 8 0 0 0-2.47-5.79z" fill="currentColor"></path><path d="M490.91 244.15l-74.8-71.56V64a16 16 0 0 0-16-16h-48a16 16 0 0 0-16 16v32l-57.92-55.38C272.77 35.14 264.71 32 256 32c-8.68 0-16.72 3.14-22.14 8.63l-212.7 203.5c-6.22 6-7 15.87-1.34 22.37A16 16 0 0 0 43 267.56L250.5 69.28a8 8 0 0 1 11.06 0l207.52 198.28a16 16 0 0 0 22.59-.44c6.14-6.36 5.63-16.86-.76-22.97z" fill="currentColor"></path></svg>`
    ),
  },
]);
//当前菜单选择key
const menusKey = ref("home");

onBeforeMount(() => {
  //请求菜单项
  axios.get("/api/v1/menuchild").then((res) => {
    if (res.data.status == 200) {
      res.data.data.map((item) => {
        menus.value.push({
          label: () =>
            h(
              RouterLink,
              {
                to: {
                  name: "home",
                },
              },
              { default: () => item.name }
            ),
          key: item.ename,
          icon: renderIcon(item.logo),
        });
      });
    }
  });
});
</script>

<style scoped></style>
