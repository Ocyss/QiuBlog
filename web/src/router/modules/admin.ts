import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/adminadminadminadminadminadminadminadmin",
    name: "admin",
    meta: {
      title: "后台管理",
    },
    component: () => import("@/views/AdminView.vue"),
    redirect: "/adminadminadminadminadminadminadminadmin/dashboard",
    children: [
      {
        path: "dashboard",
        name: "dashboard",
        component: () => import("@/views/admin/DashboardView.vue"),
        meta: {
          title: "仪表盘",
        },
      },
      {
        path: "other",
        name: "other",
        component: () => import("@/views/admin/OtherView.vue"),
        meta: {
          title: "其他设置",
        },
      },
      {
        path: "message",
        name: "message",
        component: () => import("@/views/admin/MessageView.vue"),
        meta: {
          title: "留言/问答",
        },
      },
      {
        path: "article",
        name: "article",
        component: () => import("@/views/admin/Article/ArticleView.vue"),
        meta: {
          title: "文章管理",
        },
      },
      {
        path: "article/updata/:id",
        name: "article-updata",
        component: () => import("@/views/admin/Article/content.vue"),
        meta: {
          title: "修改文章",
        },
      },
      {
        path: "article/create",
        name: "article-create",
        component: () => import("@/views/admin/Article/content.vue"),
        meta: {
          title: "发布文章",
        },
      },
    ],
  },
  {
    path: "/loginloginloginloginloginlogin",
    name: "login",
    meta: {
      title: "登录",
    },
    component: () => import("@/views/admin/LoginView.vue"),
  },
];

export default routes;
