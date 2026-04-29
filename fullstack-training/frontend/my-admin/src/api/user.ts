// API 模块：封装所有与后端的 HTTP 通信
// 使用 axios 发送请求，请求拦截器自动注入 JWT token
// 响应拦截器统一处理错误（如 401 未认证则自动登出）
import axios from "axios";
import { useUserStore } from "../stores/user";

// 创建 axios 实例，基础 URL 指向后端 API
// 避免每个请求单独指定 baseURL
const api = axios.create({
  baseURL: "http://localhost:8080/api",
  timeout: 10000, // 请求超时 10 秒
});

// ========== 请求拦截器 ==========
// 每次发送请求前执行，自动将 token 加入 Authorization header
api.interceptors.request.use(
  (config) => {
    const userStore = useUserStore();
    if (userStore.token) {
      config.headers.Authorization = `Bearer ${userStore.token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

// ========== 响应拦截器 ==========
// 每次收到响应后执行，统一处理成功响应和错误
api.interceptors.response.use(
  (response) => {
    // 成功响应只返回 response.data（去掉 axios 包装）
    return response.data;
  },
  (error) => {
    // 401 未认证：token 过期或无效，清除登录状态并跳转登录页
    if (error.response?.status === 401) {
      const userStore = useUserStore();
      userStore.logout();
      window.location.href = "/login";
    }
    // 其他错误：返回原始错误信息，由调用方处理
    return Promise.reject(error.response?.data ?? error);
  },
);

// ========== 用户相关 API ==========

// login 用户登录
export const login = (username: string, password: string) => {
  return api.post("/user/login", { username, password });
};

// getUserList 获取用户列表（支持分页和搜索）
export const getUserList = (params: {
  username?: string;
  page?: number;
  page_size?: number;
}) => {
  return api.get("/user/list", { params });
};

// createUser 创建新用户
export const createUser = (data: {
  username: string;
  password: string;
  role?: string;
}) => {
  return api.post("/user/add", data);
};

// updateUser 更新用户信息
export const updateUser = (data: {
  id: number;
  username: string;
  password?: string;
  role?: string;
}) => {
  return api.put("/user/edit", data);
};

// deleteUser 删除用户
export const deleteUser = (id: number) => {
  return api.delete(`/user/delete/${id}`);
};

// ========== 菜单相关 API ==========

// getMenuList 获取菜单列表（扁平，用于菜单管理页面）
export const getMenuList = () => {
  return api.get("/menu/list");
};

// getMenuTree 获取当前用户可见的菜单树（用于侧边栏）
export const getMenuTree = () => {
  return api.get("/menu/tree");
};

// createMenu 创建新菜单
export const createMenu = (data: {
  title: string;
  path: string;
  icon?: string;
  parent_id?: number;
  sort?: number;
  hidden?: boolean;
  role?: string;
}) => {
  return api.post("/menu/add", data);
};

// updateMenu 更新菜单
export const updateMenu = (data: {
  id: number;
  title: string;
  path: string;
  icon?: string;
  parent_id?: number;
  sort?: number;
  hidden?: boolean;
  role?: string;
}) => {
  return api.put("/menu/edit", data);
};

// deleteMenu 删除菜单
export const deleteMenu = (id: number) => {
  return api.delete(`/menu/delete/${id}`);
};