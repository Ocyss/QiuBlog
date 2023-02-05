import { createApp } from "vue";
import App from "./App.vue";
import { setupRouter } from "./router";
import { createPinia } from "pinia";
import piniaPersist from "pinia-plugin-persist";
import VueWechatTitle from "vue-wechat-title";
import VueQrcode from "@chenfengyuan/vue-qrcode";
import hljs from "highlight.js";
import "highlight.js/styles/atom-one-dark.css";

const app = createApp(App);

app.use(VueWechatTitle);
app.component(VueQrcode.name, VueQrcode);

//创建v-highlight全局指令
app.directive("highlight", function (el) {
  let blocks = el.querySelectorAll("pre code");
  blocks.forEach((block) => {
    hljs.highlightBlock(block);
  });
});

// 挂载状态管理
const pinia = createPinia();
pinia.use(piniaPersist);
app.use(pinia);

//挂载路由
setupRouter(app);

// 挂载app实例
app.mount("#app");
