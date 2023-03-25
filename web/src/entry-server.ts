import { basename } from "node:path";
import { createApp } from "./main";
import { renderToString } from "vue/server-renderer";
import { setup } from "@css-render/vue3-ssr";

export async function render(url, manifest) {
  const { app, router } = createApp();

  // set the router to the desired URL before rendering
  await router.push(url);
  await router.isReady();

  // passing SSR context object which will be available via useSSRContext()
  // @vitejs/plugin-vue injects code into a component's setup() that registers
  // itself on ctx.modules. After the render, ctx.modules would contain all the
  // components that have been instantiated during this render call.
  const { collect } = setup(app);

  const appHtml = await renderToString(app);
  const cssHtml = collect();
  return {
    cssHtml,
    appHtml,
  };
}
