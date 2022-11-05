import { ErrorPage, RedirectName, Layout } from "@/router/constant";

// 404 on a page
export const ErrorPageRoute = {
  path: "/:path(.*)*",
  name: "ErrorPage",
  component: Layout,
  meta: {
    title: "ErrorPage",
    hideBreadcrumb: true,
  },
  children: [
    {
      path: "/:path(.*)*",
      name: "ErrorPageSon",
      component: ErrorPage,
      meta: {
        title: "ErrorPage",
        hideBreadcrumb: true,
      },
    },
  ],
};
