export function febore(router) {
  router.beforeEach(async (to, from, next) => {
    if (!to.name) {
      next({ name: "exception-404" });
    } else {
      next();
    }
  });
  router.afterEach((to, _, failure) => {});
  router.onError((error) => {
    console.log(error, "路由错误");
  });
}
