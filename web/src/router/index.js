import { createRouter, createWebHashHistory } from "vue-router";
import { admin, exception, front } from "@/router/modules";
import { febore } from "./router-guards";

//需要验证权限
export const asyncRoutes = admin;

//普通路由
export const constantRouter = [...exception, ...front];

const router = createRouter({
  history: createWebHashHistory(import.meta.env.VITE_BASE_PATH),
  routes: constantRouter,
  strict: true,
  scrollBehavior: () => ({ left: 0, top: 0 }),
});

// 路由守卫
export function setupRouter(app) {
  app.use(router);
  febore(router);
}
export default router;
