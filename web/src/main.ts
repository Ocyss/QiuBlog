import { createSSRApp } from "vue";
import App from "./App.vue";
// import VueCookies from "vue-cookies";
import VueQrcode from "@chenfengyuan/vue-qrcode";
import { createHead } from "@unhead/vue";
import { createPinia } from "pinia";
// import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import { createRouter } from "./router";
import { febore } from "./router/router-guards";

// 暴露统一方法createApp
export function createApp() {
  const app = createSSRApp(App);
  app.component(VueQrcode.name, VueQrcode);

  // 挂载head插件
  const head = createHead();
  app.use(head);

  // 挂载状态管理
  const pinia = createPinia();
  // pinia.use(piniaPluginPersistedstate);

  // 挂载路由
  // setupRouter(app);
  const router = createRouter();
  febore(router);
  app.use(router);

  app.use(pinia);
  // 挂载Cookie插件
  // app.use(VueCookies, { expires: "7d" });
  return { app, router };
}
