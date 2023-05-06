import index from "@/views/FrontView.vue";
import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "home",
    meta: {},
    component: index,
    children: [
      {
        path: "/",
        name: "menuHome",
        component: () => import("@/views/front/HomeView.vue"),
        meta: {
          menukey: "home",
          title: "主页",
        },
        alias: "/home",
      },
      {
        path: "m/tag",
        name: "menuTag",
        component: () => import("@/views/front/TagView.vue"),
        meta: {
          menukey: "tag",
        },
      },
      {
        path: "m/qa",
        name: "menuQa",
        component: () => import("@/views/front/menu/QaView.vue"),
        meta: {
          menukey: "qa",
          title: "问答",
        },
      },
      {
        path: "m/message",
        name: "menuMessage",
        component: () => import("@/views/front/menu/MessageView.vue"),
        meta: {
          menukey: "message",
          title: "留言",
        },
      },
      {
        path: "m/about",
        name: "menuAbout",
        component: () => import("@/views/front/menu/AboutView.vue"),
        meta: {
          menukey: "about",
          title: "关于",
        },
      },
      {
        path: "m/:menuName",
        name: "menu",
        component: () => import("@/views/front/MenuPostView.vue"),
      },
    ],
  },
  {
    path: "/post/:pid",
    name: "post",
    component: () => import("@/views/Post/PostView.vue"),
  },
];

export default routes;
