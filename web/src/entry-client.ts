import { createApp } from "./main";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";

async function clientInit() {
  const { app, router, pinia } = createApp();
  pinia.use(piniaPluginPersistedstate);

  router.isReady().then(() => {
    app.mount("#app");
  });
}

clientInit();
