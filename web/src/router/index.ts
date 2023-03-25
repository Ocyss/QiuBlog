import {
  createRouter as _createRouter,
  createMemoryHistory,
  createWebHistory,
  RouteRecordRaw,
} from "vue-router";
import { admin, exception, front } from "@/router/modules";
import { febore } from "./router-guards";

//路由
export const constantRouter: Array<RouteRecordRaw> = [
  ...exception,
  ...front,
  ...admin,
];

// const router = createRouter({
//   history: createWebHistory(),
//   routes: constantRouter,
//   strict: true,
//   scrollBehavior: () => ({ left: 0, top: 0 }),
// });

// 路由守卫
// export function setupRouter(app) {
//   app.use(router);
//   febore(router);
// }

export function createRouter() {
  return _createRouter({
    // use appropriate history implementation for server/client
    // import.meta.env.SSR is injected by Vite.
    history: import.meta.env.SSR ? createMemoryHistory() : createWebHistory(),
    routes: constantRouter,
  });
}
