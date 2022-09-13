import { createApp } from "vue";
import App from "./App.vue";
import VueCookies from "vue-cookies";
import router from "./router";
import axios from "axios";
import { createPinia } from "pinia";
import piniaPersist from "pinia-plugin-persist";

const pinia = createPinia();
pinia.use(piniaPersist);
const app = createApp(App);
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$cookies = VueCookies;
app.use(router).use(pinia).mount("#app");
