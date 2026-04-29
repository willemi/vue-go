// 路由配置文件
// 定义所有路由规则和路由守卫（beforeEach）
import { createRouter, createWebHistory } from "vue-router";
import { useUserStore } from "../stores/user";

const routes = [
  // 登录页：不需要认证
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
      // Dashboard 主页
      {
        path: "/dashboard",
        name: "Dashboard",
        component: () => import("../views/DashboardView.vue"),
        meta: { requiresAuth: true, title: "Dashboard" },
      },
      // 用户管理：admin 和 user 角色均可访问
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
      // 菜单管理：仅 admin 可访问
      {
        path: "/menu",
        name: "MenuManagement",
        component: () => import("../views/MenuManagementView.vue"),
        meta: { requiresAuth: true, title: "菜单管理", role: ["admin"] },
      },
    ],
  },
  // 404 页面：匹配所有未定义的路由
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

// 路由守卫：在路由切换前执行认证和权限检查
// to: 目标路由，from: 来源路由
router.beforeEach((to, from) => {
  const userStore = useUserStore();
  // 从 localStorage 读取 token（刷新页面后 Pinia store 会被重置）
  const token = localStorage.getItem("token");

  // 需要认证的页面
  if (to.meta.requiresAuth) {
    if (!token) {
      // 未登录，重定向到登录页
      return "/login";
    }
    // 检查角色权限（meta.role 为数组，支持多个角色）
    if (to.meta.role && !to.meta.role.includes(userStore.userInfo.role)) {
      // 角色不匹配，重定向到 Dashboard
      return "/dashboard";
    }
    return true;
  }

  // 已登录用户访问登录页时，重定向到 Dashboard
  if (token && to.path === "/login") {
    return "/dashboard";
  }

  return true;
});

export default router;