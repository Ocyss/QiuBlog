import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";
import express from "express";
import axios from "axios";
import compression from "compression";
import serveStatic from "serve-static";
import proxy from "express-http-proxy";

export async function createServer(root = process.cwd()) {
  const __dirname = path.dirname(fileURLToPath(import.meta.url));
  const resolve = (p) => path.resolve(__dirname, p);

  const indexProd = fs.readFileSync(resolve("client/index.html"), "utf-8");

  const manifest = JSON.parse(
    fs.readFileSync(resolve("client/ssr-manifest.json"), "utf-8")
  );

  const baseUrl = "http://127.0.0.1:16879";
  const app = express();
  /**
   * @type {import('vite').ViteDevServer}
   */
  let vite;

  app.use(compression());
  app.use(
    "/",
    serveStatic(resolve("client"), {
      index: false,
    })
  );

  app.use(express.static("files"));
  app.get(["/rss/*", "/sitemap.xml"], async (req, res) => {
    const r = await axios.get(baseUrl + req.url);
    for (let h in r.headers) {
      res.set(h, r.headers[h]);
    }
    res.send(r.data);
  });

  app.use("/api", (req, res) => {
    proxy("http://127.0.0.1:16879", {
      proxyReqPathResolver: function (req) {
        return req.originalUrl;
      },
    })(req, res);
  });

  app.use("*", async (req, res) => {
    const url = req.originalUrl;

    try {
      const template = indexProd;
      const render = (await import(resolve("server/entry-server.js"))).render;

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
  app.listen(3000, () => {
    console.log(
      `
       ____  _       ____  _
      / __ \\(_)     |  _ \\| |
     | |  | |_ _   _| |_) | | ___   __ _
     | |  | | | | | |  _ <| |/ _ \\ / _\` |
     | |__| | | |_| | |_) | | (_) | (_| |
      \\___\\_\\_|\\__,_|____/|_|\\___/ \\__, |
                                    __/ |
                                   |___/
    http://localhost:3000`
    );
  })
);

//   ____  _       ____  _
//  / __ \(_)     |  _ \| |
// | |  | |_ _   _| |_) | | ___   __ _
// | |  | | | | | |  _ <| |/ _ \ / _` |
// | |__| | | |_| | |_) | | (_) | (_| |
//  \___\_\_|\__,_|____/|_|\___/ \__, |
//                                __/ |
//                               |___/
