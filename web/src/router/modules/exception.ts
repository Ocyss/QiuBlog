import { Layout } from "@/router/constant";
import { AlertCircle } from "@vicons/ionicons5";
import { renderIcon } from "@/utils/index";
import { RouteRecordRaw } from "vue-router";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/exception",
    name: "Exception",
    redirect: "/exception/403",
    component: Layout,
    meta: {
      title: "异常页面",
      icon: renderIcon(AlertCircle),
      sort: 3,
    },
    children: [
      {
        path: "403",
        name: "exception-403",
        meta: {
          title: "403",
        },
        component: () => import("@/views/exception/403.vue"),
      },
      {
        path: "404",
        name: "exception-404",
        meta: {
          title: "404",
        },
        component: () => import("@/views/exception/404.vue"),
      },
      {
        path: "500",
        name: "exception-500",
        meta: {
          title: "500",
        },
        component: () => import("@/views/exception/500.vue"),
      },
    ],
  },
];

export default routes;
