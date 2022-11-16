import { createApp } from "vue";
import App from "./App.vue";
import VueCookies from "vue-cookies";
import { setupRouter } from "./router";
import { createPinia } from "pinia";
import piniaPersist from "pinia-plugin-persist";
import VueWechatTitle from "vue-wechat-title";
import VueQrcode from "@chenfengyuan/vue-qrcode";

const app = createApp(App);
//依赖注入cookies
app.provide("$cookies", VueCookies);

app.use(VueWechatTitle);
app.component(VueQrcode.name, VueQrcode);

// 挂载状态管理
const pinia = createPinia();
pinia.use(piniaPersist);
app.use(pinia);

//挂载路由
setupRouter(app);

// 挂载app实例
app.mount("#app");
