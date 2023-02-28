import { setTitle } from "@/utils";

const getCookie = (name) =>
  document.cookie.match(`[;\s+]?${name}=([^;]*)`)?.pop();

export function febore(router) {
  router.beforeEach(async (to, from, next) => {
    if (!to.name) {
      //判断有没有路由
      next({ name: "exception-404" });
    }
    if (to.meta.title != undefined) {
      setTitle(to.meta.title);
    }

    if (to.matched[0].name == "admin" && !getCookie("token")) {
      //前往后台，判断是否登陆
      next({ name: "login" });
    } else if (to.name == "login") {
      //前往登陆，判断是否登陆
      if (getCookie("token")) {
        next({ name: "admin" });
      } else {
        next();
      }
    } else {
      next();
    }
  });
  router.afterEach((to, _, failure) => {});
  router.onError((error) => {
    console.log(error, "路由错误");
  });
}
