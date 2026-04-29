// Pinia 用户状态管理
// 管理登录令牌（token）和用户信息（userInfo），并持久化到 localStorage
// 登录状态在页面刷新后通过 localStorage 恢复
import { defineStore } from "pinia";
import { ref } from "vue";

interface UserInfo {
  id: number;
  username: string;
  role: string;
}

// 定义 user store，命名为 "user"
export const useUserStore = defineStore("user", () => {
  // token：从 localStorage 初始化，页面刷新后仍有效
  const token = ref<string>(localStorage.getItem("token") || "");
  // userInfo：存储用户基本信息，同样从 localStorage 恢复
  const userInfo = ref<UserInfo>(
    JSON.parse(localStorage.getItem("userInfo") || "{}"),
  );

  // setToken 保存登录令牌到 store 和 localStorage
  function setToken(newToken: string) {
    token.value = newToken;
    localStorage.setItem("token", newToken);
  }

  // setUserInfo 保存用户信息到 store 和 localStorage
  function setUserInfo(user: UserInfo) {
    userInfo.value = user;
    localStorage.setItem("userInfo", JSON.stringify(user));
  }

  // logout 清除所有登录状态并刷新页面
  // localStorage 清除后，路由守卫会检测到无 token 跳转登录页
  function logout() {
    token.value = "";
    userInfo.value = { id: 0, username: "", role: "" };
    localStorage.removeItem("token");
    localStorage.removeItem("userInfo");
    location.reload();
  }

  return {
    token,
    userInfo,
    setToken,
    setUserInfo,
    logout,
  };
});