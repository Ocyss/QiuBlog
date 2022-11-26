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
        meta: {
          menukey: "home",
        },
      },
      {
        path: "m/tag",
        name: "menuTag",
        component: () => import("@/views/front/TagView.vue"),
        meta: {
          menukey: "tag",
        },
      },
      {
        path: "m/qa",
        name: "menuQa",
        component: () => import("@/views/front/menu/QaView.vue"),
        meta: {
          menukey: "qa",
        },
      },
      {
        path: "m/message",
        name: "menuMessage",
        component: () => import("@/views/front/menu/MessageView.vue"),
        meta: {
          menukey: "message",
        },
      },
      {
        path: "m/about",
        name: "menuAbout",
        component: () => import("@/views/front/menu/AboutView.vue"),
        meta: {
          menukey: "about",
        },
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
