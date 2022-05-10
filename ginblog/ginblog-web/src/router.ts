import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
  RouteRecordRaw,
} from "vue-router";
import { setDocumentTitle } from "./utils/meta";
import NotFound from "./components/NotFound.vue";
import Layout from "./layout/Index.vue";
import Main from "./views/Main.vue";
import UserProfile from "./views/system/UserProfile.vue";
import UserManage from "./views/system/UserManage.vue";
import CustomDirective from "./views/others/CustomDirective.vue";
import Charts from "./views/others/Charts.vue";
import Icons from "./views/others/Icons.vue";
import Login from "./views/Login.vue";
import Notice from "./views/system/Notice.vue";
let routes: RouteRecordRaw[] = [
  { path: "/", redirect: "/admin/login" },
  { path: "/admin/login", component: Login, name: "login" },

  {
    path: "/admin",
    component: Layout,

    children: [
      {
        path: "home",
        component: Main,
        name: "adminHome",
        meta: {
          title: "首页",
        },
      },
      {
        path: "notice",
        component: Notice,
        name: "notice",
        meta: {
          title: "通知公告",
        },
      },
      {
        path: "userManage",
        component: UserManage,
        name: "userManage",
        meta: {
          title: "用户管理",
        },
      },
      {
        path: "profile",
        component: UserProfile,
        name: "profile",
        meta: {
          title: "用户信息",
        },
      },
      {
        path: "directives",
        component: CustomDirective,
        name: "directives",
        meta: {
          title: "自定义指令",
        },
      },
      {
        path: "charts",
        component: Charts,
        name: "charts",
        meta: {
          title: "图表",
        },
      },
      {
        path: "icons",
        component: Icons,
        name: "icons",
        meta: {
          title: "图标",
        },
      },

      {
        path: "/:pathMatch(.*)*",
        name: "NotFound",
        component: NotFound,
        meta: {
          title: "404",
        },
      },
    ],
  },
];
const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
router.beforeEach((to, from, next) => {
  to.meta &&
    typeof to.meta.title !== "undefined" &&
    setDocumentTitle(`${to.meta.title} `);
  next();
});
export default router;
