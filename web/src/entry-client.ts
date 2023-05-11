import { createApp } from "./main";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";
import VueCookies from 'vue-cookies'

// default options config: { expires: '1d', path: '/', domain: '', secure: '', sameSite: 'Lax' }


async function clientInit() {
  const { app, router, pinia } = createApp();
  pinia.use(piniaPluginPersistedstate);
  app.use(VueCookies, { expires: '7d'})
  router.isReady().then(() => {
    app.mount("#app");
  });
}

clientInit();
    