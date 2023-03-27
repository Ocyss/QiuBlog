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
  // legacy: { buildSsrCjsExternalHeuristics: true },
  // build: {
  //   sourcemap: false,
  //   minify: "esbuild",
  //   chunkSizeWarningLimit: 800,

  // rollupOptions: {
  //   output: {
  //     //静态资源分类打包
  //     // chunkFileNames: "assets/js/[name]-[hash].js",
  //     // entryFileNames: "assets/js/[name]-[hash].js",
  //     // assetFileNames: "assets/[ext]/[name]-[hash].[ext]",
  //     manualChunks(id) {
  //       if (id.includes("naive-ui") || id.includes("vue")) {
  //         return;
  //       }
  //       //静态资源分拆打包
  //       if (id.includes("node_modules")) {
  //         return id
  //           .toString()
  //           .split("node_modules/")[1]
  //           .split("/")[0]
  //           .toString();
  //       }
  //     },
  //   },
  // },
  // },
});
