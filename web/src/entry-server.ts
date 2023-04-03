import { basename } from "node:path";
import { createApp } from "./main";
import { renderToString } from "vue/server-renderer";
import { setup } from "@css-render/vue3-ssr";
import { renderSSRHead } from "@unhead/ssr";

export async function render(url, manifest) {
  const { app, router, pinia, head } = createApp();

  await router.push(url);
  await router.isReady();

  const { collect } = setup(app);
  const ctx: any = {};
  const appHtml = await renderToString(app, ctx);
  const cssHtml = collect();
  const preloadLinks = renderPreloadLinks(ctx.modules, manifest);

  const headPayload = await renderSSRHead(head);
  return { appHtml, cssHtml, preloadLinks, headPayload };
}

function renderPreloadLinks(modules, manifest) {
  let links = "";
  const seen = new Set();
  modules.forEach((id) => {
    const files = manifest[id];
    if (files) {
      files.forEach((file) => {
        if (!seen.has(file)) {
          seen.add(file);
          const filename = basename(file);
          if (manifest[filename]) {
            for (const depFile of manifest[filename]) {
              links += renderPreloadLink(depFile);
              seen.add(depFile);
            }
          }
          links += renderPreloadLink(file);
        }
      });
    }
  });
  return links;
}

function renderPreloadLink(file) {
  if (file.endsWith(".js")) {
    return `<link rel="modulepreload" crossorigin href="${file}">`;
  } else if (file.endsWith(".css")) {
    return `<link rel="stylesheet" href="${file}">`;
  } else if (file.endsWith(".woff")) {
    return ` <link rel="preload" href="${file}" as="font" type="font/woff" crossorigin>`;
  } else if (file.endsWith(".woff2")) {
    return ` <link rel="preload" href="${file}" as="font" type="font/woff2" crossorigin>`;
  } else if (file.endsWith(".gif")) {
    return ` <link rel="preload" href="${file}" as="image" type="image/gif">`;
  } else if (file.endsWith(".jpg") || file.endsWith(".jpeg")) {
    return ` <link rel="preload" href="${file}" as="image" type="image/jpeg">`;
  } else if (file.endsWith(".png")) {
    return ` <link rel="preload" href="${file}" as="image" type="image/png">`;
  } else {
    // TODO
    return "";
  }
}
