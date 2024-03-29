<template>
  <NMenu
    :options="menusRef"
    :collapsed="collapsed"
    :value="menusKey"
    :collapsed-width="64"
    :collapsed-icon-size="20"
    :indent="24"
    @update:value="clickMenuItem"
  />
</template>

<script setup lang="ts">
import { ref, h, onServerPrefetch, inject, computed, onMounted } from "vue";
import { NIcon, useMessage } from "naive-ui";
import { RouterLink, useRouter, useRoute } from "vue-router";
import api from "@/api";
import { useDesignSettingStore } from "@/store/modules/designSetting";
import { useHead } from "@unhead/vue";
import axios from "axios";

const props = defineProps(["collapsed"]);
const designStore = useDesignSettingStore();
const route = useRoute();
const router = useRouter();
const message = useMessage();
//当前菜单选择key
const menusKey = ref("home");
const menus = ref([
  {
    name: "主页",
    ename: "Home",
    linkName: "home",
    link: "home",
    icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512"><path d="M261.56 101.28a8 8 0 0 0-11.06 0L66.4 277.15a8 8 0 0 0-2.47 5.79L63.9 448a32 32 0 0 0 32 32H192a16 16 0 0 0 16-16V328a8 8 0 0 1 8-8h80a8 8 0 0 1 8 8v136a16 16 0 0 0 16 16h96.06a32 32 0 0 0 32-32V282.94a8 8 0 0 0-2.47-5.79z" fill="currentColor"></path><path d="M490.91 244.15l-74.8-71.56V64a16 16 0 0 0-16-16h-48a16 16 0 0 0-16 16v32l-57.92-55.38C272.77 35.14 264.71 32 256 32c-8.68 0-16.72 3.14-22.14 8.63l-212.7 203.5c-6.22 6-7 15.87-1.34 22.37A16 16 0 0 0 43 267.56L250.5 69.28a8 8 0 0 1 11.06 0l207.52 198.28a16 16 0 0 0 22.59-.44c6.14-6.36 5.63-16.86-.76-22.97z" fill="currentColor"></path></svg>`,
  },
  {
    name: "问答",
    ename: "Q&A",
    linkName: "menuQa",
    link: "qa",
    icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512"><path d="M64 480H48a32 32 0 0 1-32-32V112a32 32 0 0 1 32-32h16a32 32 0 0 1 32 32v336a32 32 0 0 1-32 32z" fill="currentColor"></path><path d="M240 176a32 32 0 0 0-32-32h-64a32 32 0 0 0-32 32v28a4 4 0 0 0 4 4h120a4 4 0 0 0 4-4z" fill="currentColor"></path><path d="M112 448a32 32 0 0 0 32 32h64a32 32 0 0 0 32-32v-30a2 2 0 0 0-2-2H114a2 2 0 0 0-2 2z" fill="currentColor"></path><rect x="112" y="240" width="128" height="144" rx="2" ry="2" fill="currentColor"></rect><path d="M320 480h-32a32 32 0 0 1-32-32V64a32 32 0 0 1 32-32h32a32 32 0 0 1 32 32v384a32 32 0 0 1-32 32z" fill="currentColor"></path><path d="M495.89 445.45l-32.23-340c-1.48-15.65-16.94-27-34.53-25.31l-31.85 3c-17.59 1.67-30.65 15.71-29.17 31.36l32.23 340c1.48 15.65 16.94 27 34.53 25.31l31.85-3c17.59-1.67 30.65-15.71 29.17-31.36z" fill="currentColor"></path></svg>`,
  },
  {
    name: "留言",
    ename: "Message",
    linkName: "menuMessage",
    link: "message",
    icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512"><path d="M408 48H104a72.08 72.08 0 0 0-72 72v192a72.08 72.08 0 0 0 72 72h24v64a16 16 0 0 0 26.25 12.29L245.74 384H408a72.08 72.08 0 0 0 72-72V120a72.08 72.08 0 0 0-72-72zM160 248a32 32 0 1 1 32-32a32 32 0 0 1-32 32zm96 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32zm96 0a32 32 0 1 1 32-32a32 32 0 0 1-32 32z" fill="currentColor"></path></svg>`,
  },
  {
    name: "关于",
    ename: "About",
    linkName: "menuAbout",
    link: "about",
    icon: `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 512 512"><path d="M336 336H32a16 16 0 0 1-14-23.81l152-272a16 16 0 0 1 27.94 0l152 272A16 16 0 0 1 336 336z" fill="currentColor"></path><path d="M336 160a161.07 161.07 0 0 0-32.57 3.32l74.47 133.27A48 48 0 0 1 336 368H183.33A160 160 0 1 0 336 160z" fill="currentColor"></path></svg>`,
  },
]);

if (!route.meta.menukey) {
  menusKey.value = route.params.menuName as string;
} else {
  menusKey.value = route.meta.menukey as string;
}
//获取菜单名并设置标题
function getMenuName(key) {
  let menu = menus.value.find((item) => {
    return key == item.link;
  });
  if (menu == undefined) {
    menu = menus.value[0];
  }
  if (
    route.name == "home" ||
    route.name == "menuQa" ||
    route.name == "menuMessage" ||
    route.name == "menuAbout" ||
    route.name == "menu"
  ) {
    useHead({ title: designStore.getLocale ? menu.name : menu.ename });
  }
}
//切换左侧菜单
function clickMenuItem(key, item) {
  menusKey.value = key;
  useHead({ title: designStore.getLocale ? item.name : item.ename });
}
//图标赋值
function renderIcon(icon) {
  return () => h(NIcon, { innerHTML: icon });
}

async function getMenus() {
  //请求菜单项
  const res = await api.menuchild.gets();

  res.data.map((item) => {
    menus.value.splice(1, 0, {
      name: item.name,
      ename: item.ename,
      linkName: "menu",
      link: item.link,
      icon: item.logo,
    });
  });
}
onMounted(() => {
  if (menus.value.length == 4) {
    getMenus();
  }
  getMenuName(menusKey.value);
});
onServerPrefetch(() => {
  getMenus();
  getMenuName(menusKey.value);
});

// let menusRef = undefined;
//主页菜单
const menusRef = computed(() => {
  return menus.value.map((item) => {
    return {
      label: () =>
        h(
          RouterLink,
          {
            to: {
              name: item.linkName,
              params:
                item.linkName == "menu"
                  ? {
                      menuName: item.link,
                    }
                  : {},
            },
          },
          { default: () => (designStore.getLocale ? item.name : item.ename) }
        ),

      name: item.name,
      ename: item.ename,
      key: item.link,
      icon: renderIcon(item.icon),
    };
  });
});
</script>

<style scoped></style>
