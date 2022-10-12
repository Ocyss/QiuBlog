import { createRouter, createWebHashHistory } from "vue-router";
// import { RedirectRoute } from "@/router/base";
import { admin, exception, front } from "@/router/modules";
//需要验证权限
export const asyncRoutes = admin;

//普通路由 无需验证权限
export const constantRouter = [
  ...exception,
  ...front,
  // RedirectRoute,
  ...asyncRoutes,
];

const router = createRouter({
  history: createWebHashHistory(""),
  routes: constantRouter,
  strict: true,
  scrollBehavior: () => ({ left: 0, top: 0 }),
});

export default router;
