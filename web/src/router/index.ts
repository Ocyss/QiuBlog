import {
  createRouter as _createRouter,
  createMemoryHistory,
  createWebHistory,
  RouteRecordRaw,
} from "vue-router";
import { admin, exception, front } from "@/router/modules";

//路由
export const constantRouter: Array<RouteRecordRaw> = [
  ...exception,
  ...front,
  ...admin,
];

export function createRouter() {
  return _createRouter({
    history: import.meta.env.SSR ? createMemoryHistory() : createWebHistory(),
    routes: constantRouter,
  });
}
