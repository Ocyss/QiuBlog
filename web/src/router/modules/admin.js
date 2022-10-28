const routes = [
  {
    path: "/admin",
    name: "admin",
    meta: {
      title: "后台管理",
    },
    component: () => import("@/views/AdminView.vue"),
    redirect: "/admin/dashboard",
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
];

export default routes;
