import { createApp } from "./main";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import { renderToString } from "vue/server-renderer";

async function clientInit() {
  const { app, router, pinia } = createApp();
  pinia.use(piniaPluginPersistedstate);

  router.isReady().then(() => {
    app.mount("#app");
  });
}

clientInit();
