import { RouteRecordRaw } from "vue-router";
function randomString() {
  let chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
  let result = "";
  for (let i = 16; i > 0; --i)
    result += chars[Math.floor(Math.random() * chars.length)];
  return "/" + result;
}

const adminPath = import.meta.env.PROD ? randomString() : "/admin";
const LoginPath = import.meta.env.PROD ? randomString() : "/login";
const routes: Array<RouteRecordRaw> = [
  {
    path: adminPath,
    meta: {
      title: "后台管理",
      admin: true,
    },
    component: () => import("@/views/AdminView.vue"),
    children: [
      {
        path: "",
        name: "admin",
        component: () => import("@/views/admin/DashboardView.vue"),
        meta: {
          title: "仪表盘",
          admin: true,
        },
      },
      {
        path: "other",
        name: "other",
        component: () => import("@/views/admin/OtherView.vue"),
        meta: {
          title: "其他设置",
          admin: true,
        },
      },
      {
        path: "message",
        name: "message",
        component: () => import("@/views/admin/MessageView.vue"),
        meta: {
          title: "留言/问答",
          admin: true,
        },
      },
      {
        path: "article",
        name: "article",
        component: () => import("@/views/admin/Article/ArticleView.vue"),
        meta: {
          title: "文章管理",
          admin: true,
        },
      },
      {
        path: "article/updata/:id",
        name: "article-updata",
        component: () => import("@/views/admin/Article/content.vue"),
        meta: {
          title: "修改文章",
          admin: true,
        },
      },
      {
        path: "article/create",
        name: "article-create",
        component: () => import("@/views/admin/Article/content.vue"),
        meta: {
          title: "发布文章",
          admin: true,
        },
      },
    ],
  },
  {
    path: LoginPath,
    name: "login",
    meta: {
      title: "登录",
    },
    component: () => import("@/views/admin/LoginView.vue"),
  },
];

export default routes;
