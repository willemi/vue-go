// Package main Vue 应用入口
// 创建 Vue 实例，注册路由、状态管理、UI 组件库
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { createPinia } from "pinia";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import "./style.css";

// 创建 Vue 应用实例
const app = createApp(App);

// 注册 Pinia 状态管理
app.use(createPinia());

// 注册 Vue Router（路由守卫在 router/index.ts 中配置）
app.use(router);

// 注册 Element Plus UI 组件库
app.use(ElementPlus);

// 全局注册所有 Element Plus 图标组件，使 <component :is="iconName" /> 可以通过字符串名称动态渲染
// 例如：<el-icon><component :is="'User'" /></el-icon>
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

// 将应用挂载到 #app DOM 节点
app.mount("#app");