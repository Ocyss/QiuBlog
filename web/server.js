import fs from "fs";
import path from "node:path";
import { fileURLToPath } from "url";
import express from "express";
import { createServer as createViteServer } from "vite";

const isTest = process.env.VITEST;

export async function createServer(
  root = process.cwd(),
  isProd = process.env.NODE_ENV === "production",
  hmrPort
) {
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

  const app = express();

  // const vite = await createViteServer({
  //   server: { middlewareMode: true },
  //   appType: "custom",
  // });

  // app.use(vite.middlewares);
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
        watch: {
          // During tests we edit the files too fast and sometimes chokidar
          // misses change events, so enforce polling for consistency
          usePolling: true,
          interval: 100,
        },
        hmr: {
          port: hmrPort,
        },
      },
      appType: "custom",
    });
    // use vite's connect instance as middleware
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

  app.use("*", async (req, res) => {
    const url = req.originalUrl;

    try {
      // let template = fs.readFileSync(
      //   path.resolve(__dirname, "index.html"),
      //   "utf-8"
      // );

      // template = await vite.transformIndexHtml(url, template);

      // const { render } = import("./dist/server/entry-server.js");

      // const { appHtml, cssHtml } = await render(url);
      let template, render;
      if (!isProd) {
        // always read fresh template in dev
        template = fs.readFileSync(resolve("index.html"), "utf-8");
        template = await vite.transformIndexHtml(url, template);
        render = (await vite.ssrLoadModule("/src/entry-server.ts")).render;
      } else {
        template = indexProd;
        // @ts-ignore
        render = (await import("./dist/server/entry-server.js")).render;
      }

      const [appHtml, cssHtml, preloadLinks] = await render(url, manifest);

      const html = template
        .replace(`<!--ssr-outlet-->`, appHtml)
        .replace(`<!--css-outlet-->`, cssHtml)
        .replace(`<!--preload-links-->`, preloadLinks);

      res.status(200).set({ "Content-Type": "text/html" }).end(html);
    } catch (e) {
      vite && vite.ssrFixStacktrace(e);
      console.log(e.stack);
      res.status(500).end(e.stack);
    }
  });

  return { app, vite };
}

createServer().then(({ app }) =>
  app.listen(6879, () => {
    console.log("http://localhost:6879");
  })
);
