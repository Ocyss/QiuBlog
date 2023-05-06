import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
import express from "express";
import proxy from "express-http-proxy-2";
import { log } from "node:console";

const isTest = process.env.VITEST;

export async function createServer(
  root = process.cwd(),
  isProd = process.env.NODE_ENV === "production"
) {
  console.log(isProd);
  const __dirname = path.dirname(fileURLToPath(import.meta.url));
  const resolve = (p) => path.resolve(__dirname, p);

  const indexProd = isProd
    ? fs.readFileSync(resolve("dist/client/index.html"), "utf-8")
    : "";

  const manifest = isProd
    ? JSON.parse(
        fs.readFileSync(resolve("dist/client/ssr-manifest.json"), "utf-8")
      )
    : {};
  const baseUrl = isProd ? "api.邱.cf" : "127.0.0.1:3000";
  const app = express();
  /**
   * @type {import('vite').ViteDevServer}
   */
  let vite;

  if (!isProd) {
    vite = await (
      await import("vite")
    ).createServer({
      base: "/",
      root,
      logLevel: isTest ? "error" : "info",
      server: {
        middlewareMode: true,
      },
      appType: "custom",
    });
    app.use(vite.middlewares);
  } else {
    app.use((await import("compression")).default());
    app.use(
      "/",
      (await import("serve-static")).default(resolve("dist/client"), {
        index: false,
      })
    );
  }

  app.use(
    ["/rss/*", "/sitemap.xml"],
    proxy(baseUrl, {
      proxyReqPathResolver: function (req) {
        return req.baseUrl;
      },
    })
  );

  app.use("*", async (req, res) => {
    const url = req.originalUrl;

    try {
      let template, render;
      if (!isProd) {
        template = fs.readFileSync(resolve("index.html"), "utf-8");
        template = await vite.transformIndexHtml(url, template);
        render = (await vite.ssrLoadModule("/src/entry-server.ts")).render;
      } else {
        template = indexProd;
        render = (await import("./dist/server/entry-server.js")).render;
      }

      const { appHtml, cssHtml, preloadLinks, headPayload, teleports } =
        await render(url, manifest);

      const html = template
        .replace(`<!--ssr-outlet-->`, appHtml)
        .replace(`<!--css-outlet-->`, cssHtml)
        .replace(`<!--affix-outlet-->`, teleports["#affixContent"])
        // .replace(`<!--preload-links-->`, preloadLinks)
        .replace(` [!--htmlAttrs--]`, headPayload.htmlAttrs)
        .replace(` [!--bodyAttrs--]`, headPayload.bodyAttrs)
        .replace(`<!--bodyTags-->`, headPayload.bodyTags)
        .replace(`<!--bodyTagsOpen-->`, headPayload.bodyTagsOpen)
        .replace(`<!--headTags-->`, headPayload.headTags);

      res.status(200).set({ "Content-Type": "text/html" }).end(html);
    } catch (e) {
      vite && vite.ssrFixStacktrace(e);

      res.status(500).end(e.stack);
    }
  });

  return { app, vite };
}

createServer().then(({ app }) =>
  app.listen(process.env.npm_package_config_port, () => {
    console.log("http://localhost:" + process.env.npm_package_config_port);
  })
);
