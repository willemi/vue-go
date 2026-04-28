import { createRouter, createWebHistory } from "vue-router";
import { useUserStore } from "../stores/user";

const routes = [
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/LoginView.vue"),
    meta: { requiresAuth: false },
  },
  {
    path: "/",
    name: "Layout",
    component: () => import("../layout/Layout.vue"),
    redirect: "/dashboard",
    children: [
      {
        path: "/dashboard",
        name: "Dashboard",
        component: () => import("../views/DashboardView.vue"),
        meta: { requiresAuth: true, title: "Dashboard" },
      },
      {
        path: "/user",
        name: "UserManagement",
        component: () => import("../views/UserManagementView.vue"),
        meta: {
          requiresAuth: true,
          title: "用户管理",
          role: ["admin", "user"],
        },
      },
      {
        path: "/menu",
        name: "MenuManagement",
        component: () => import("../views/MenuManagementView.vue"),
        meta: { requiresAuth: true, title: "菜单管理", role: ["admin"] },
      },
    ],
  },
  {
    path: "/:pathMatch(.*)*",
    name: "NotFound",
    component: () => import("../views/NotFoundView.vue"),
    meta: { requiresAuth: false },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// 路由守卫
router.beforeEach((to, from) => {
  const userStore = useUserStore();
  const token = localStorage.getItem("token");

  if (to.meta.requiresAuth) {
    if (!token) {
      return "/login";
    }
    // 检查角色权限
    if (to.meta.role && !to.meta.role.includes(userStore.userInfo.role)) {
      return "/dashboard";
    }
    return true;
  }
  if (token && to.path === "/login") {
    return "/dashboard";
  }
  return true;
});

export default router;
