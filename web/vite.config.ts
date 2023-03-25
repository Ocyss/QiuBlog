// vite.config.ts
import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { NaiveUiResolver } from "unplugin-vue-components/resolvers";
import { resolve } from "path";
import UnheadVite from "@unhead/addons/vite";
//打包分析
// import { visualizer } from "rollup-plugin-visualizer";
// import VitePluginCompression from "vite-plugin-compression";//压缩
//CDN 有问题暂时不管了
// import { Plugin as importToCDN, autoComplete } from "vite-plugin-cdn-import";
// import externalGlobals from "rollup-plugin-external-globals";
function pathResolve(dir) {
  return resolve(process.cwd(), ".", dir);
}

const alias = {
  "@": pathResolve("src") + "/",
};

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    // VitePluginCompression(),//压缩gzip
    AutoImport({
      imports: [
        "vue",
        {
          "naive-ui": [
            "useDialog",
            "useMessage",
            "useNotification",
            "useLoadingBar",
          ],
        },
      ],
    }),
    Components({
      resolvers: [NaiveUiResolver()],
    }),
    UnheadVite(),
    // 打包分析
    // visualizer({
    //   emitFile: true, //是否被触摸
    //   filename: "test.html", //生成分析网页文件名
    //   open: true, //在默认用户代理中打开生成的文件
    //   gzipSize: true, //从源代码中收集 gzip 大小并将其显示在图表中
    //   brotliSize: true, //从源代码中收集 brotli 大小并将其显示在图表中
    // }),
    //CDN 配置
    // importToCDN({
    //   modules: [
    //     autoComplete("vue"),
    //     autoComplete("axios"),

    //     {
    //       name: "highlight",
    //       path: "https://cdn.bootcdn.net/ajax/libs/highlight.js/11.7.0/es/highlight.min.js",
    //       css: "https://cdn.bootcdn.net/ajax/libs/highlight.js/11.7.0/styles/atom-one-dark.min.css",
    //     },
    //     {
    //       name: "pinia",
    //       path: "https://cdn.bootcdn.net/ajax/libs/pinia/2.0.26/pinia.d.ts",
    //     },
    //     {
    //       name: "wangeditor",
    //       path: "https://cdn.bootcdn.net/ajax/libs/wangeditor5/5.1.23/index.min.js",
    //       css: "https://cdn.bootcdn.net/ajax/libs/wangeditor5/5.1.23/css/style.css",
    //     },
    //     {
    //       name: "lottie-web",
    //       path: "https://cdn.bootcdn.net/ajax/libs/lottie-web/5.9.6/lottie.d.ts",
    //     },
    //     {
    //       name: "marked",
    //       path: "https://cdn.bootcdn.net/ajax/libs/marked/4.2.3/lib/marked.esm.js",
    //     },
    //     {
    //       name: "scss",
    //       path: "https://cdn.bootcdn.net/ajax/libs/sass.js/0.11.1/sass.js",
    //     },
    //   ],
    // }),
  ],
  resolve: {
    alias,
  },
  server: {
    host: "0.0.0.0",
    port: 6879,
    proxy: {
      "^/(api|config|sitemap|about.md)": {
        target: "http://127.0.0.1:3000",
        changeOrigin: true,
      },
    },
  },
  css: {
    preprocessorOptions: {
      scss: {
        modifyVars: {},
        javascriptEnabled: true,
        additionalData: `@import "src/styles/index.scss";`,
      },
    },
  },
  build: {
    //CDM 配置
    // rollupOptions: {
    //   // 👇 告诉打包工具 "vue-demi" 也是外部依赖项 👇
    //   external: ["vue", "vue-demi"],
    //   plugins: [
    //     externalGlobals({
    //       vue: "Vue",
    //       // 👇 配置 vue-demi 全局变量 👇
    //       "vue-demi": "VueDemi",
    //     }),
    //   ],
    // },
  },
});
