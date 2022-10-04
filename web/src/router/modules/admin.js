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
        component: () => import("@/views/admin/ArticleView.vue"),
        meta: {
          title: "文章管理",
        },
      },
    ],
  },
];

export default routes;
