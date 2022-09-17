import { createRouter, createWebHashHistory } from "vue-router";
import index from "@/views/HomeView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    meta: {
      title: "主页",
    },
    component: index,
  },
  {
    path: "/admin",
    name: "admin",
    meta: {
      title: "后台管理",
    },
    component: () => import("@/views/AdminView.vue"),
  },
  {
    path: "/admin/article",
    name: "article",
    component: () => import("@/views/admin/ArticleView.vue"),
    meta: {
      title: "文章管理",
    },
  },
  {
    path: "/admin/other",
    name: "other",
    component: () => import("@/views/admin/OtherView.vue"),
    meta: {
      title: "其他设置",
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
