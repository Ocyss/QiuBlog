import { basename } from "node:path";
import { createApp } from "./main";
import { renderToString } from "vue/server-renderer";
import { renderSSRHead } from "@unhead/ssr";

export async function render(url, manifest) {
  const { app, router, pinia, head, collect } = createApp();

  await router.push(url);
  await router.isReady();

  const ctx: any = {};
  const appHtml = await renderToString(app, ctx);
  const cssHtml = collect();
  const preloadLinks = renderPreloadLinks(cssHtml, ctx.modules, manifest);
  const headPayload = await renderSSRHead(head);
  return { appHtml, preloadLinks, headPayload };
}

function renderPreloadLinks(css, modules, manifest) {
  let links = css;
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
