import index from "@/views/FrontView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    meta: {
      title: "主页",
    },
    component: index,
    redirect: "/m/home",
    children: [
      {
        path: "m/home",
        name: "menuHome",
        component: () => import("@/views/front/HomeView.vue"),
      },
      {
        path: "m/tag",
        name: "menuTag",
        component: () => import("@/views/front/TagView.vue"),
      },
      {
        path: "m/qa",
        name: "menuQa",
        component: () => import("@/views/front/menu/QaView.vue"),
      },
      {
        path: "m/message",
        name: "menuMessage",
        component: () => import("@/views/front/menu/MessageView.vue"),
      },
      {
        path: "m/about",
        name: "menuAbout",
        component: () => import("@/views/front/menu/AboutView.vue"),
      },
      {
        path: "m/:menuName",
        name: "menu",
        component: () => import("@/views/front/MenuPostView.vue"),
      },
    ],
  },
  {
    path: "/post/:pid",
    name: "post",
    component: () => import("@/views/Post/PostView.vue"),
  },
];

export default routes;
